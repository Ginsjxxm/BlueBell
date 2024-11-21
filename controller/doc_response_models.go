package controller

import (
	"BlueBell/models"
)

type _ResponsePostList struct {
	Code    ResCode                 `json:"code"`
	Message string                  `json:"message"`
	Data    []*models.ApiPostDetail `json:"data"`
}

type _PostFirst struct {
	CommunityID int64  `json:"community_id,string" db:"community_id" binding:"required"`
	Title       string `json:"title" db:"title" binding:"required"`
	Content     string `json:"content" db:"content" binding:"required"`
}
