package service

import "app/internal/db"

type Organizations interface {
	OrganizationCreate(*PaymentFullRequest) (int64, error)
}

type organizationServices struct {
	db db.DB
}

func (org *organizationServices) OrganizationCreate(resp *PaymentFullRequest) (int64, error) {
	var o db.Organizations
	o.Name = resp.OrganizationName
	id, err := org.db.OrganizationCreate(&o)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func NewOrganizationServices(db db.DB) Organizations {
	return &organizationServices{db: db}
}
