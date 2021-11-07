package dao

import (
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

func connectionDataBase() (*dynamo.DB, error) {
	disableSSLFlg, err := strconv.ParseBool(os.Getenv("DISABLE_SSL_FLG"))
	if err != nil {
		return nil, err
	}

	return dynamo.New(session.Must(session.NewSession(&aws.Config{
		Region:     aws.String(os.Getenv("REGION")),
		Endpoint:   aws.String(os.Getenv("ENDPOINT")),
		DisableSSL: aws.Bool(disableSSLFlg),
	}))), nil
}
