package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"myApp/logic"
	"myApp/models"
	"strconv"
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

func GetPostDetailHandler(c *gin.Context) {
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParm)
		return
	}
	data, err := logic.GetPostByID(pid)
	if err != nil {
		zap.L().Error("logic.GetPostByID failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
func GetPostDetailList(c *gin.Context) {
	offset, limit := getPageInfo(c)
	data, err := logic.GetpostList(offset, limit)
	if err != nil {
		zap.L().Error("logic.GetpostList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)

}
