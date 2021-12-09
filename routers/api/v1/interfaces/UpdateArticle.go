package interfaces

type UpdateArticleRequest struct {
	ID          int    `json:"id" validate:"required,gt=0"`
	Tag_ID      int    `json:"tag_id" validate:"gt=0"`
	Title       string `json:"title" validate:"lte=100"`
	Desc        string `json:"desc" validate:"lte=255"`
	Content     string `json:"content" validate:"lte=65535"`
	Modified_By string `json:"modified_by" validate:"required,lte=100"`
	State       int    `json:"state" validate:"oneof=-1 1"`
}

type UpdateArticleResponse struct {
	BaseResponse `json:"status"`
}
