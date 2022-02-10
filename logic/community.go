package logic

import (
	"myApp/dao/mysql"
	"myApp/models"
)

//社区列表
func GetCommunity() ([]*models.Community, error) {
	return mysql.GetCommunity()
}

//查询社区分类详情
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
