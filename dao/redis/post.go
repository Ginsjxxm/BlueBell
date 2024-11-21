package redis

import (
	"BlueBell/models"
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

func GetPostIDByInOrder(p *models.ParamPostList) ([]string, error) {
	//从redis获取id
	ctx := context.Background()
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostTimeZSet)
	}
	limit := p.Limit
	offset := p.Limit + p.Offset
	return client.ZRevRange(ctx, key, limit, offset).Result()
}

// GetPostVoteData 根据ids查询每一篇帖子的赞成票的数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	//data = make([]int64, 0, len(ids))
	ctx := context.Background()
	//for _, id := range ids {
	//	key := getRedisKey(id)
	//	v := client.ZCount(ctx, KeyPostVotedZSetPF+key, "1", "1").Val()
	//	data = append(data, v)
	//}
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		pipeline.ZCount(ctx, key, "1", "1")
	}
	cmder, err := pipeline.Exec(ctx)
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(ids))
	for _, cmd := range cmder {
		v := cmd.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return data, nil
}

func GetCommunityPostIDByInOrder(p *models.ParamPostList) ([]string, error) {
	ctx := context.Background()
	orderkey := getRedisKey(KeyPostTimeZSet) // 默认是时间
	if p.Order == models.OrderScore {        // 按照分数请求
		orderkey = KeyPostScoreZSet
	}
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityID)))
	key := KeyPostTimeZSet + ":" + strconv.Itoa(int(p.CommunityID))
	pipeline := client.Pipeline()
	if client.Exists(ctx, key).Val() < 1 {
		pipeline.ZInterStore(context.Background(), key, &redis.ZStore{
			Keys:      []string{cKey, orderkey},
			Aggregate: "MAX",
		})
		pipeline.Expire(ctx, key, time.Hour)
	}
	_, err := pipeline.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return getIDsFormKey(key, p.Limit, p.Offset)
}

func getIDsFormKey(key string, limit, offset int64) ([]string, error) {
	offset = limit + offset
	ctx := context.Background()
	return client.ZRevRange(ctx, key, limit, offset).Result()
}
