CREATE TABLE if not exists imports.records
(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    county VARCHAR(300),
    sex VARCHAR(100),
    age VARCHAR(50),
    onset_date TIMESTAMP,
    death_date TIMESTAMP,
    admission_date TIMESTAMP,
    case_count INT,
    death_count INT,
    hospitalized_count INT
);


alter table imports.records
    owner to covid_county;
