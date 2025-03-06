package models

import "time"

type User struct {
	ID 			int64  		`json:"userId"`
	UserName 	string 		`json:"username"`
	Password 	string 		`json:"password"`
	CreatedAt 	time.Time 	`json:"createdAt"`
}