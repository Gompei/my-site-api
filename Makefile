-include .env

MAKEFILE_DIR:=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))

lambda-build:
	GOOS=linux GOARCH=amd64 go build -o handler ./cmd/my-site-api/main.go
	zip lambda.zip handler

.PHONY: deploy
deploy: lambda-build
	aws lambda update-function-code \
		--function-name my-site-api \
		--region us-east-1 \
		--zip-file fileb://$(MAKEFILE_DIR)/lambda.zip
