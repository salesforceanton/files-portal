package files_portal

import "mime/multipart"

type User struct {
	Id       int    `json:"-"        db:"id"`
	Email    string `json:"email"    db:"email"         binding:"required"`
	Username string `json:"username" db:"username"      binding:"required"`
	Password string `json:"password" db:"password_hash" binding:"required"`
}

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type File struct {
	Id          int    `json:"id"           db:"id"`
	Size        int    `json:"size"         db:"size"`
	Name        string `json:"name"         db:"name"`
	Url         string `json:"url"          db:"url"`
	CreatedDate string `json:"created_date" db:"created_date"`
	OwnerId     int    `json:"owner_id"     db:"owner_id"`
}

type FileItem struct {
	Filename string
	Source   multipart.File
	Size     int64
}

type FilesResponse struct {
	Data []File
}
