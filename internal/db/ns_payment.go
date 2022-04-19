package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

const (
	paymentTable         = "payment"
	counterpartysTable   = "counterpartys"
	expenditureItemTable = "expenditure"
	projectTable         = "project"
)

type dbPayment interface {
	ExpenditureItemCreate(*ExpenditureItem, *logrus.Logger) error
	ExpenditureItemFindeByName(*ExpenditureItem, *logrus.Logger) error
	CounterpartyCreate(*Counterparty, *logrus.Logger) error
	CounterpartyFindeByName(*Counterparty, *logrus.Logger) error
	ProjectCreate(*Project, *logrus.Logger) error
	ProjectFindeByName(*Project, *logrus.Logger) error
	PaymentCreate(*Payment, *logrus.Logger) error
}

type Payment struct {
	ID                 *int       `json:"id"`
	Doctype            *string    `json:"doctype"`
	CreatedAt          *time.Time `json:"createdAt"`
	OrganizationName   *string    `json:"OrganizationeName"`
	CounterpartyID     *int       `json:"counterpartyId"`
	IncomingCurrency   *int       `json:"incomingCurrency"`
	ExpendableCurrency *int       `json:"expendableCurrency"`
	Purpose            *string    `json:"purpose"`
	ExpenditureItemID  *int       `json:"expenditureItmeId"`
	ProjectID          *int       `json:"projectId"`
	Comments           string     `json:"comments"`
}

type ExpenditureItem struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}
type Counterparty struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}
type Project struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

func (db *clietn) ExpenditureItemCreate(ex *ExpenditureItem, log *logrus.Logger) error {
	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", expenditureItemTable)
	row := db.QueryRow(context.TODO(), query, ex.Name)
	if err := row.Scan(&ex.ID); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			fmt.Println(pgErr)
			if pgErr == pgx.ErrNoRows {
				return pgErr
			}
			newErr := fmt.Errorf("SQL Error: %s, Deatil: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			log.Panic(newErr)
			return nil
		}
		return err
	}
	return nil
}
func (db *clietn) ExpenditureItemFindeByName(ex *ExpenditureItem, log *logrus.Logger) error {
	query := fmt.Sprintf("SELECT id FROM %s WHERE name = $1", expenditureItemTable)
	row := db.QueryRow(context.TODO(), query, ex.Name)
	if err := row.Scan(&ex.ID); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr == pgx.ErrNoRows {
				return pgErr
			}
			newErr := fmt.Errorf("aSQL Error: %s, Deatil: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			log.Panic(newErr)
			return nil
		}
		return err
	}
	return nil
}

func (db *clietn) CounterpartyFindeByName(c *Counterparty, log *logrus.Logger) error {
	query := fmt.Sprintf("SELECT id FROM %s WHERE name = $1", counterpartysTable)
	row := db.QueryRow(context.TODO(), query, c.Name)
	if err := row.Scan(&c.ID); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr == pgx.ErrNoRows {
				return pgErr
			}
			newErr := fmt.Errorf("bSQL Error: %s, Deatil: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			log.Panic(newErr)
			return nil
		}
		return err
	}
	return nil
}

func (db *clietn) CounterpartyCreate(c *Counterparty, log *logrus.Logger) error {
	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", counterpartysTable)
	row := db.QueryRow(context.TODO(), query, c.Name)
	if err := row.Scan(&c.ID); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr == pgx.ErrNoRows {
				return pgErr
			}
			newErr := fmt.Errorf("cSQL Error: %s, Deatil: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			log.Panic(newErr)
			return nil
		}
		return err
	}
	return nil
}
func (db *clietn) ProjectFindeByName(pr *Project, log *logrus.Logger) error {
	query := fmt.Sprintf("SELECT id FROM %s WHERE name = $1", projectTable)
	row := db.QueryRow(context.TODO(), query, pr.Name)
	if err := row.Scan(&pr.ID); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr == pgx.ErrNoRows {
				return pgErr
			}
			newErr := fmt.Errorf("fSQL Error: %s, Deatil: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			log.Panic(newErr)
			return nil
		}
		return err
	}
	return nil
}

func (db *clietn) ProjectCreate(pr *Project, log *logrus.Logger) error {
	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", projectTable)
	row := db.QueryRow(context.TODO(), query, pr.Name)
	if err := row.Scan(&pr.ID); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr == pgx.ErrNoRows {
				return pgErr
			}
			newErr := fmt.Errorf("gSQL Error: %s, Deatil: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			log.Panic(newErr)
			return nil
		}
		return err
	}
	return nil
}

func (db *clietn) PaymentCreate(p *Payment, log *logrus.Logger) error {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (document_type, time, organization, counterparty_id, incoming_currency, expendable_currency, purpose, expenditure_id, project_id, comments) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id", paymentTable)
	row := db.QueryRow(context.TODO(), query, p.Doctype, p.CreatedAt, p.OrganizationName, p.CounterpartyID, p.IncomingCurrency, p.ExpendableCurrency, p.Purpose, p.ExpenditureItemID, p.ProjectID, p.Comments)
	err := row.Scan(&id)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgErr == pgx.ErrNoRows {
				return pgErr
			}
			newErr := fmt.Errorf("zSQL Error: %s, Deatil: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			log.Panic(newErr)
			return nil
		}
		return err
	}
	return nil
}
