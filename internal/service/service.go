package service

import "app/pkg/db"

type Service struct {
	Payment       Payment
	Organizations Organizations
}

func NewService(db db.DB) *Service {
	return &Service{
		Payment:       NewPaymentServices(db),
		Organizations: NewOrganizationServices(db),
	}
}
