package service

import "app/internal/db"

type Organizations interface {
	OrganizationCreate(*OrganizationRequest) (int64, error)
}

type organizationServices struct {
	db db.DB
}

type OrganizationRequest struct {
	Name string `json:"name"`
}

func (org *organizationServices) OrganizationCreate(resp *OrganizationRequest) (int64, error) {
	var o db.Organizations
	o.Name = resp.Name
	id, err := org.db.OrganizationCreate(&o)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func NewOrganizationServices(db db.DB) Organizations {
	return &organizationServices{db: db}
}
