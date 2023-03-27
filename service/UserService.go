package service

import (
	"GolangAPI/pojo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindByUserId(c *gin.Context) {
	user := pojo.FindByUserId(c.Param("Id"))
	if user.UserId == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, user)
}
