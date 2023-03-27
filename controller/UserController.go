package controller

import (
	session "GolangAPI/middlewares"
	"GolangAPI/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/user", session.SetSession())

	// user.GET("/:Id", service.FindByUserId)
	user.POST("/create", service.CreateUser)
	//LoginUser
	user.POST("/login", service.LoginUser)
	user.GET("/check", service.CheckUserSession)

	user.Use(session.AuthSession())
	{
		user.GET("/logout", service.LogoutUser)

	}
	user.GET("/deleteuser", service.DeleteUser)
}
