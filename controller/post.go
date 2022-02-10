package controller

import (
	"github.com/gin-gonic/gin"
	"myApp/logic"
	"myApp/models"
)

func CreatePostHandler(c *gin.Context) {

	//c.ShouldBindJSON(c)
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CodeInvalidParm)
		return
	}
	userID, err := GetCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	if err := logic.CreatePost(p); err != nil {
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, CodeSuccess)
}
