package repository

import (
	"context"

	"github.com/Gompei/my-site-api/internal/app/domain/object"
)

type Article interface {
	ListArticles(ctx context.Context) ([]*object.Article, error)
	GetArticle(ctx context.Context, id string) (*object.Article, error)
}
