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

func CreateUser(user User) User {
	DB.DBconnect.Create(&user)
	return user
}

// LoginPassword
func LoginUser(name string, password string) User {
	user := User{}
	DB.DBconnect.Where("User_Name = ? and Password = ?", name, password).First(&user)
	return user
}

func CheckFindByUserName(userName string) User {
	var user User
	// fmt.Println("帳號確認:" + userName)
	DB.DBconnect.Raw("select * from USERS where user_name = ?", userName).Scan(&user)
	// if user.UserName != "" {
	// 	return true
	// }
	return user
}

func CheckFindByUserPassword(password string) bool {
	var user User
	// fmt.Println("密碼確認:" + password)
	DB.DBconnect.Raw("select * from USERS where password = ?", password).Scan(&user)
	if user.UserName != "" {
		return true
	}
	return false
}

// DeleteUser
func DeleteUser(userId string) bool {
	user := User{}
	result := DB.DBconnect.Where("user_id = ?", userId).Delete(&user)
	// log.Println(result)
	// if result.RowsAffected == 0 {
	// 	return false
	// }
	return result.RowsAffected > 0
}
