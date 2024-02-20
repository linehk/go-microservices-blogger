#!/bin/bash

set -e

for service_name in user blog post page comment
do
    CGO_ENABLED=0 GOOS=linux go build \
        -o ./bin/ \
        ./service/"$service_name"/rpc/"$service_name".go
done