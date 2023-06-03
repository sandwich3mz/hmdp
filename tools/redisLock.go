package tools

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"log"
	"strconv"
	"time"
)

const (
	// 加锁脚本
	/*lockScript = "if  redis.call('GET', KEYS[1]) == ARGV[1] " +
	"then redis.call('EXPIRE', KEYS[1], ARGV[2]) " +
	"return 1 " +
	"else " +
	"return redis.call('SET', KEYS[1], ARGV[1], 'NX', 'PX', ARGV[2]) " +
	"end"*/

	// 解锁的lua脚本
	unlockScript = "if  redis.call('GET', KEYS[1]) == ARGV[1] " +
		"then return redis.call('DEL', KEYS[1]) " +
		"else " +
		"return 0 " +
		"end"

	// 看门狗脚本
	watchDogScript = "if  redis.call('GET', KEYS[1]) == ARGV[1] " +
		"then return redis.call('EXPIRE', KEYS[1], ARGV[2]) " +
		"else " +
		"return 0 " +
		"end"
)

type DispersedLock struct {
	// redis客户端
	client *redis.Client
	// 过期时间
	expire uint64
	// 锁的key
	key string
	// 锁的value 防止被别人获取
	value string
	// 解锁通知通道
	unlockCh chan struct{}
}

// NewDispersedLock 返回分布式锁
func NewDispersedLock(client *redis.Client, key string, expire uint64) *DispersedLock {
	l := &DispersedLock{
		client: client,
		expire: expire,
		key:    key,
		value:  uuid.NewString(),
	}
	l.unlockCh = make(chan struct{}, 0)
	return l
}

// TryLock 加锁
func (l *DispersedLock) TryLock(ctx context.Context) (bool, error) {

	res, err := l.client.SetNX(ctx, l.key, l.value, time.Duration(l.expire)*time.Second).Result()
	if err != nil {
		return false, err
	}
	if !res {
		err := errors.New("fail to lock")
		return false, err
	} else {
		go l.watchDog(ctx)
	}
	return true, nil
}

// UnLock 解锁
func (l *DispersedLock) UnLock(ctx context.Context) (bool, error) {
	res, err := l.client.Eval(ctx, unlockScript, []string{l.key}, []string{l.value}).Result()
	if err != nil {
		return false, err
	}
	if res == 0 {
		err := errors.New("fail to del the lock")
		return false, err
	}
	close(l.unlockCh)
	return true, nil
}

// 看门狗
func (l *DispersedLock) watchDog(ctx context.Context) {
	loopTime := time.Duration(l.expire*2/3) * time.Second
	expTicker := time.NewTicker(loopTime)

	for {
		select {
		// 定时任务激活看门狗
		case <-expTicker.C:
			res, err := l.client.Eval(ctx, watchDogScript, []string{l.key}, []string{l.value, strconv.Itoa(int(l.expire))}).Result()
			if err != nil {
				log.Printf("watchDog Error : %v", err)
				return
			}
			if res == 0 {
				return
			}
			log.Printf("锁 %v 续费成功", l.key)
			// 关闭看门狗
		case <-l.unlockCh:
			return
		}
	}
}

//func TryLock(ctx context.Context, key string) bool {
//	client := global.App.Redis
//	isSuccess, err := client.SetNX(ctx, key, "1", LockShopTtl*time.Second).Result()
//	if err != nil {
//		log.Printf("Failed to get the lock: %v", err)
//	}
//	return isSuccess
//}
//
//func UnLock(key string) {
//	client := global.App.Redis
//	client.Del(context.Background(), key)
//}
