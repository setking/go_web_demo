package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"myApp/models"
)

const secret = "123456"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
)

//数据库操作封装成函数
//待logic层根据业务需求调用
//检查用户名是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username=?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

//数据库插入新的用户记录
func InsertUser(user *models.User) (err error) {
	user.Password = MD5Password(user.Password)
	sqlStr := `insert into user(user_id, username, password) values(?, ?, ?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

//MD5加密
func MD5Password(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

//检查登录用户信息
func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id, username, password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	fmt.Printf("userData: %v\n", user)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	password := MD5Password(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}
