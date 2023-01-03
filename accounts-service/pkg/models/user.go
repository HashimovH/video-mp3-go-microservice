package models

type User struct{
	Id int `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email string `json:"email"`
	Password string `json:"password"`
}