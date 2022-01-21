package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//定义状态码
type ResponseCode struct {
	Code CodeType    `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(c *gin.Context, code CodeType) {
	rd := &ResponseCode{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}
func ResponseErrorWithMsg(c *gin.Context, code CodeType, msg interface{}) {
	rd := &ResponseCode{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	rd := &ResponseCode{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)
}
