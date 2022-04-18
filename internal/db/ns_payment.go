package db

import (
	"context"
	"fmt"
	"time"
)

const (
	paymentTable       = "payment"
	counterpartysTable = "counterpartys"
	organizationsTable = "organizations"
	projectTable       = "project"
)

type dbPayment interface {
	OrganizationCreate(*Organizations) (int64, error)
	CounterpartyCreate(*Counterparty) (int64, error)
	ProjectCreate(*Project) (int64, error)
	PaymentCreate(*Payment) (int64, error)
}

type Payment struct {
	ID                 *int64     `json:"id"`
	Doctype            *string    `json:"doctype"`
	CreatedAt          *time.Time `json:"createdAt"`
	OrganizationID     *int64     `json:"organizationId"`
	CounterpartyID     *int64     `json:"counterpartyId"`
	IncomingCurrency   *int64     `json:"incomingCurrency"`
	ExpendableCurrency *int64     `json:"expendableCurrency"`
	Purpose            *string    `json:"purpose"`
	Expenditure        *string    `json:"expenditure"`
	ProjectID          *int64     `json:"projectId"`
	Comments           string     `json:"comments"`
}

type Organizations struct {
	ID   *int64  `json:"id"`
	Name *string `json:"name"`
}
type Counterparty struct {
	ID   *int64  `json:"id"`
	Name *string `json:"name"`
}
type Project struct {
	ID   *int64  `json:"id"`
	Name *string `json:"name"`
}

func (db *clietn) OrganizationCreate(org *Organizations) (int64, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", organizationsTable)
	row := db.QueryRow(context.TODO(), query, org.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	id64 := int64(id)
	org.ID = &id64
	return *org.ID, nil
}
func (db *clietn) CounterpartyCreate(c *Counterparty) (int64, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", counterpartysTable)
	row := db.QueryRow(context.TODO(), query, c.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	id64 := int64(id)

	c.ID = &id64
	return *c.ID, nil
}

func (db *clietn) ProjectCreate(pr *Project) (int64, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", projectTable)
	row := db.QueryRow(context.TODO(), query, pr.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	id64 := int64(id)
	pr.ID = &id64
	return *pr.ID, nil
}

func (db *clietn) PaymentCreate(p *Payment) (int64, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (document_type, time, organization_id, counterparty_id, incoming_currency, expendable_currency, purpose, expenditure, project_id, comments) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id", paymentTable)
	row := db.QueryRow(context.TODO(), query, p.Doctype, p.CreatedAt, p.OrganizationID, p.CounterpartyID, p.IncomingCurrency, p.ExpendableCurrency, p.Purpose, p.Expenditure, p.ProjectID, p.Comments)
	err := row.Scan(&id)
	if err != nil {

		return 0, err
	}
	id64 := int64(id)
	p.ID = &id64
	return *p.ID, nil
}
