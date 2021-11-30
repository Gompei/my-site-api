package object

import "encoding/json"

type Article struct {
	ID              int64    `dynamodbav:"id" json:"id,omitempty"`
	Title           string   `dynamodbav:"title" json:"title,omitempty"`
	SubTitle        string   `dynamodbav:"sub_title" json:"sub_title,omitempty"`
	ImageURL        string   `dynamodbav:"image_url" json:"image_url,omitempty"`
	CategoryTag     []string `dynamodbav:"category_tag" json:"category_tag,omitempty"`
	Description     string   `dynamodbav:"description" json:"description,omitempty"`
	Content         string   `dynamodbav:"content" json:"content,omitempty"`
	CreateTimeStamp string   `dynamodbav:"create_time_stamp" json:"create_time_stamp,omitempty"`
	UpdateTimeStamp string   `dynamodbav:"update_time_stamp" json:"update_time_stamp,omitempty"`
	PublicFlg       bool     `dynamodbav:"public_flg" json:"public_flg,omitempty"`
	DeleteFlg       bool     `dynamodbav:"delete_flg" json:"delete_flg,omitempty"`
}

func ToArticleStruct(str string) (*Article, error) {
	var article Article
	err := json.Unmarshal([]byte(str), &article)
	if err != nil {
		return nil, err
	}
	return &article, nil
}
