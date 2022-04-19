package service

import (
	"app/internal/db"

	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type ExpenditureService interface {
	ExpenditureCreate(*ExpendaturesFullRequest, *logrus.Logger) (int, error)
}

type ExpendaturesFullRequest struct {
	CounterpartyName    string  `json:"counterpartyName"`
	ConditionOne        string  `jspm:"conditionOne"`
	ProjectName         string  `json:"projectName"`
	ConditionTwo        string  `jspm:"conditionTwo"`
	Comments            string  `json:"comments"`
	ConditionThee       string  `jspm:"conditionThree"`
	Purpose             string  `json:"purpose"`
	ExpenditureItemName *string `json:"expenditureItmeName"`
}

func (r *ExpendaturesFullRequest) fillToExpendatures(e *db.Expenditures, cpID, exID, prID int) {
	e.CounterpartyID = cpID
	e.ProjectID = prID
	e.ExpenditureItemID = &exID
	e.Comments = r.Comments
	e.ConditionOne = r.ConditionOne
	e.ConditionTwo = r.ConditionTwo
	e.ConditionThee = r.ConditionThee
	e.Purpose = r.Purpose
}

func (p *PaymentService) ExpenditureCreate(resp *ExpendaturesFullRequest, log *logrus.Logger) (int, error) {
	var e db.Expenditures

	cpID, err := p.counterparty.CounterpartyFindeByName(&resp.CounterpartyName, log)
	if err == pgx.ErrNoRows {
		cpID, err = p.counterparty.CounterpartyCreate(&resp.CounterpartyName, log)
		if err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}

	exID, err := p.expenditureItem.ExpenditureItemCreate(resp.ExpenditureItemName, log)
	if err == pgx.ErrNoRows {
		exID, err = p.expenditureItem.ExpenditureItemCreate(resp.ExpenditureItemName, log)
		if err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}

	prID, err := p.project.ProjectCreate(&resp.ProjectName, log)
	if err == pgx.ErrNoRows {
		prID, err = p.project.ProjectCreate(&resp.ProjectName, log)
		if err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}

	resp.fillToExpendatures(&e, cpID, exID, prID)
	err = p.db.ExpenditureCreate(&e, log)
	if err != nil {
		return 0, err
	}
	return *e.ID, nil
}
