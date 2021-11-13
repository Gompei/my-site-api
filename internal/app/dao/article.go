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
	articleHeadlineTableName             = "article_headline_table"
	articleHeadlineTablePartitionKeyName = "service_name"
	articleHeadlinePartitionKeyValue     = "my_site"
)

const (
	articleContentTableName             = "article_content_table"
	articleContentTablePartitionKeyName = "article_id"
)

type Article struct {
	db *dynamodb.DynamoDB
}

func NewArticle(db *dynamodb.DynamoDB) repository.Article {
	return &Article{db: db}
}

//
//func FindArticleHeadlines() {}
//
//func ListArticleHeadlines() {}
//
//func GetArticleHeadlines() {}
//
//// GetArticle 記事IDを元に、記事データを検索して返却します
//func (r *Article) GetArticle(ctx context.Context, id int64) (*object.Article, error) {
//	i := strconv.FormatInt(id, 10)
//
//	// 記事見出し取得
//
//	// 記事内容
//	result, err := r.db.GetItemWithContext(ctx, &dynamodb.GetItemInput{
//		TableName: aws.String(articleContentTableName),
//		Key: map[string]*dynamodb.AttributeValue{
//			articleContentTablePartitionKeyName: {
//				N: aws.String(i),
//			},
//		},
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	articleContent := &object.ArticleContent{}
//	err = dynamodbattribute.UnmarshalMap(result.Item, articleContent)
//
//	return article, err
//}

// PutArticle　記事データを登録します
func (r *Article) PutArticle(ctx context.Context, article *object.Article) error {
	// 記事見出し登録
	av, err := dynamodbattribute.MarshalMap(article.Headline)
	if err != nil {
		return err
	}
	_, err = r.db.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(articleHeadlineTableName),
		Item:      av,

		ReturnConsumedCapacity:      aws.String("NONE"),
		ReturnItemCollectionMetrics: aws.String("NONE"),
		ReturnValues:                aws.String("NONE"),
	})

	// 記事内容登録
	av, err = dynamodbattribute.MarshalMap(article.Content)
	if err != nil {
		return err
	}
	_, err = r.db.PutItemWithContext(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(articleContentTableName),
		Item:      av,

		ReturnConsumedCapacity:      aws.String("NONE"),
		ReturnItemCollectionMetrics: aws.String("NONE"),
		ReturnValues:                aws.String("NONE"),
	})

	return err
}

// DeleteArticle　記事IDを元に、記事データを削除します(物理削除)
//func (r *Article) DeleteArticle(ctx context.Context, id int64) error {
//	i := strconv.FormatInt(id, 10)
//
//	// 記事見出し削除
//
//	// 記事内容削除
//	_, err := r.db.DeleteItemWithContext(ctx, &dynamodb.DeleteItemInput{
//		TableName: aws.String(articleContentTableName),
//		Key: map[string]*dynamodb.AttributeValue{
//			articleContentTablePartitionKeyName: {
//				S: aws.String(i),
//			},
//		},
//
//		ReturnConsumedCapacity:      aws.String("NONE"),
//		ReturnItemCollectionMetrics: aws.String("NONE"),
//		ReturnValues:                aws.String("NONE"),
//	})
//
//	return err
//}
