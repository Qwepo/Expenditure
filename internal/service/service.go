package service

import "app/internal/db"

type Service struct {
	Organizations Organizations
	Counterparty  Counterparty
	Payment       Payment
}

func NewService(db db.DB) *Service {
	return &Service{
		Organizations: NewOrganizationServices(db),
		Counterparty: NewCounterpartyServices(db),
		Payment:       NewPaymentServices(db),
	}
}
