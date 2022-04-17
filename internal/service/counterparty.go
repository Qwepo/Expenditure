package service

import "app/internal/db"

type Counterparty interface {
	CounterpartyCreate(*PaymentFullRequest) (int64, error)
}

type сounterpartyServices struct {
	db db.DB
}

func (c *сounterpartyServices) CounterpartyCreate(resp *PaymentFullRequest) (int64, error) {
	var cp db.Counterparty
	cp.Name = resp.CounterpartyName
	id, err := c.db.CounterpartyCreate(&cp)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func NewCounterpartyServices(db db.DB) Counterparty {
	return &сounterpartyServices{db: db}
}
