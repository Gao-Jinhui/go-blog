package interfaces

type DeleteTagRequest struct {
	ID int `json:"id" validator:"required,gt=0"`
}

type DeleteTagResponse struct {
	BaseResponse `json:"status"`
	Error        []string `json:"error"`
}
