package interfaces

type AddArticleRequest struct {
	Tag_ID     int    `json:"tag_id" validate:"gt=0"`
	Title      string `json:"title" validate:"required"`
	Desc       string `json:"desc"`
	Content    string `json:"content" validate:"required"`
	Created_By string `json:"created_by" validate:"required"`
	State      int    `json:"state" validate:"oneof=-1 1"`
}

type AddArticleResponse struct {
	BaseResponse `json:"status"`
}
