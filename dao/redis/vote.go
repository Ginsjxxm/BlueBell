package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"math"
	"strconv"
	"time"
)

const (
	onWeekInSeconds = 7 * 24 * 60 * 60
	scoreVote       = 432
)

var (
	ErrVoteTimeExpired = errors.New("投票时间已过")
	ErrVoteError       = errors.New("不允许重复投票")
)

func VoteForPost(userID, postID string, value float64) (err error) {
	ctx := context.Background()
	exists, err := client.Exists(ctx, getRedisKey(KeyPostTimeZSet)).Result()
	if err != nil || exists == 0 {
		// 键不存在
		return errors.New("键不存在")
	}

	parseTime := client.ZScore(ctx, getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-parseTime > onWeekInSeconds {
		return ErrVoteTimeExpired
	}
	//更新帖子分数
	//先查之前的投票记录
	ov := client.ZScore(ctx, getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()
	var op float64
	//好好研究
	if ov == value {
		return ErrVoteError
	}
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value)
	//更新帖子分数
	pipeline := client.Pipeline()
	pipeline.ZIncrBy(ctx, getRedisKey(KeyPostScoreZSet), op*diff*scoreVote, postID)

	if value == 0 {
		pipeline.ZRem(ctx, getRedisKey(KeyPostVotedZSetPF+postID), userID)
	} else {
		pipeline.ZAdd(ctx, getRedisKey(KeyPostVotedZSetPF+postID), &redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err = pipeline.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
func CreatePost(postID uint64, communityId int64) (err error) {
	ctx := context.Background()
	pipeline := client.TxPipeline()
	pipeline.ZAdd(ctx, getRedisKey(KeyPostTimeZSet), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	//帖子
	pipeline.ZAdd(ctx, getRedisKey(KeyPostScoreZSet), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityId)))
	pipeline.ZAdd(ctx, cKey, &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	_, err = pipeline.Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
