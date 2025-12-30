package zset

import (
	"context"
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type Zset struct {
	Key    string
	Score  int64
	Member string
	redis  *redis.Client
}

func NewZset(ctx context.Context, key string) *Zset {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	// test connection
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	return &Zset{Key: key, redis: client}
}

func (z *Zset) Add(ctx context.Context, score int64, member string) error {
	return z.redis.ZAdd(ctx, z.Key, &redis.Z{
		Score:  float64(score),
		Member: member,
	}).Err()
}

func (z *Zset) Get(ctx context.Context, member string) (float64, error) {
	return z.redis.ZScore(ctx, z.Key, member).Result()
}

func (z *Zset) Remove(ctx context.Context, member string) error {
	return z.redis.ZRem(ctx, z.Key, member).Err()
}

func (z *Zset) GetAll(ctx context.Context) ([]string, error) {
	return z.redis.ZRange(ctx, z.Key, 0, -1).Result()
}

func (z *Zset) GetGreaterThan(ctx context.Context, score float64) ([]string, error) {
	return z.redis.ZRangeByScore(ctx, z.Key, &redis.ZRangeBy{
		Min:    strconv.FormatFloat(score, 'f', -1, 64),
		Max:    "+inf",
		Offset: 0,
		Count:  100,
	}).Result()
}

func (z *Zset) GetGreaterThanOrEqual(ctx context.Context, score float64) ([]string, error) {
	return z.redis.ZRangeByScore(ctx, z.Key, &redis.ZRangeBy{
		Min:    strconv.FormatFloat(score, 'f', -1, 64),
		Max:    "+inf",
		Offset: 0,
		Count:  100,
	}).Result()
}

func (z *Zset) GetLessThan(ctx context.Context, score float64) ([]string, error) {
	return z.redis.ZRangeByScore(ctx, z.Key, &redis.ZRangeBy{
		Min:    "-inf",
		Max:    strconv.FormatFloat(score, 'f', -1, 64),
		Offset: 0,
		Count:  100,
	}).Result()
}

func (z *Zset) GetBetween(ctx context.Context, min float64, max float64) ([]string, error) {
	return z.redis.ZRangeByScore(ctx, z.Key, &redis.ZRangeBy{
		Min:    strconv.FormatFloat(min, 'f', -1, 64),
		Max:    strconv.FormatFloat(max, 'f', -1, 64),
		Offset: 0,
		Count:  100,
	}).Result()
}

func (z *Zset) RemoveAll(ctx context.Context) error {
	return z.redis.Del(ctx, z.Key).Err()
}
