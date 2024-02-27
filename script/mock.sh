#!/bin/bash

set -e

for service_name in user blog post page comment
do
  mockgen -source=./service/"$service_name"/rpc/"$service_name"service/"$service_name"_service.go \
    -destination=./service/"$service_name"/rpc/"$service_name"service/mock_"$service_name"_service.go \
    -package="$service_name"service \
    -self_package=github.com/linehk/go-microservices-blogger/service/"$service_name"/rpc/"$service_name"service \
    github.com/linehk/go-microservices-blogger/service/"$service_name"/rpc/"$service_name"service
done