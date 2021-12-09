package interfaces

type BaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ValidatorErrorResponse struct {
	BaseResponse
	Error []string `json:"error"`
}
