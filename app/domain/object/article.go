package object

import "encoding/json"

type Article struct {
	ID    int64  `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	//SubTitle         string    `json:"sub_title,omitempty"`
	//CategoryTag      string    `json:"category_tag,omitempty"`
	//Description      string    `json:"description,omitempty"`
	//Content          string    `json:"content,omitempty"`
	//NumberOfAccesses int64     `json:"number_of_accesses,omitempty"`
	//CreateTimeStamp  TimeStamp `json:"create_time_stamp,omitempty"`
	//UpdateTimeStamp  TimeStamp `json:"update_time_stamp,omitempty"`
}

func ToArticleStruct(str string) (*Article, error) {
	var article Article
	err := json.Unmarshal([]byte(str), &article)
	if err != nil {
		return nil, err
	}
	return &article, nil
}
