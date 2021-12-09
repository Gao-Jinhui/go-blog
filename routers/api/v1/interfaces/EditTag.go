package interfaces

type EditTagRequest struct {
	ID         int    `json:"id" validate:"required"`
	ModifiedBy string `json:"modified_by" validate:"required,lte=100"`
	Name       string `json:"name" validate:"required,lte=100"`
}

type EditTagResponse struct {
	BaseResponse `json:"status"`
	Error        []string `json:"error"`
}
