#!/bin/bash

set -e

readonly name="$1"

goctl rpc protoc \
		./service/"$name"/rpc/"$name".proto \
		--go_out=./service/"$name"/rpc/ \
		--go-grpc_out=./service/"$name"/rpc/ \
		--zrpc_out=./service/"$name"/rpc/ \
		-style go_zero