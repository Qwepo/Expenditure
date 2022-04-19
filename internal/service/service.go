package service

import "app/internal/db"

type Service struct {
	ExpenditureItem ExpenditureItem
	Counterparty    Counterparty
	Project         Project
	Payment         Payment
}

func NewService(db db.DB) *Service {
	ex := NewExpenditureServices(db)
	cp := NewCounterpartyServices(db)
	pr := NewProjectServices(db)
	return &Service{
		ExpenditureItem: ex,
		Counterparty:    cp,
		Project:         pr,
		Payment:         NewPaymentServices(db, cp, ex, pr),
	}
}
