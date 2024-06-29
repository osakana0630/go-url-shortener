package store

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// Top level declarations for the storeService and Redis context
var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

// CacheDuration
// 現実の使用においては、キャッシュの有効期間は設定せず、キャッシュがいっぱいになった場合には、
// 取得頻度が低い値が自動的にキャッシュからパージされ、RDBMSに戻されるように、
// LRU(Least Recently Used)ポリシー構成を設定する必要があります。
const CacheDuration = 6 * time.Hour

// StorageService
// Define the struct wrapper around raw Redis client
type StorageService struct {
	redisClient *redis.Client
}

// InitializeStore
// Initializing the store service and return a store pointer
func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

// SaveUrlMapping
// 元URLと生成された短いURLをマッピングを保存できるようにしたい。
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
	fmt.Printf("Saved shortUrl: %s - originalUrl: %s\n", shortUrl, originalUrl)
}

/*
RetrieveInitialUrl
短いURLをもとに元の長いURLを取得する
*/
func RetrieveInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}
