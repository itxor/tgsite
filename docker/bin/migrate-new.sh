#!/bin/bash

source .env

docker-compose run --rm -u www-data migrate create -ext js -dir ./schema -seq "$1"
