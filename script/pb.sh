#!/bin/bash

set -e

for service_name in user blog post page comment
do
  goctl rpc protoc \
    ./service/"$service_name"/rpc/"$service_name".proto \
    --go_out=./service/"$service_name"/rpc/ \
    --go-grpc_out=./service/"$service_name"/rpc/ \
    --zrpc_out=./service/"$service_name"/rpc/ \
    -style go_zero
done