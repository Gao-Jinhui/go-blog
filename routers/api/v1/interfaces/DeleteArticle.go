package interfaces

type DeleteArticleRequest struct {
	ID int `json:"id" validate:"required,gt=0"`
}

type DeleteArticleResponse struct {
	BaseResponse `json:"status"`
}
