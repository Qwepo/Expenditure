package service

import (
	"app/internal/db"

	"github.com/sirupsen/logrus"
)

type ExpenditureItem interface {
	ExpenditureItemCreate(*string, *logrus.Logger) (int, error)
	ExpenditureFindeByName(*string, *logrus.Logger) (int, error)
}

type expenditureServices struct {
	db  db.DB
	log *logrus.Logger
}

func (exp *expenditureServices) ExpenditureItemCreate(name *string, log *logrus.Logger) (int, error) {
	var ex db.ExpenditureItem
	ex.Name = name
	err := exp.db.ExpenditureItemFindeByName(&ex, log)
	if err == nil {
		return *ex.ID, nil
	}
	err = exp.db.ExpenditureItemCreate(&ex, log)
	if err != nil {
		return 0, err
	}
	return *ex.ID, nil
}
func (exp *expenditureServices) ExpenditureFindeByName(name *string, log *logrus.Logger) (int, error) {
	var ex db.ExpenditureItem
	ex.Name = name
	err := exp.db.ExpenditureItemFindeByName(&ex, log)
	if err != nil {
		return 0, err
	}
	return *ex.ID, nil
}

func NewExpenditureServices(db db.DB, log *logrus.Logger) ExpenditureItem {
	return &expenditureServices{db: db, log: log}
}
