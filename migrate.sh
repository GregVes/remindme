#!/bin/bash


migrate -path ./db/migration -database "postgresql://${REMINDME_DB_USER}:${REMINDME_DB_PASSWORD}@localhost:${REMINDME_DB_PORT}/${REMINDME_DB_NAME}?sslmode=disable" -verbose up
