package service

import "app/internal/db"

type ExpenditureItem interface {
	ExpenditureItemCreate(*string) (int, error)
	ExpenditureFindeByName(name *string) (int, error)
}

type expenditureServices struct {
	db db.DB
}

func (exp *expenditureServices) ExpenditureItemCreate(name *string) (int, error) {
	var ex db.ExpenditureItem
	ex.Name = name
	err := exp.db.ExpenditureItemFindeByName(&ex)
	if err == nil {
		return *ex.ID, nil
	}
	err = exp.db.ExpenditureItemCreate(&ex)
	if err != nil {
		return 0, err
	}
	return *ex.ID, nil
}
func (exp *expenditureServices) ExpenditureFindeByName(name *string) (int, error) {
	var ex db.ExpenditureItem
	ex.Name = name
	err := exp.db.ExpenditureItemFindeByName(&ex)
	if err != nil {
		return 0, err
	}
	return *ex.ID, nil
}

func NewExpenditureServices(db db.DB) ExpenditureItem {
	return &expenditureServices{db: db}
}
