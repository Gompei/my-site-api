package object

import "encoding/json"

type Article struct {
	Headline ArticleHeadline
	Content  ArticleContent
}

type ArticleHeadline struct {
	ID               int64  `dynamodbav:"article_id" json:"id,omitempty"`
	Title            string `dynamodbav:"title" json:"title,omitempty"`
	SubTitle         string `dynamodbav:"sub_title" json:"sub_title,omitempty"`
	ImageURL         string `dynamodbav:"image_url" json:"image_url,omitempty"`
	CategoryTag      string `dynamodbav:"category_tag" json:"category_tag,omitempty"`
	Description      string `dynamodbav:"description" json:"description,omitempty"`
	NumberOfAccesses int64  `dynamodbav:"content" json:"number_of_accesses,omitempty"`
	CreateTimeStamp  string `dynamodbav:"create_time_stamp" json:"create_time_stamp,omitempty"`
	UpdateTimeStamp  string `dynamodbav:"update_time_stamp" json:"update_time_stamp,omitempty"`
	DeleteFlg        bool   `dynamodbav:"delete_flg" json:"delete_flg,omitempty"`
}

type ArticleContent struct {
	ID              int64  `dynamodbav:"article_id" json:"id,omitempty"`
	Content         string `dynamodbav:"content" json:"content,omitempty"`
	CreateTimeStamp string `dynamodbav:"create_time_stamp" json:"create_time_stamp,omitempty"`
	UpdateTimeStamp string `dynamodbav:"update_time_stamp" json:"update_time_stamp,omitempty"`
	DeleteFlg       bool   `dynamodbav:"delete_flg" json:"delete_flg,omitempty"`
}

func ToArticleStruct(str string) (*Article, error) {
	var article Article
	err := json.Unmarshal([]byte(str), &article)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func ToArticleHeadlineStruct(str string) (*ArticleHeadline, error) {
	var article ArticleHeadline
	err := json.Unmarshal([]byte(str), &article)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func ToArticleContentStruct(str string) (*ArticleContent, error) {
	var article ArticleContent
	err := json.Unmarshal([]byte(str), &article)
	if err != nil {
		return nil, err
	}
	return &article, nil
}
