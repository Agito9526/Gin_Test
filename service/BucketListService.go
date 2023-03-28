package service

import (
	"GolangAPI/middlewares"
	"GolangAPI/pojo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateList(c *gin.Context) {
	list := pojo.Bucketlist{}
	err := c.BindJSON(&list)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}
	sessionID := strconv.Itoa(middlewares.GetSession(c))
	createStatus := pojo.GreateList(sessionID, list.ListTitle, list.Content)
	if !createStatus {
		c.JSON(http.StatusBadRequest, "Error")
		return
	}
	returnList := pojo.FindLastListByUserId(sessionID)
	c.JSON(http.StatusOK, returnList)
}

func FindAllListByUserId(c *gin.Context) {
	sessionID := strconv.Itoa(middlewares.GetSession(c))
	list := pojo.FindAllListByUserId(sessionID)
	user := pojo.FindByUserId(sessionID)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
		"list": list,
	})
}

func FindListByListId(c *gin.Context) {
	list := pojo.Bucketlist{}
	err := c.BindJSON(&list)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}
	returnList := pojo.FindListByListId(strconv.Itoa(list.ListId))
	if returnList.ListId == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, returnList)
}

func UpdateList(c *gin.Context) {
	list := pojo.Bucketlist{}
	err := c.BindJSON(&list)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}
	list = pojo.UpdateList(strconv.Itoa(list.ListId), list)
	returnList := pojo.FindListByListId(strconv.Itoa(list.ListId))
	c.JSON(http.StatusOK, returnList)
}

func UpdateListStatus(c *gin.Context) {
	list := pojo.Bucketlist{}
	err := c.BindJSON(&list)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}
	if !pojo.UpdateListStatus(strconv.Itoa(list.ListId)) {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	list = pojo.FindListByListId(strconv.Itoa(list.ListId))
	c.JSON(http.StatusOK, list)
}

func DeleteList(c *gin.Context) {
	list := pojo.Bucketlist{}
	err := c.BindJSON(&list)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}
	listStatus := pojo.DeleteList(strconv.Itoa(list.ListId))
	if !listStatus {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "刪除成功",
	})
}
