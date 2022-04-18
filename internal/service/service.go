package service

import "app/internal/db"

type Service struct {
	Organizations Organizations
	Counterparty  Counterparty
	Project       Project
	Payment       Payment
}

func NewService(db db.DB) *Service {
	org := NewOrganizationServices(db)
	cp := NewCounterpartyServices(db)
	pr := NewProjectServices(db)
	return &Service{
		Organizations: org,
		Counterparty:  cp,
		Project:       pr,
		Payment:       NewPaymentServices(db, cp, org, pr),
	}
}
