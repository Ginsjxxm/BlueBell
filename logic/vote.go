package logic

import (
	"BlueBell/dao/redis"
	"BlueBell/models"
	"strconv"
)

func VoteForPost(userID uint64, p *models.ParamVoteData) (err error) {
	ParseEnd, err := strconv.ParseFloat(p.Direction, 64)
	if err != nil {
		return err
	}
	err = redis.VoteForPost(strconv.Itoa(int(userID)), strconv.Itoa(int(p.PostID)), ParseEnd)
	if err != nil {
		return err
	}
	return nil
}
