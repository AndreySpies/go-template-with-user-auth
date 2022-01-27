package entity

import "time"

type User struct {
	Id           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Password     []byte    `json:"password"`
	Birthday     time.Time `json:"birthday"`
	Location     string    `json:"location"`
	TotalBalance float64   `json:"total_balance"`
	CreatedAt    time.Time `json:"created_at"`
}
