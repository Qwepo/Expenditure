CREATE TABLE expenditure
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
    organization VARCHAR(150) NOT NULL,
    counterparty_id INTEGER REFERENCES counterpartys(id) NOT NULL,
    incoming_currency NUMERIC NOT NULL,
    expendable_currency NUMERIC NOT NULL,
    purpose VARCHAR(150) NOT NULL,
    expenditure_id INTEGER REFERENCES expenditure(id) NOT NULL,
    project_id INTEGER REFERENCES project(id) NOT NULL,
    comments VARCHAR(150)
);

CREATE TABLE article_of_expenditures
(
    id SERIAL PRIMARY KEY,
    counterparty_id INTEGER REFERENCES counterpartys(id),
    condition_one VARCHAR(100),
    project_id INTEGER REFERENCES project(id),
    condition_two VARCHAR(100),
    comments VARCHAR(150),
    condition_three VARCHAR(100),
    purpose VARCHAR(150),
    expenditure_id INTEGER REFERENCES expenditure(id) NOT NULL
);