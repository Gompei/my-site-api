package dao

import (
	"context"

	"github.com/Gompei/my-site-api/internal/app/domain/object"
	"github.com/Gompei/my-site-api/internal/app/domain/repository"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	tableName                = "article_table"
	partitionKey             = "id"
	listProjectionExpression = "id,title,sub_title,image_url,category_tag,description,create_time_stamp,update_time_stamp,public_flg,delete_flg"
)

type Article struct {
	db *dynamodb.DynamoDB
}

func NewArticle(db *dynamodb.DynamoDB) repository.Article {
	return &Article{db: db}
}

// ListArticles　記事データを全件取得します
// 本来であればパーティションキー・ソートキーを組み合わせた検索・取得制限を実施するべきだが、
// パーティションキーを記事ID(一意)にしているため、上記の実装に修正しても内部的な処理(計算量)は同じになってしまうと考えている。(今後調査)
// よって、データ量も多くないため全件返すように実装している。(ハンドラ関数内でデータを加工)
func (r *Article) ListArticles(ctx context.Context) ([]*object.Article, error) {
	scanOut, err := r.db.ScanWithContext(ctx, &dynamodb.ScanInput{
		TableName:            aws.String(tableName),
		ProjectionExpression: aws.String(listProjectionExpression),
	})
	if err != nil {
		return nil, err
	}

	var articles []*object.Article
	for _, s := range scanOut.Items {
		var article *object.Article
		err = dynamodbattribute.UnmarshalMap(s, &article)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

// GetArticle 記事IDを元に、記事データを検索して返却します
func (r *Article) GetArticle(ctx context.Context, id string) (*object.Article, error) {
	result, err := r.db.GetItemWithContext(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			partitionKey: {
				N: aws.String(id),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	var article *object.Article
	err = dynamodbattribute.UnmarshalMap(result.Item, &article)

	return article, err
}

// PutArticle　記事データを登録・更新します
func (r *Article) PutArticle(ctx context.Context, article *object.Article) error {
	av, err := dynamodbattribute.MarshalMap(article)
	if err != nil {
		return err
	}

	_, err = r.db.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,

		ReturnConsumedCapacity:      aws.String("NONE"),
		ReturnItemCollectionMetrics: aws.String("NONE"),
		ReturnValues:                aws.String("NONE"),
	})

	return err
}

// DeleteArticle　記事IDを元に、記事データを削除します(物理削除)
func (r *Article) DeleteArticle(ctx context.Context, id string) error {
	_, err := r.db.DeleteItemWithContext(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			partitionKey: {
				N: aws.String(id),
			},
		},

		ReturnConsumedCapacity:      aws.String("NONE"),
		ReturnItemCollectionMetrics: aws.String("NONE"),
		ReturnValues:                aws.String("NONE"),
	})

	return err
}
