package message

type Paging struct {
	NumPerPage int    `json:"numPerPage" bson:"numPerPage"`
	Sort       string `json:"sort" bson:"sort"`
	Page       int    `json:"page" bson:"page"`
}
