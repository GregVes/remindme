#!/bin/bash

export PGPASSWORD=$POSTGRES_PASSWORD;

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE USER $REMINDME_DB_USER WITH PASSWORD '$REMINDME_DB_PASSWORD';
  CREATE DATABASE $REMINDME_DB_NAME;
  GRANT ALL PRIVILEGES ON DATABASE $REMINDME_DB_NAME TO $REMINDME_DB_USER;
EOSQL