package service

import (
	"app/internal/db"

	"github.com/sirupsen/logrus"
)

type Counterparty interface {
	CounterpartyCreate(*string, *logrus.Logger) (int, error)
	CounterpartyFindeByName(*string, *logrus.Logger) (int, error)
}

type сounterpartyServices struct {
	db db.DB
	log *logrus.Logger
}

func (c *сounterpartyServices) CounterpartyCreate(name *string, log *logrus.Logger) (int, error) {
	var cp db.Counterparty
	cp.Name = name
	err := c.db.CounterpartyFindeByName(&cp, log)
	if err == nil {
		return *cp.ID, nil
	}

	err = c.db.CounterpartyCreate(&cp, log)
	if err != nil {
		return 0, err
	}
	return *cp.ID, nil
}

func (c *сounterpartyServices) CounterpartyFindeByName(name *string, log *logrus.Logger) (int, error) {
	var cp db.Counterparty
	cp.Name = name
	err := c.db.CounterpartyFindeByName(&cp, log)
	if err != nil {
		return 0, err
	}
	return *cp.ID, nil
}

func NewCounterpartyServices(db db.DB, log *logrus.Logger) Counterparty {
	return &сounterpartyServices{db: db, log: log}
}
