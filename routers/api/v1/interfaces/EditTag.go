package interfaces

type EditTagRequest struct {
	ID         int    `json:"id" validator:"required"`
	ModifiedBy string `json:"modified_by" validator:"required,lte=100"`
	Name       string `json:"name" validator:"required,lte=100"`
}

type EditTagResponse struct {
	BaseResponse `json:"status"`
	Error        []string `json:"error"`
}
