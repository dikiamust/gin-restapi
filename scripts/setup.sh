#!/bin/bash
export $(cat .env | xargs)
migrate -database "${DATABASE_URL}" -path ./migrations up
