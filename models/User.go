package models

type User struct {
	ID       int `gorm:"primaryKey"`
	Username string
	Email    string
}
