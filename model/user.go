package model

type User struct {
	ID int `json:"id"`
	Name string `json:"name" binding:"required,min=2,max=30,alpha"`
	Email string `json:"email" binding:"required,email"`
}