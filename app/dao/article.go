package dao

import (
	"github.com/Gompei/my-site-api/app/domain/object"
	"github.com/Gompei/my-site-api/app/domain/repository"
	"github.com/guregu/dynamo"
)

const (
	tableName    = "article_table"
	partitionKey = "article_id"
)

type Article struct {
	client *dynamo.DB
}

func NewArticle(client *dynamo.DB) repository.Article {
	return &Article{client: client}
}

// GetAllArticle 全件の記事データを返却します
func (r *Article) GetAllArticle() ([]*object.Article, error) {
	var articles []*object.Article
	if err := r.client.Table(tableName).Scan().All(&articles); err != nil {
		return nil, err
	}
	return articles, nil
}

// GetArticle 記事IDを元に、記事データを検索して返却します
func (r *Article) GetArticle(id int64) (*object.Article, error) {
	var result *object.Article
	if err := r.client.Table(tableName).Get(partitionKey, id).One(&result); err != nil {
		return nil, err
	}
	return result, nil
}

// PutArticle　記事データを登録します
func (r *Article) PutArticle(article *object.Article) error {
	if err := r.client.Table(tableName).Put(&article).Run(); err != nil {
		return err
	}
	return nil
}

// DeleteArticle　記事IDを元に、記事データを削除します
func (r *Article) DeleteArticle(article *object.Article) error {
	if err := r.client.Table(tableName).Delete(partitionKey, article.ID).Run(); err != nil {
		return err
	}
	return nil
}
