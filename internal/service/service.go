package service

import (
	"app/internal/db"

	"github.com/sirupsen/logrus"
)

type Service struct {
	ExpenditureItem ExpenditureItem
	Counterparty    Counterparty
	Project         Project
	Payment         Payment
}

func NewService(db db.DB, log *logrus.Logger) *Service {
	ex := NewExpenditureServices(db, log)
	cp := NewCounterpartyServices(db, log)
	pr := NewProjectServices(db, log)
	return &Service{
		ExpenditureItem: ex,
		Counterparty:    cp,
		Project:         pr,
		Payment:         NewPaymentServices(db, log, cp, ex, pr),
	}
}
