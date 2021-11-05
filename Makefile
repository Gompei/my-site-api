init:
	# https://docs.aws.amazon.com/ja_jp/lambda/latest/dg/go-image.html#go-image-other
	go mod tidy
	mkdir -p ~/.aws-lambda-rie && curl -Lo ~/.aws-lambda-rie/aws-lambda-rie \
    https://github.com/aws/aws-lambda-runtime-interface-emulator/releases/latest/download/aws-lambda-rie \
    && chmod +x ~/.aws-lambda-rie/aws-lambda-rie

local:
	docker-compose build
	docker-compose up -d

build:
	GOOS=linux GOARCH=amd64 go build -o handler
	zip lambda.zip handler
