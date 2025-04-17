package data

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisMessageCache struct {
	rdb         *redis.Client
	expireAfter time.Duration
}

// Store 把离线消息存入缓存
func (d *pushRepo) Store(uid uint, msg string) error {
	ctx := context.Background()
	key := fmt.Sprintf("O%d", uid)

	//fmt.Println("Store() -> UID:", uid)
	//fmt.Println("Store() -> Key:", key)
	//fmt.Println("Store() -> Message:", msg)
	//fmt.Println("Store() -> ExpireAfter:", d.data.RD.expireAfter)

	// 使用 Redis pipeline 保证原子性
	pipe := d.data.RD.rdb.TxPipeline()

	// 添加消息到 List（RPUSH 保持 FIFO）
	pipe.RPush(ctx, key, msg)

	// 设置过期时间（每次都重置 TTL）
	pipe.Expire(ctx, key, d.data.RD.expireAfter)

	_, err := pipe.Exec(ctx)
	return err
}

// PopAll 获取所有离线消息
func (d *pushRepo) PopAll(uid uint) ([]string, error) {
	ctx := context.Background()
	key := fmt.Sprintf("O%d", uid)

	// 使用 pipeline 读 + 删，减少网络开销
	pipe := d.data.RD.rdb.TxPipeline()

	// 获取所有消息（保留顺序）
	getCmd := pipe.LRange(ctx, key, 0, -1)

	// 删除该 key，防止重复投递
	pipe.Del(ctx, key)

	_, err := pipe.Exec(ctx)
	if err != nil {
		return nil, err
	}

	return getCmd.Val(), nil
}
