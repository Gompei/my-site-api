package repository

import (
	"github.com/Gompei/my-site-api/app/domain/object"
)

type Article interface {
	GetAllArticle() ([]*object.Article, error)
	GetArticle(id int64) (*object.Article, error)
	PutArticle(article *object.Article) error
	DeleteArticle(article *object.Article) error
}
