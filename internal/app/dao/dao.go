package dao

import (
	"github.com/Gompei/my-site-api/internal/app/domain/repository"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Dao interface {
	Article() repository.Article
}

type dao struct {
	db *dynamodb.DynamoDB
}

func New() (Dao, error) {
	db, err := NewDynamo()
	if err != nil {
		return nil, err
	}

	return &dao{
		db: db,
	}, nil
}

func (d *dao) Article() repository.Article {
	return NewArticle(d.db)
}
