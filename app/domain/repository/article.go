package repository

import (
	"github.com/Gompei/my-site-api/app/domain/object"
)

type Article interface {
	GetArticle(article *object.Article) (*object.Article, error)
	PutArticle(article *object.Article) error
	DeleteArticle(article *object.Article) error
}
