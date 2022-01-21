package logic

import (
	"myApp/dao/mysql"
	"myApp/models"
	"myApp/pkg/jwt"
	"myApp/pkg/snowflake"
)

//存放业务逻辑代码

func SignUp(p *models.ParamSignUp) (err error) {
	//判断用户不存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	//生成UID
	userID := snowflake.GenID()
	Users := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//snowflake.GenID()
	//密码加密
	//保存进数据库
	return mysql.InsertUser(Users)
}

func Login(p *models.ParamLogin) (token string, err error) {
	//判断用户不存在
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return "", err
	}
	return jwt.GenToken(user.UserID, user.Username)
}
