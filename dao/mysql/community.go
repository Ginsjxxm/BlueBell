package mysql

import (
	"BlueBell/models"
	"database/sql"
	"errors"
	"go.uber.org/zap"
)

func GetCommunityList() (data []*models.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	if err = db.Select(&data, sqlStr); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Warn("there is no community")
			err = nil
		}
	}
	return
}

func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select community_id,community_name,introduction,create_time from community where community_id = ?`
	err = db.Get(community, sqlStr, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = ErrorInvalidID
		}
	}
	return community, err
}
