#!/usr/bin/env bash

DYNAMODB_ENDPOINT="http://localhost:8000"
DYNAMODB_TABLE_NAME="article_table"

: create dynamodb table
aws dynamodb create-table \
  --table-name "$DYNAMODB_TABLE_NAME" \
  --attribute-definitions AttributeName=id,AttributeType=N \
  --key-schema AttributeName=id,KeyType=HASH \
  --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
  --endpoint-url "$DYNAMODB_ENDPOINT"

: put test data
for i in $(seq 100)
do
  DATE=$(date -v-"$i"d '+%Y-%m-%dT%H:%M:%S%z')
  aws dynamodb put-item  \
    --table-name "$DYNAMODB_TABLE_NAME" \
    --item "{\"id\": $(i),\"title\": \"example\",\"sub_title\": \"example\",\"image_url\": \"example\",\"category_tag\": \"example\",\"description\": \"example\",\"content\": \"example\",\"create_time_stamp\": $(DATE),\"update_time_stamp\": $(DATE)}" \
    --endpoint-url "$DYNAMODB_ENDPOINT"
done
