package service

import "app/pkg/db"

type Organizations interface {
	OrganizationCreate(*OrganizationRequest) (int64, error)
}

type OrganizationServices struct {
	db db.DB
}

type OrganizationRequest struct {
	Name string `json:"name"`
}

func (org *OrganizationServices) OrganizationCreate(resp *OrganizationRequest) (int64, error) {
	var o db.Organizations
	o.Name = resp.Name
	id, err := org.db.OrganizationCreate(&o)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func NewOrganizationServices(db db.DB) Organizations {
	return &OrganizationServices{db: db}
}
