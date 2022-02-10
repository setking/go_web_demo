package logic

import (
	"myApp/dao/mysql"
	"myApp/models"
	"myApp/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	p.ID = snowflake.GenID()
	return mysql.CreatePost(p)
}
