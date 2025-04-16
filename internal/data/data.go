package data

import (
	"github.com/WH-5/push-service/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewPushRepo)

// Data .
type Data struct {
	WSD *WSData
	RD  *RedisMessageCache
}

//type Other struct {
//	MessageExpiredTimeHour int32
//}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {

	//初始化ws
	ws := NewWSData()

	//初始化redis
	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Database),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})
	cleanup := func() {
		logHelper := log.NewHelper(logger)
		log.NewHelper(logger).Info("closing the data resources")
		if err := rdb.Close(); err != nil {
			logHelper.Errorf("failed to close Redis DB: %v", err)
		}
	}
	exa := c.Redis.MessageExpiredTimeHour.AsDuration()
	//exa := 1 * time.Hour
	return &Data{WSD: ws, RD: &RedisMessageCache{rdb: rdb, expireAfter: exa}}, cleanup, nil
}
