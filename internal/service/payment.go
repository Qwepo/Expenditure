package service

import (
	"app/internal/db"
	"time"
)

type Payment interface {
	PaymentCreate(*PaymentFullRequest) (int64, error)
}

type PaymentService struct {
	db            db.DB
	counterparty  Counterparty
	organizaotion Organizations
	project       Project
}

type PaymentFullRequest struct {
	Doctype            *string `json:"doctype"`
	OrganizationName   *string `json:"organizationName"`
	CounterpartyName   *string `json:"counterpartyName"`
	IncomingCurrency   *int64  `json:"incomingCurrency"`
	ExpendableCurrency *int64  `json:"expendableCurrency"`
	Purpose            *string `json:"purpose"`
	Expenditure        *string `json:"expenditure"`
	ProjectName        *string `json:"projectName"`
	Comments           string  `json:"comments"`
}

func (r *PaymentFullRequest) fillTo(p *db.Payment) {
	p.Doctype = r.Doctype
	p.IncomingCurrency = r.IncomingCurrency
	p.ExpendableCurrency = r.ExpendableCurrency
	p.Purpose = r.Purpose
	p.Expenditure = r.Expenditure
	p.Comments = r.Comments
}

func (p *PaymentService) PaymentCreate(resp *PaymentFullRequest) (int64, error) {
	var payment db.Payment
	time := time.Now()
	cpID, err := p.counterparty.CounterpartyCreate(resp)
	if err != nil {
		return 0, err
	}
	orgID, err := p.organizaotion.OrganizationCreate(resp)
	if err != nil {
		return 0, err
	}

	prID, err := p.project.ProjectCreate(resp)
	if err != nil {
		return 0, err
	}
	payment.CounterpartyID = &cpID
	payment.OrganizationID = &orgID
	payment.ProjectID = &prID
	payment.CreatedAt = &time
	resp.fillTo(&payment)
	id, err := p.db.PaymentCreate(&payment)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func NewPaymentServices(db db.DB, counterparty Counterparty, organizaotion Organizations, projcet Project) Payment {
	return &PaymentService{db: db, counterparty: counterparty, organizaotion: organizaotion, project: projcet}
}
