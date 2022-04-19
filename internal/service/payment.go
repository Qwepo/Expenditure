package service

import (
	"app/internal/db"
	"time"

	"github.com/jackc/pgx/v4"
)

type Payment interface {
	PaymentCreate(*PaymentFullRequest) (int, error)
}

type PaymentService struct {
	db              db.DB
	counterparty    Counterparty
	expenditureItem ExpenditureItem
	project         Project
}

type PaymentFullRequest struct {
	Doctype             *string `json:"doctype"`
	OrganizationName    *string `json:"organizationName"`
	CounterpartyName    *string `json:"counterpartyName"`
	IncomingCurrency    *int    `json:"incomingCurrency"`
	ExpendableCurrency  *int    `json:"expendableCurrency"`
	Purpose             *string `json:"purpose"`
	ExpenditureItemName *string `json:"expenditureItemName"`
	ProjectName         *string `json:"projectName"`
	Comments            string  `json:"comments"`
}

func (r *PaymentFullRequest) fillToPayment(p *db.Payment, cpID, exID, prID int) {
	time := time.Now()
	p.CreatedAt = &time
	p.Doctype = r.Doctype
	p.IncomingCurrency = r.IncomingCurrency
	p.ExpendableCurrency = r.ExpendableCurrency
	p.OrganizationName = r.OrganizationName
	p.Purpose = r.Purpose
	p.Comments = r.Comments
	p.CounterpartyID = &cpID
	p.ExpenditureItemID = &exID
	p.ProjectID = &prID
}

func (p *PaymentService) PaymentCreate(resp *PaymentFullRequest) (int, error) {
	var payment db.Payment
	cpID, err := p.counterparty.CounterpartyFindeByName(resp.CounterpartyName)
	if err == pgx.ErrNoRows {
		cpID, err = p.counterparty.CounterpartyCreate(resp.CounterpartyName)
		if err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}

	exID, err := p.expenditureItem.ExpenditureItemCreate(resp.ExpenditureItemName)
	if err == pgx.ErrNoRows {
		exID, err = p.expenditureItem.ExpenditureItemCreate(resp.ExpenditureItemName)
		if err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}

	prID, err := p.project.ProjectCreate(resp.ProjectName)
	if err == pgx.ErrNoRows {
		prID, err = p.project.ProjectCreate(resp.ProjectName)
		if err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}
	
	resp.fillToPayment(&payment, cpID, exID, prID)

	err = p.db.PaymentCreate(&payment)
	if err != nil {
		return 0, err
	}
	return *payment.ID, nil
}

func NewPaymentServices(db db.DB, counterparty Counterparty, expenditureItem ExpenditureItem, projcet Project) Payment {
	return &PaymentService{db: db, counterparty: counterparty, expenditureItem: expenditureItem, project: projcet}
}
