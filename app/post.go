package app

import "mime/multipart"

type PostRequest struct {
	Username string               `form:"username" binding:"required,min=5,max=100"`
	Email    string               `form:"email" binding:"required"`
	Password string               `form:"password" binding:"required,min=8,max=30"`
	Photo    multipart.FileHeader `form:"photo" binding:"required"`
}

type PhotoRequest struct {
	Filename string
	Size     int
	Mimetype string
}
