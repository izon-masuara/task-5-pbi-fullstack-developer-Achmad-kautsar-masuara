package app

type PostRequest struct {
	Username string `json:"username" binding:"required,min=5,max=100"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=30"`
}

type GetLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
