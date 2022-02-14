package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const ContextUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

func GetCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

//分页
func getPageInfo(c *gin.Context) (int64, int64) {
	offsetStr := c.Query("offset") //页数
	limitStr := c.Query("limit")   //每页显示多少条
	var (
		offset int64
		limit  int64
		err    error
	)
	offset, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 1
	}
	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 10
	}
	return offset, limit
}
