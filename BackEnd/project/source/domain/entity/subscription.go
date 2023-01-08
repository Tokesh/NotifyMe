package entity

type Subscription struct {
	Id        int    `json:"subs_id"`
	Name      string `json:"subs_name"`
	Code      string `json:"subs_code"`
	Category1 string `json:"category1"`
	Category2 string `json:"category2"`
}
