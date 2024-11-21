package models

//定义请求的参数结果体

const OrderTime = "time"
const OrderScore = "score"

type ParamSignUp struct {
	Username   string `json:"username" binding:"required" `
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"rePassword" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	PostID    int64  `json:"post_id,string" binding:"required"`         //帖子id
	Direction string `json:"direction" binding:"required,oneof=1 0 -1"` //帖子(赞成1，反对负一
}

type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"`
	Limit       int64  `json:"limit" form:"limit"`
	Offset      int64  `json:"offset" form:"offset"`
	Order       string `json:"order" form:"order"`
}
