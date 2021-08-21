#!/bin/bash

source .env

docker-compose run --rm migrate -path ./schema -database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" up
