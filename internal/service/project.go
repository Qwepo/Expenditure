package service

import (
	"app/internal/db"
)

type Project interface {
	ProjectCreate(*string) (int, error)
	ProjectFindeByName(name *string) (int, error)
}

type projectServices struct {
	db db.DB
}

func (pro *projectServices) ProjectCreate(name *string) (int, error) {
	var pr db.Project
	pr.Name = name
	err := pro.db.ProjectFindeByName(&pr)
	if err == nil {
		return *pr.ID, nil
	}
	err = pro.db.ProjectCreate(&pr)
	if err != nil {
		return 0, err
	}
	return *pr.ID, nil
}

func (pro *projectServices) ProjectFindeByName(name *string) (int, error) {
	var pr db.Project
	pr.Name = name
	err := pro.db.ProjectFindeByName(&pr)
	if err != nil {
		return 0, err
	}
	return *pr.ID, nil
}

func NewProjectServices(db db.DB) Project {
	return &projectServices{db: db}
}
