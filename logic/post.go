package logic

import (
	"BlueBell/dao/mysql"
	"BlueBell/dao/redis"
	"BlueBell/models"
	"BlueBell/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	//生成post id
	p.ID, _ = snowflake.GenID()
	//保存入库
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.ID, p.CommunityID)
	if err != nil {
		return err
	}
	return nil
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

func GetPostList2(p *models.ParamPostList) ([]*models.ApiPostDetail, error) {
	//redis查询id列表
	ids, err := redis.GetPostIDByInOrder(p)
	if err != nil {
		zap.L().Error("redis.GetPostIDByInOrder(p) failed", zap.Error(err))
		return nil, err
	}
	if len(ids) == 0 {
		return []*models.ApiPostDetail{}, nil
	}
	data := make([]*models.ApiPostDetail, 0)
	//根据列表查询数据
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		zap.L().Error("mysql.GetPostListByIDs(ids) failed", zap.Error(err))
		return nil, err
	}
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		zap.L().Error("redis.GetPostVoteData(ids) failed", zap.Error(err))
		return nil, err
	}
	for idx, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(post.AuthorID) failed", zap.Error(err))
			return nil, err
		}
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed", zap.Error(err))
			return nil, err
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return data, nil
}

func GetCommunityList2(p *models.ParamPostList) ([]*models.ApiPostDetail, error) {
	ids, err := redis.GetCommunityPostIDByInOrder(p)
	if err != nil {
		zap.L().Error("redis.GetPostIDByInOrder(p) failed", zap.Error(err))
		return nil, err
	}
	if len(ids) == 0 {
		return []*models.ApiPostDetail{}, nil
	}
	data := make([]*models.ApiPostDetail, 0)
	//根据列表查询数据
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		zap.L().Error("mysql.GetPostListByIDs(ids) failed", zap.Error(err))
		return nil, err
	}
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		zap.L().Error("redis.GetPostVoteData(ids) failed", zap.Error(err))
		return nil, err
	}
	for idx, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(post.AuthorID) failed", zap.Error(err))
			return nil, err
		}
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed", zap.Error(err))
			return nil, err
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return data, nil
}

func GetPostListNew(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	if p.CommunityID == 0 {
		data, err = GetPostList2(p)
	} else {
		data, err = GetCommunityList2(p)
	}
	if err != nil {
		return nil, err
	}
	return data, nil
}
