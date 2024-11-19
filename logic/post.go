package logic

import (
	"BlueBell/dao/mysql"
	"BlueBell/models"
	"BlueBell/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	//生成post id
	p.ID, _ = snowflake.GenID()
	//保存入库
	return mysql.CreatePost(p)
}

func GetPostByID(pid int64) (data *models.ApiPostDetail, err error) {
	//查询拼接接口下的数据
	data = new(models.ApiPostDetail)
	post, err := mysql.GetPostByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(pid) failed", zap.Error(err), zap.Int64("pid", pid))
		return
	}
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID(post.AuthorID)failed", zap.Error(err), zap.Int64("pid", pid))
		return
	}
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID)failed", zap.Error(err), zap.Int64("pid", pid))
		return
	}
	data.AuthorName = user.Username
	data.Post = post
	data.CommunityDetail = community
	return
}

func GetPostList(offset int64, limit int64) (data []*models.ApiPostDetail, err error) {
	data = make([]*models.ApiPostDetail, 0)
	posts, err := mysql.GetPostList(offset, limit)
	if err != nil {
		return nil, err
	}
	for _, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			return nil, err
		}
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			return nil, err
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}

	return
}
