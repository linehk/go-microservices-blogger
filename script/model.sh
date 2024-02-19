#!/bin/bash

set -e

service_name="$1"
table_name="$2"

goctl model pg datasource \
    --url="postgresql://dev:dev@localhost:5432/go_blogger?sslmode=disable" \
    --table="$table_name" -c -dir ./service/"$service_name"/rpc/model/ \
    -style go_zero