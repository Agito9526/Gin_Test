package service

import (
	"GolangAPI/middlewares"
	"GolangAPI/pojo"
	"crypto/md5"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// func FindByUserId(c *gin.Context) {
// 	user := pojo.FindByUserId(c.Param("Id"))
// 	if user.UserId == 0 {
// 		c.JSON(http.StatusNotFound, "Error")
// 		return
// 	}
// 	c.JSON(http.StatusOK, user)
// }

func CreateUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}
	if pojo.CheckFindByUserName(user.UserName).UserName != "" {
		c.JSON(http.StatusNotFound, "帳號重複")
		return
	}
	//新增SALT
	user.Salt = middlewares.RandomAtoZWith0to9(20)

	//MD5+SALT
	data := []byte(user.Password + user.Salt)
	has := md5.Sum(data)
	user.Password = fmt.Sprintf("%x", has)

	//新增帳號
	newUser := pojo.CreateUser(user)

	returnUser := pojo.LoginUser(newUser.UserName, newUser.Password)
	// returnUser.Password = ""
	// returnUser.Salt = ""
	c.JSON(http.StatusOK, returnUser)
}

// Login User
func LoginUser(c *gin.Context) {
	sessionId := middlewares.GetSession(c)
	if sessionId != 0 {
		c.JSON(http.StatusUnauthorized, "已登入")
		return
	}

	name := c.PostForm("Name")
	user := pojo.CheckFindByUserName(name)
	if user.UserName == "" {
		c.JSON(http.StatusNotFound, "帳號錯誤或者不存在")
		return
	}

	data := []byte(c.PostForm("Password") + user.Salt)
	has := md5.Sum(data)
	password := fmt.Sprintf("%x", has)
	if !pojo.CheckFindByUserPassword(password) {
		c.JSON(http.StatusNotFound, "密碼錯誤")
		return
	}

	returnUser := pojo.LoginUser(name, password)
	// returnUser.Password = ""
	// returnUser.Salt = ""
	middlewares.SaveSession(c, user.UserId)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Successfully",
		"User":    returnUser,
		"Session": middlewares.GetSession(c),
	})
}

// Logout User
func LogoutUser(c *gin.Context) {
	middlewares.ClearSession(c)
	c.JSON(http.StatusOK, gin.H{
		"message": "登出成功",
	})
}

// Check User Session
func CheckUserSession(c *gin.Context) {
	sessionId := middlewares.GetSession(c)
	if sessionId == 0 {
		c.JSON(http.StatusUnauthorized, "沒有登入")
		return
	}
	Id := strconv.Itoa(sessionId)
	user := pojo.FindByUserId(Id)
	// user.Password = ""
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	sessionId := middlewares.GetSession(c)
	middlewares.ClearSession(c)
	Id := strconv.Itoa(sessionId)
	user := pojo.DeleteUser(Id)
	if !user {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "刪除成功",
	})
}

func UpdateUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
		return
	}

	sessionId := middlewares.GetSession(c)
	Id := strconv.Itoa(sessionId)
	selectUser := pojo.FindByUserId(Id)

	data := []byte(user.Password + selectUser.Salt)
	has := md5.Sum(data)
	password := fmt.Sprintf("%x", has)
	user.Password = password

	if password == selectUser.Password {
		c.JSON(http.StatusUnauthorized, "密碼重複")
		return
	}

	user = pojo.UpdateUser(Id, user)
	// if user.UserId == 0 {
	// 	c.JSON(http.StatusNotFound, "NotFound")
	// 	return
	// }
	c.JSON(http.StatusOK, user)
}
