package pojo

import (
	DB "GolangAPI/database"
)

type User struct {
	UserId   int    `json:"Id"`
	UserName string `json:"Name"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
	Salt     string `json:"Salt"`
}

func FindByUserId(userId string) User {
	var user User
	DB.DBconnect.Raw("select * from USERS where user_id = ?", userId).Scan(&user)
	return user
}
