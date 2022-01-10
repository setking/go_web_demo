package controller

import (
	"myApp/logic"
	"myApp/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

//处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	//校验参数
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	//if len(p.RePassword) == 0 || len(p.Password) == 0 || len(p.Username) == 0 || len(p.RePassword) != len(p.Password) {
	//	zap.L().Error("SignUp with invalid param")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数错误",
	//	})
	//	return
	//}
	//业务处理
	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "注册失败!",
		})
		return
	}

	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "请求成功！！！",
	})
}

//处理登录请求的函数
func LoginHandler(c *gin.Context) {
	//校验参数
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	//业务处理
	if err := logic.Login(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码错误!",
		})
		return
	}

	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "请求成功！！！",
	})
}
