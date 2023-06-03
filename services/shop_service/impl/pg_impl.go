package impl

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"hmdp/ent"
	"hmdp/ent/shop"
	"hmdp/global"
	"hmdp/model"
	"hmdp/services/shop_service"
	"hmdp/tools"
	"log"
	"strconv"
	"time"
)

type pgImpl struct {
	dbClient *ent.Client
}

func NewPgImpl(dbClient *ent.Client) shop_service.Repository {
	return &pgImpl{
		dbClient: dbClient,
	}
}

func (p *pgImpl) QueryShopById(ctx context.Context, id string) (*ent.Shop, error) {
	res, err := p.queryWithPassThrough(ctx, id)
	return res, err
}

// 缓存空值解决缓存穿透
func (p *pgImpl) queryWithPassThrough(ctx context.Context, id string) (*ent.Shop, error) {
	client := global.App.Redis
	key := tools.CacheShopKey + id

	// 查询缓存
	shopJson, err := client.Get(context.Background(), key).Result()
	// 抛出除了查不到键的错误
	if err != nil && err != redis.Nil {
		log.Printf("Failed to get qShop from Redis: %v", err)
		return nil, err
	}
	var qShop *ent.Shop
	if shopJson != "" {
		if err := json.Unmarshal([]byte(shopJson), &qShop); err != nil {
			log.Printf("Failed to unmarshal qShop from Redis: %v", err)
		}
		return qShop, nil
	} else if shopJson == "" && err == nil {
		// 键值对存在，但是内容为空，说明是缓存的空数据
		return nil, nil
	}

	qId, _ := strconv.Atoi(id)
	// 查询数据库
	qShop, err = p.dbClient.Shop.Query().
		Where(shop.IDEQ(int64(qId))).
		Only(context.Background())
	if err != nil {
		// 缓存空值 解决缓存穿透
		client.Set(context.Background(), key, nil, tools.CacheShopTtl*time.Second)
		log.Printf("Failed to query qShop from database: %v", err)
		return nil, err
	}

	qShopBytes, err := json.Marshal(qShop)
	// 缓存到redis
	_, err = client.Set(context.Background(), key, qShopBytes, tools.CacheShopTtl*time.Minute).Result()
	if err != nil {
		log.Printf("Failed to set qShop to redis: %v", err)
	}
	return qShop, nil
}

// 互斥锁解决缓存击穿
/*func (p *pgImpl) queryWithMutex(ctx context.Context, id string) (*ent.Shop, error) {
	key := tools.CacheShopKey + id
	client := global.App.Redis
	// 查询店铺缓存
	shopJson, err := client.Get(context.Background(), key).Result()

	// 抛出除了找不到键值对以外的错误
	if err != nil && err != redis.Nil {
		log.Printf("Failed to get qShop from Redis: %v", err)
		return nil, err
	}
	// 缓存命中
	var qShop *ent.Shop
	if shopJson != "" {
		err := json.Unmarshal([]byte(shopJson), qShop)
		if err != nil {
			log.Printf("Failed to unmarshal qShop from Redis: %v", err)
			return nil, err
		}
		// 缓存命中则返回结果
		return qShop, nil
	}

	// 缓存不存在 实现缓存重构
	// 获取互斥锁
	lockKey := tools.LockShopKey + id
	isLock := tools.TryLock(ctx, lockKey)
	// 获取互斥锁失败
	if !isLock {
		// 失败则休眠重试
		time.Sleep(50 * time.Millisecond)
		return p.queryWithMutex(ctx, id)
	}
	qId, _ := strconv.Atoi(id)
	// 成功,查询数据库
	res, _ := p.dbClient.Shop.Query().
		Where(shop.IDEQ(int64(qId))).
		Only(context.Background())
	// 不存在, 返回空值
	if res == nil {
		client.Set(context.Background(), key, nil, tools.CacheNullTtl*time.Minute)
		return nil, nil
	}

	// 存在
	// 序列化
	marshal, _ := json.Marshal(res)
	// 保存到redis
	client.Set(context.Background(), key, marshal, tools.CacheShopTtl*time.Minute)
	defer tools.UnLock(key)
	return res, nil
}*/

// 使用逻辑过期解决缓存击穿
/*func (p *pgImpl) queryWithLogicalExpire(ctx context.Context, id string) (*ent.Shop, error) {
	key := tools.CacheShopKey + id
	client := global.App.Redis
	// 查询缓存
	shopJson, err := client.Get(ctx, key).Result()

	// 抛出除了找不到键值对以外的错误
	if err != nil && err != redis.Nil {
		return nil, nil
	}

	// 键值对存在，但是内容为空，说明是缓存的空数据
	if shopJson == "" && err == nil {
		return nil, nil
	}

	// 缓存命中
	// 反序列化json
	var redisData model.RedisData
	err = json.Unmarshal([]byte(shopJson), &redisData)
	if err != nil {
		return nil, err
	}
	var shopData ent.Shop
	if err := json.Unmarshal(redisData.Data, &shopData); err != nil {
		return nil, err
	}
	expireTime := redisData.ExpireTime
	// 判断是否过期
	if expireTime.After(time.Now()) {
		// 未过期直接返回
		return &shopData, nil
	}

	// 已过期，进行缓存重建
	lockKey := tools.LockShopKey + id
	isLock := tools.TryLock(ctx, lockKey)
	// 判断是否获取锁成功
	if isLock {
		go func() {
			qId, _ := strconv.Atoi(id)
			p.saveShop2Redis(ctx, int64(qId), 20)
		}()
		defer tools.UnLock(lockKey)
	}
	return &shopData, nil
}*/

// 缓存预热\重建缓存
func (p *pgImpl) saveShop2Redis(ctx context.Context, id int64, expireSeconds int64) {
	client := global.App.Redis
	// 查询店铺数据
	qShop, err := p.dbClient.Shop.Query().Where(shop.IDEQ(id)).Only(context.Background())
	if err != nil {
		log.Printf("Failed to query qShop from database: %v", err)
		return
	}
	shopJson, err := json.Marshal(qShop)
	if err != nil {
		log.Printf("Failed to format qShop to json: %v", err)
		return
	}

	// 封装逻辑过期时间
	var redisData model.RedisData
	redisData.Data = shopJson
	redisData.ExpireTime = time.Now().Add(time.Duration(expireSeconds) * time.Second)
	keyID := strconv.Itoa(int(id))
	key := tools.CacheShopKey + keyID
	marshal, _ := json.Marshal(redisData)
	client.Set(ctx, key, marshal, -1)
}

func (p *pgImpl) QueryShopByType(ctx context.Context, typeId int64, current int64, x float64, y float64) (res []*ent.Shop) {
	if x == 0 || y == 0 {
		// 不需要坐标查询
		res, _ = p.dbClient.Shop.
			Query().
			Where(shop.TypeID(uint64(typeId))).
			Offset(int(current-1) * 6).
			Limit(6).
			All(ctx)
	}
	return res
}
