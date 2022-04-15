package service

import "app/pkg/db"

type Payment interface {
}

type Expenditure interface {
}

type Service struct {
	Payment
	Expenditure
}

func NewService(db db.DB) *Service {
	return &Service{}
}
