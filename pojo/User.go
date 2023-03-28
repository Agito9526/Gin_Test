package pojo

import (
	DB "GolangAPI/database"
)

type User struct {
	UserId   int    `json:"UserId"`
	UserName string `json:"Name"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
	Salt     string `json:"Salt"`
}

func FindByUserId(userId string) User {
	var user User
	DB.DBconnect.Raw("select * from USERS where USER_ID = ?", userId).Scan(&user)
	return user
}

func CreateUser(user User) User {
	DB.DBconnect.Create(&user)
	return user
}

// LoginPassword
func LoginUser(name string, password string) User {
	user := User{}
	DB.DBconnect.Where("USER_NAME = ? and PASSWORD = ?", name, password).First(&user)
	return user
}

func CheckFindByUserName(userName string) User {
	var user User
	DB.DBconnect.Raw("select * from USERS where USER_NAME = ?", userName).Scan(&user)
	return user
}

func CheckFindByUserPassword(password string) bool {
	var user User
	DB.DBconnect.Raw("select * from USERS where PASSWORD = ?", password).Scan(&user)
	if user.UserName != "" {
		return true
	}
	return false
}

func UpdateUser(UserId string, user User) User {
	DB.DBconnect.Model(&user).Where("USER_ID = ?", UserId).Updates(user)
	return user
}

// DeleteUser
func DeleteUser(userId string) bool {
	user := User{}
	result := DB.DBconnect.Where("USER_ID = ?", userId).Delete(&user)
	return result.RowsAffected > 0
}
