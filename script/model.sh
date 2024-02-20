#!/bin/bash

set -e

gen_model_by_table_and_service() {
  goctl model pg datasource \
    --url="postgresql://dev:dev@localhost:5432/go_blogger?sslmode=disable" \
    --table="$1" -c -dir ./service/"$2"/rpc/model/ \
    -style go_zero
}

for table_name in app_user locale
do
  gen_model_by_table_and_service $table_name user
done

for table_name in blog blog_user_info page_views
do
  gen_model_by_table_and_service $table_name blog
done

for table_name in post location label image post_user_info author
do
  gen_model_by_table_and_service $table_name post
done

for table_name in page
do
  gen_model_by_table_and_service $table_name page
done

for table_name in comment
do
  gen_model_by_table_and_service $table_name comment
done