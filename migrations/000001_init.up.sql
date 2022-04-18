CREATE TABLE organizations
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL
);

CREATE TABLE counterpartys
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL
);

CREATE TABLE project
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL
);

CREATE TABLE payment
 (
    id SERIAL PRIMARY KEY,
    document_type VARCHAR(150) NOT NULL,
    time TIMESTAMP NOT NULL,
    organization_id INTEGER  REFERENCES organizations(id)  NOT NULL,
    counterparty_id INTEGER REFERENCES counterpartys(id) NOT NULL,
    incoming_currency NUMERIC NOT NULL,
    expendable_currency NUMERIC NOT NULL,
    purpose VARCHAR(150) NOT NULL,
    expenditure VARCHAR(150) NOT NULL,
    project_id INTEGER REFERENCES project(id) NOT NULL,
    comments VARCHAR(150)
);