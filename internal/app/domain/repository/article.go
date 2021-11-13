package repository

import (
	"context"

	"github.com/Gompei/my-site-api/internal/app/domain/object"
)

type Article interface {
	PutArticle(ctx context.Context, article *object.Article) error
	//GetAllArticle() ([]*object.Article, error)
	//GetArticle(id int64) (*object.Article, error)
	//DeleteArticle(article *object.Article) error
}
