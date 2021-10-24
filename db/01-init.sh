#!/bin/bash

export PGPASSWORD=$POSTGRES_PASSWORD;

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE USER $REMINDME_DB_USER WITH PASSWORD '$REMINDME_DB_PASSWORD';
  CREATE DATABASE $REMINDME_DB_NAME;
  GRANT ALL PRIVILEGES ON DATABASE $REMINDME_DB_NAME TO $REMINDME_DB_USER;
  \connect remindme;
  CREATE TABLE IF NOT EXISTS reminder (
    id serial PRIMARY KEY,
    chat_id INT NOT NULL,
    chat_message VARCHAR NOT NULL,
    is_recurrent BOOLEAN DEFAULT FALSE,
    is_everyday BOOLEAN DEFAULT FALSE,
    weekly_day VARCHAR(10) DEFAULT NULL,
    monthly_day INT DEFAULT NULL,
    annual_date VARCHAR(30) DEFAULT NULL,
    unique_date date,
    unique_time time NOT NULL,
    CONSTRAINT valid_month_date CHECK (monthly_day BETWEEN 1 AND 31)
  );
EOSQL