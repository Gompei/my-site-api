package dao

import (
	"github.com/Gompei/my-site-api/app/domain/object"
	"github.com/Gompei/my-site-api/app/domain/repository"
	"github.com/guregu/dynamo"
)

const (
	tableName = ""
	hashKey   = ""
)

type Article struct {
	client *dynamo.DB
}

func NewArticle(client *dynamo.DB) repository.Article {
	return &Article{client: client}
}

// GetArticle 記事IDを元に、記事データを検索して返却します
func (r *Article) GetArticle(article *object.Article) (*object.Article, error) {
	var result *object.Article
	if err := r.client.Table(tableName).Get(hashKey, article.ID).One(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *Article) PutArticle(article *object.Article) error {
	if err := r.client.Table(tableName).Put(&article).Run(); err != nil {
		return err
	}
	return nil
}

func (r *Article) DeleteArticle(article *object.Article) error {
	if err := r.client.Table(tableName).Delete(hashKey, article.ID).Run(); err != nil {
		return err
	}
	return nil
}
