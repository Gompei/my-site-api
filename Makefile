-include .env

MAKEFILE_DIR:=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))

build:
	GOOS=linux GOARCH=amd64 go build -o handler ./cmd/my-site-api/main.go
	zip lambda.zip handler

.PHONY: deploy
deploy: build
	aws lambda update-function-code \
		--function-name my-site-api \
		--region ap-northeast-1 \
		--zip-file fileb://$(MAKEFILE_DIR)/lambda.zip
