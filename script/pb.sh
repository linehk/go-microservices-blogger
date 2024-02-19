#!/bin/bash

set -e

readonly service_name="$1"

goctl rpc protoc \
		./service/"$service_name"/rpc/"$service_name".proto \
		--go_out=./service/"$service_name"/rpc/ \
		--go-grpc_out=./service/"$service_name"/rpc/ \
		--zrpc_out=./service/"$service_name"/rpc/ \
		-style go_zero