#!/usr/bin/env bash

: create dynamodb table
aws dynamodb create-table \
  --table-name "$DYNAMODB_TABLE_NAME" \
  --attribute-definitions AttributeName=id,AttributeType=N \
  --key-schema AttributeName=id,KeyType=HASH \
  --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
  --endpoint-url "$DYNAMODB_ENDPOINT" > /dev/null 2>&1

aws dynamodb create-table \
  --table-name "$DYNAMODB_COUNT_TABLE_NAME" \
  --attribute-definitions AttributeName=name,AttributeType=S \
  --key-schema AttributeName=name,KeyType=HASH \
  --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
  --endpoint-url "$DYNAMODB_ENDPOINT" > /dev/null 2>&1

: put test data
for id in $(seq 10)
do
  # Mac
  # DATE=$(date -v-"$id"d '+%Y-%m-%d')
  # Linux
  DATE=$(date --date "$id days ago" '+%Y-%m-%d')
  echo $DATE
  aws dynamodb put-item  \
    --region ap-northeast-1 \
    --table-name "$DYNAMODB_TABLE_NAME" \
    --endpoint-url "$DYNAMODB_ENDPOINT" \
    --item "{ \"id\": { \"N\": \"$id\"},\"create_time_stamp\": {\"S\": \"$DATE\"},\"update_time_stamp\": {\"S\": \"$DATE\"},\"title\": {\"S\": \"example\"},\"sub_title\": { \"S\": \"example\"},\"image_url\": {\"S\": \"example\"},\"category_tag\": {\"S\": \"example\"},\"description\": {\"S\": \"example\"},\"content\": {\"S\": \"example\"}}"
done

aws dynamodb put-item  \
    --region ap-northeast-1 \
    --table-name "$DYNAMODB_COUNT_TABLE_NAME" \
    --endpoint-url "$DYNAMODB_ENDPOINT" \
    --item "{ \"name\": { \"S\": \"article_count_id\"},\"id\": {\"N\": \"10\"}}"
