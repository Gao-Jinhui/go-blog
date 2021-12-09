package interfaces

import "go-blog/models"

type GetArticleByIDRequest struct {
	ID *int `json:"id" validate:"required,gt=0"`
}

type GetArticleByIDResponse struct {
	BaseResponse `json:"status"`
	Data         models.Article `json:"data"`
}

type GetArticlesByTagRequest struct {
	Tag_ID int `json:"tag_id" validate:"required,gt=0"`
	State  int `json:"state" validate:"required,oneof=-1 1"`
}

type GetArticlesByTagResponse struct {
	BaseResponse `json:"status"`
	Data         []models.Article `json:"data"`
}
