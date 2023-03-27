package controller

import (
	"GolangAPI/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/user")

	user.GET("/:Id", service.FindByUserId)
}
