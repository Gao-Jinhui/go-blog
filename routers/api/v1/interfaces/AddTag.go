package interfaces

type AddTagRequest struct {
	Name      string `json:"name" validate:"required,max=20"`
	State     int    `json:"state" validate:"oneof=0 1"`
	CreatedBy string `json:"created_by" validate:"required,lte=100"`
}

type AddTagResponse struct {
	BaseResponse `json:"status"`
	Error        []string `json:"error"`
}
