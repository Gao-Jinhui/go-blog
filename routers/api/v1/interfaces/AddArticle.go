package interfaces

type AddArticleRequest struct {
	ID        int    `json:"tag_id" validate:"gt=0"`
	Title     string `json:"title" validate:"required"`
	Desc      string `json:"desc"`
	Content   string `json:"content" validate:"required"`
	CreatedBy string `json:"created_by" validate:"required"`
	State     int    `json:"state" validate:"oneof=0 1"`
}

type AddArticleResponse struct {
	BaseResponse `json:"status"`
	Error        []string `json:"error"`
}
