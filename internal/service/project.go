package service

import (
	"app/internal/db"
)

type Project interface {
	ProjectCreate(*PaymentFullRequest) (int64, error)
}

type projectServices struct {
	db db.DB
}

func (pro *projectServices) ProjectCreate(resp *PaymentFullRequest) (int64, error) {
	var pr db.Project

	pr.Name = resp.ProjectName
	id, err := pro.db.ProjectCreate(&pr)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func NewProjectServices(db db.DB) Project {
	return &projectServices{db: db}
}
