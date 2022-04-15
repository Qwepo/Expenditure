package db

import (
	"context"
	"fmt"
)

const (
	paymentTable       = "payment"
	counterpartysTable = "counterpartys"
	organizationsTable = "organizations"
)

type dbPayment interface {
	OrganizationCreate(*Organizations) (int64 error)
	CounterpartyCreate(*Counterparty) (int64 error)
	PaymentCreate(*Payment) (int64 error)
}

type Payment struct {
	ID                 int64  `json:"id"`
	Doctype            string `json:"doctype"`
	Time               int64  `json:"time"`
	OrganizationID     int64  `json:"organizationId"`
	CounterpartyID     int64  `json:"counterpartyId"`
	IncomingCurrency   int64  `json:"incomingCurrency"`
	ExpendableCurrency int64  `json:"expendableCurrency"`
	Purpose            string `json:"purpose"`
	Expenditure        string `json:"expenditure"`
	Comments           string `json:"comments"`
}

type Organizations struct {
	ID   int64  `json:"id"`
	Name string `json:"name" `
}
type Counterparty struct {
	ID   int64  `json:"id"`
	Name string `json:"name" `
}

func (db *clietn) OrganizationCreate(org *Organizations) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s name VALUES $1", organizationsTable)
	row := db.QueryRow(context.TODO(), query, org.Name)
	if err := row.Scan(org.ID); err != nil {
		return 0, err
	}
	return org.ID, nil
}
func (db *clietn) CounterpartyCreate(c *Counterparty) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s name VALUES $1", counterpartysTable)
	row := db.QueryRow(context.TODO(), query, c.Name)
	if err := row.Scan(c.ID); err != nil {
		return 0, err
	}
	return c.ID, nil
}

func (db *clietn) PaymentCreate(p *Payment) (int64, error) {
	query := fmt.Sprintf("INSERT INTO %s (document_type, time, organization_id, counterparty_id, incoming_currency, expendable_currency, purpose, expenditure, comments) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id", paymentTable)
	row := db.QueryRow(context.TODO(), query, p.Doctype, p.Time, p.OrganizationID, p.CounterpartyID, p.IncomingCurrency, p.ExpendableCurrency, p.Purpose, p.Expenditure, p.Comments)
	if err := row.Scan(p.ID); err != nil {
		return 0, err
	}
	return p.ID, nil
}
