package interfaces

type GetAuthRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type GetAuthResponse struct {
	BaseResponse
}
