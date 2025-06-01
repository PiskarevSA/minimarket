#!/usr/bin/bash
sqlc generate -f services/accrual/.sqlc.yaml
sqlc generate -f services/gophermart/.sqlc.yaml