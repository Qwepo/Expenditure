package service

import (
	"app/internal/db"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type Payment interface {
	PaymentCreate(*PaymentFullRequest, *logrus.Logger) (int, error)
}

type PaymentService struct {
	db              db.DB
	log             *logrus.Logger
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

func (p *PaymentService) PaymentCreate(resp *PaymentFullRequest, log *logrus.Logger) (int, error) {
	var payment db.Payment
	cpID, err := p.counterparty.CounterpartyFindeByName(resp.CounterpartyName, log)
	if err == pgx.ErrNoRows {
		cpID, err = p.counterparty.CounterpartyCreate(resp.CounterpartyName, log)
		if err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}

	exID, err := p.expenditureItem.ExpenditureItemCreate(resp.ExpenditureItemName, log)
	if err == pgx.ErrNoRows {
		exID, err = p.expenditureItem.ExpenditureItemCreate(resp.ExpenditureItemName, log)
		if err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}

	prID, err := p.project.ProjectCreate(resp.ProjectName, log)
	if err == pgx.ErrNoRows {
		prID, err = p.project.ProjectCreate(resp.ProjectName, log)
		if err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}

	resp.fillToPayment(&payment, cpID, exID, prID)

	err = p.db.PaymentCreate(&payment, log)
	if err != nil {
		return 0, err
	}
	return *payment.ID, nil
}

func NewPaymentServices(db db.DB, log *logrus.Logger, counterparty Counterparty, expenditureItem ExpenditureItem, projcet Project) Payment {
	return &PaymentService{db: db, log: log, counterparty: counterparty, expenditureItem: expenditureItem, project: projcet}
}
