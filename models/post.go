package models

import "time"

// Post 内存对其
type Post struct {
	ID          uint64    `json:"id" db:"post_id" binding:"required"`
	AuthorID    uint64    `json:"author_id" db:"author_id" binding:"required"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

type ApiPostDetail struct {
	AuthorName       string `json:"author_name"`
	*Post            `json:"post"`
	*CommunityDetail `json:"community"`
}
