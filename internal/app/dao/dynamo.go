package dao

import (
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewDynamo() (*dynamodb.DynamoDB, error) {
	disableSSLFlg, err := strconv.ParseBool(os.Getenv("DISABLE_SSL_FLG"))
	if err != nil {
		return nil, err
	}

	region := os.Getenv("REGION")
	if region == "" {
		region = "us-east-1"
	}

	endpoint := os.Getenv("ENDPOINT")
	if endpoint == "" {
		endpoint = "http://127.0.0.1:8000"
	}

	return dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region:     aws.String(region),
		Endpoint:   aws.String(endpoint),
		DisableSSL: aws.Bool(disableSSLFlg),
	}))), nil
}
