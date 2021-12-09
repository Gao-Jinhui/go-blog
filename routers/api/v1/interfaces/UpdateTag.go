package interfaces

type UpdateTagRequest struct {
	ID         int    `json:"id" validate:"required"`
	ModifiedBy string `json:"modified_by" validate:"required,lte=100"`
	Name       string `json:"name" validate:"required,lte=100"`
}

type UpdateTagResponse struct {
	BaseResponse `json:"status"`
	Error        []string `json:"error"`
}
