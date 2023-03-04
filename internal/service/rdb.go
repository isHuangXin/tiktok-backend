package service

import (
	"github.com/go-redis/redis"
	initialization "github.com/isHuangxin/tiktok-backend/init"
	"math/rand"
	"sync"
	"time"
)

var (
	redisClient *redis.Client
	redisOnce   sync.Once
)

func initRedis() {
	redisOnce.Do(func() {
		redisClient = initialization.GetRDB()
	})
}

func getFavoriteRandomTime() time.Duration {
	return time.Duration(int64(videoFavoriteExpireTime) + rand.Int63n(int64(12*time.Hour)))
}
