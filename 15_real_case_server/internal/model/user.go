package model

type User struct {
	ID    int64  `json:"id" db:"id"`
	Name  string `json:"name" db:"name" binding:"required,min=1,max=100"`
	Email string `json:"email" db:"email" binding:"required,email"`
}
