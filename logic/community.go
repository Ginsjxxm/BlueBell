package logic

import (
	"BlueBell/dao/mysql"
	"BlueBell/models"
)

func GetCommunityList() (data []*models.Community, err error) {
	//查找所有的community
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (data *models.CommunityDetail, err error) {
	return mysql.GetCommunityDetailByID(id)
}
