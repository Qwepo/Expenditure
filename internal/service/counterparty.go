package service

import "app/internal/db"

type Counterparty interface {
	CounterpartyCreate(*string) (int, error)
	CounterpartyFindeByName(name *string) (int, error)
}

type сounterpartyServices struct {
	db db.DB
}

func (c *сounterpartyServices) CounterpartyCreate(name *string) (int, error) {
	var cp db.Counterparty
	cp.Name = name
	err := c.db.CounterpartyFindeByName(&cp)
	if err == nil {
		return *cp.ID, nil
	}

	err = c.db.CounterpartyCreate(&cp)
	if err != nil {
		return 0, err
	}
	return *cp.ID, nil
}

func (c *сounterpartyServices) CounterpartyFindeByName(name *string) (int, error) {
	var cp db.Counterparty
	cp.Name = name
	err := c.db.CounterpartyFindeByName(&cp)
	if err != nil {
		return 0, err
	}
	return *cp.ID, nil
}

func NewCounterpartyServices(db db.DB) Counterparty {
	return &сounterpartyServices{db: db}
}
