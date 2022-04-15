package service

import "app/internal/db"

type Counterparty interface {
	CounterpartyCreate(*CounterpartyRequest) (int64, error)
}

type сounterpartyServices struct {
	db db.DB
}

type CounterpartyRequest struct {
	Name string `json:"name"`
}

func (c *сounterpartyServices) CounterpartyCreate(resp *CounterpartyRequest) (int64, error) {
	var cp db.Counterparty
	cp.Name = resp.Name
	id, err := c.db.CounterpartyCreate(&cp)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func NewCounterpartyServices(db db.DB) Counterparty {
	return &сounterpartyServices{db: db}
}
