package interfaces

import "go-blog/models"

type GetTagsRequest struct {
	Name  string `json:"name"`
	State int    `json:"state"`
}

type GetTagsResponse struct {
	BaseResponse `json:"status"`
	Lists        []models.Tag `json:"lists"`
	Total        int          `json:"total"`
}
