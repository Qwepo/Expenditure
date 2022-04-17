package service

import "app/internal/db"

type Service struct {
	Organizations Organizations
	Counterparty  Counterparty
	Payment       Payment
}

func NewService(db db.DB) *Service {
	org := NewOrganizationServices(db)
	cp := NewCounterpartyServices(db)
	return &Service{
		Organizations: org,
		Counterparty:  cp,
		Payment:       NewPaymentServices(db, cp, org),
	}
}
