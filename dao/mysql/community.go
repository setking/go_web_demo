package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"myApp/models"
)

//社区列表
func GetCommunity() (Community []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community "
	if err := db.Select(&Community, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("Community is no community in db")
			err = nil
		}
	}
	return
}
