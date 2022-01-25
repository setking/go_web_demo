package logic

import (
	"myApp/dao/mysql"
	"myApp/models"
)

//社区列表
func GetCommunity() ([]*models.Community, error) {
	return mysql.GetCommunity()
}
