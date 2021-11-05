build:
	GOOS=linux GOARCH=amd64 go build -o handler
	zip lambda.zip handler
