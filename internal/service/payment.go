package service

import (
	"app/internal/db"
)

type Payment interface {
	PaymentCreate(*PaymentFullRequest) (int64, error)
}

type PaymentService struct {
	db            db.DB
	counterparty  *Counterparty
	organizaotion *Organizations
}

type PaymentFullRequest struct {
	Doctype            string `json:"doctype"`
	OrganizationName   string `json:"organizationName"`
	CounterpartyName   string `json:"counterpartyName"`
	OrganizationID     int64  `json:"organizationId"`
	CounterpartyID     int64  `json:"counterpartyId"`
	IncomingCurrency   int64  `json:"incomingCurrency"`
	ExpendableCurrency int64  `json:"expendableCurrency"`
	Purpose            string `json:"purpose"`
	Expenditure        string `json:"expenditure"`
	Comments           string `json:"comments"`
}

func (r *PaymentFullRequest) fillTo(p *db.Payment) {
	r.Doctype = p.Doctype
	r.OrganizationID = p.OrganizationID
	r.CounterpartyID = p.CounterpartyID
	r.IncomingCurrency = p.IncomingCurrency
	r.ExpendableCurrency = p.ExpendableCurrency
	r.Purpose = p.Purpose
	r.Expenditure = p.Expenditure
	r.Comments = p.Comments
}

func (p *PaymentService) PaymentCreate(resp *PaymentFullRequest) (int64, error) {
	var payment db.Payment
	resp.fillTo(&payment)
	id, err := p.db.PaymentCreate(&payment)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func NewPaymentServices(db db.DB, counterparty Counterparty, organizaotion Organizations) Payment {
	return &PaymentService{db: db, counterparty: &counterparty, organizaotion: &organizaotion}
}
