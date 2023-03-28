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

	user.GET("/list/all", service.FindAllListByUserId)
	user.GET("/list", service.FindListByListId)

	user.POST("/list/up", service.UpdateList)

	user.Use(session.AuthSession())
	{
		user.GET("/logout", service.LogoutUser)
		user.GET("/deleteuser", service.DeleteUser)
		user.PUT("/updateuser", service.UpdateUser)
		user.POST("/list/create", service.CreateList)
		user.POST("/list/upstatus", service.UpdateListStatus)
		user.POST("/list/deletelist", service.DeleteList)
	}
}
