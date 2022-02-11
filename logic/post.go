package logic

import (
	"go.uber.org/zap"
	"myApp/dao/mysql"
	"myApp/models"
	"myApp/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	p.ID = snowflake.GenID()
	return mysql.CreatePost(p)
}
func GetPostByID(pid int64) (data *models.UserData, err error) {
	post, err := mysql.GetPostPid(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostPid failed", zap.Error(err))
		return
	}

	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysqlGetUserById failed.", zap.Error(err))
		return
	}
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysqlGetUserById failed.", zap.Error(err))
		return
	}
	data = &models.UserData{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}
