#!/bin/bash

set -e

for service_name in user blog post page comment
do
    ./bin/"$service_name" -f ./service/"$service_name"/rpc/etc/"$service_name".yaml &
done