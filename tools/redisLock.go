package tools

import (
	"context"
	"hmdp/global"
	"log"
	"time"
)

func TryLock(ctx context.Context, key string) bool {
	client := global.App.Redis
	isSuccess, err := client.SetNX(ctx, key, "1", LockShopTtl*time.Second).Result()
	if err != nil {
		log.Printf("Failed to get the lock: %v", err)
	}
	return isSuccess
}

func UnLock(key string) {
	client := global.App.Redis
	client.Del(context.Background(), key)
}
