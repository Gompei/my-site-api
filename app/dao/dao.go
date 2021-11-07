package dao

import (
	"github.com/Gompei/my-site-api/app/domain/repository"
	"github.com/guregu/dynamo"
)

type (
	Dao interface {
		Article() repository.Article
	}

	dao struct {
		client *dynamo.DB
	}
)

func New() (Dao, error) {
	client, err := connectionDataBase()
	if err != nil {
		return nil, err
	}
	return &dao{
		client: client,
	}, nil
}

func (d *dao) Article() repository.Article {
	return NewArticle(d.client)
}
