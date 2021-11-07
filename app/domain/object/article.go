package object

type Article struct {
	ID               int64     `json:"id"`
	Title            string    `json:"title"`
	SubTitle         string    `json:"sub_title,omitempty"`
	CategoryTag      string    `json:"category_tag,omitempty"`
	Description      string    `json:"description,omitempty"`
	Content          string    `json:"content,omitempty"`
	NumberOfAccesses int64     `json:"number_of_accesses,omitempty"`
	CreateTimeStamp  TimeStamp `json:"create_time_stamp,omitempty"`
	UpdateTimeStamp  TimeStamp `json:"update_time_stamp,omitempty"`
}
