#!/usr/bin/env bash

DYNAMODB_TABLE_NAME="article_table"

: put test data
for id in $(seq 10)
do
  # Mac
  DATE=$(date -v-"$id"d '+%Y-%m-%d')
  # Linux
  # DATE=$(date --date "$id days ago" '+%Y-%m-%d')
  aws dynamodb put-item  \
    --region us-east-1 \
    --table-name "$DYNAMODB_TABLE_NAME" \
    --item "{ \"id\": { \"N\": \"$id\"},\"create_time_stamp\": {\"S\": \"$DATE\"},\"update_time_stamp\": {\"S\": \"$DATE\"},\"title\": {\"S\": \"example\"},\"sub_title\": { \"S\": \"example\"},\"image_url\": {\"S\": \"example\"},\"category_tag\": {\"SS\": [\"example1\", \"example2\",\"example3\"]},\"description\": {\"S\": \"example\"},\"content\": {\"S\": \"example\"}}"
done
