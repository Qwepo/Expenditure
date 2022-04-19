package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/sirupsen/logrus"
)

type dbExpenditures interface {
	ExpenditureCreate(*Expenditures, *logrus.Logger) error
}

type Expenditures struct {
	ID                *int   `json:"id"`
	CounterpartyID    int    `json:"counterpartyId"`
	ConditionOne      string `jspm:"conditionOne"`
	ProjectID         int    `json:"projectId"`
	ConditionTwo      string `jspm:"conditionTwo"`
	Comments          string `json:"comments"`
	ConditionThee     string `jspm:"conditionThree"`
	Purpose           string `json:"purpose"`
	ExpenditureItemID *int   `json:"expenditureItmeId"`
}

func (db *clietn) ExpenditureCreate(ex *Expenditures, log *logrus.Logger) error {
	query := fmt.Sprintf("INSERT INTO %s (counterparty_id, condition_one, project_id, condition_two, comments, condition_three, purpose, expenditure_id ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", expenditureItemTable)
	row := db.QueryRow(context.TODO(), query, ex.CounterpartyID, ex.ConditionOne, ex.ProjectID, ex, ex.ConditionTwo, ex.Comments, ex.ConditionThee, ex.Purpose, ex.ExpenditureItemID)
	if err := row.Scan(&ex.ID); err != nil  {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			newErr := fmt.Errorf("SQL Error: %s, Deatil: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState())
			log.Panic(newErr)
			return nil
		}
		return err
	}
	return nil
}
