package service

import (
	"app/internal/db"

	"github.com/sirupsen/logrus"
)

type Project interface {
	ProjectCreate(*string, *logrus.Logger) (int, error)
	ProjectFindeByName(*string, *logrus.Logger) (int, error)
}

type projectServices struct {
	log *logrus.Logger
	db  db.DB
}

func (pro *projectServices) ProjectCreate(name *string, log *logrus.Logger) (int, error) {
	var pr db.Project
	pr.Name = name
	err := pro.db.ProjectFindeByName(&pr, log)
	if err == nil {
		return *pr.ID, nil
	}
	err = pro.db.ProjectCreate(&pr, log)
	if err != nil {
		return 0, err
	}
	return *pr.ID, nil
}

func (pro *projectServices) ProjectFindeByName(name *string, log *logrus.Logger) (int, error) {
	var pr db.Project
	pr.Name = name
	err := pro.db.ProjectFindeByName(&pr, log)
	if err != nil {
		return 0, err
	}
	return *pr.ID, nil
}

func NewProjectServices(db db.DB, log *logrus.Logger) Project {
	return &projectServices{db: db, log: log}
}
