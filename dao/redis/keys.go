package redis

// redis key
const (
	KeyPrefix          = "blue:"
	KeyPostTimeZSet    = "post:time"
	KeyPostScoreZSet   = "post:score"
	KeyPostVotedZSetPF = "post:voted:"
	KeyCommunitySetPF  = "community:"
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
