package interfaces

type AddArticleRequest struct {
	ID        int    `json:"tag_id" validator:"gt=0"`
	Title     string `json:"title" validator:"required"`
	Desc      string `json:"desc"`
	Content   string `json:"content" validator:"required"`
	CreatedBy string `json:"created_by" validator:"required"`
	State     int    `json:"state" validator:"oneof=0 1"`
}

type AddArticleResponse struct {
	BaseResponse `json:"status"`
	Error        []string `json:"error"`
}
