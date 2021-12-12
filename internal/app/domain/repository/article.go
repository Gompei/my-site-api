package repository

import (
	"context"

	"github.com/Gompei/my-site-api/internal/app/domain/object"
)

type Article interface {
	ListArticles(ctx context.Context) ([]*object.Article, error)
	PutArticle(ctx context.Context, article *object.Article) error
	GetArticle(ctx context.Context, id string) (*object.Article, error)
	DeleteArticle(ctx context.Context, id string) error
	GetCountID(ctx context.Context) (int64, error)
	PutCountID(ctx context.Context, id int64) error
}
