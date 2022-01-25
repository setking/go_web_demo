package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"myApp/logic"
)

//----社区列表----

func Community(c *gin.Context) {
	data, err := logic.GetCommunity()
	if err != nil {
		zap.L().Error("logic.GetCommunity() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
