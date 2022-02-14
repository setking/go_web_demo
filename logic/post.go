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
func GetpostList(offset, limit int64) (data []*models.UserData, err error) {
	posts, err := mysql.GetPostList(offset, limit)
	if err != nil {
		return nil, err
	}
	data = make([]*models.UserData, 0, len(posts))
	for _, post := range posts {
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysqlGetUserById failed.", zap.Error(err))
			continue
		}
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysqlGetUserById failed.", zap.Error(err))
			continue
		}
		postDetail := &models.UserData{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}
