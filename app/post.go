package app

type PostRequest struct {
	Username string `json:"username" binding:"required,min=5,max=50"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=8,max=30"`
}

type GetLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PostPhoto struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption" binding:"required,max=255"`
	PhotoUrl string `json:"photo_url" binding:"required,max=255"`
}
