-include .env

MAKEFILE_DIR:=$(dir $(abspath $(lastword $(MAKEFILE_LIST))))

init:
	# https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/go-image.html#go-image-other
	go mod tidy
	mkdir -p ~/.aws-lambda-rie && curl -Lo ~/.aws-lambda-rie/aws-lambda-rie \
    https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie \
    && chmod +x ~/.aws-lambda-rie/aws-lambda-rie

local:
	docker-compose build
	docker-compose up -d

lambda-test:
	curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{}'

build:
	GOOS=linux GOARCH=amd64 go build -o handler
	zip lambda.zip handler

deploy: build
	aws lambda update-function-code \
		--function-name my-site-api \
		--region us-east-1 \
		--zip-file fileb://$(MAKEFILE_DIR)/lambda.zip

api-test:
	curl https://${API_DOMAIN}/api --header 'x-api-key: ${API_KEY}'
