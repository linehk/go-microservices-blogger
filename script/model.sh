#!/bin/bash

set -e

name="$1"
table_name="$1"

if [ "$1" == "user" ]
then table_name="app_user"
fi

goctl model pg datasource \
    --url="postgresql://dev:dev@localhost:5432/go_blogger?sslmode=disable" \
    --table="$table_name" -c -dir ./service/"$name"/rpc/model/ \
    -style go_zero