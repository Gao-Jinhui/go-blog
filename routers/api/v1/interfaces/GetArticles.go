package interfaces

import "go-blog/models"

type GetArticleByIDRequest struct {
	ID int `json:"id" validate:"required,gt=0"`
}

type GetArticleByIDResponse struct {
	BaseResponse `json:"status"`
	Data         models.Article `json:"data"`
}
