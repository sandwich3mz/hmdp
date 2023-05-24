package impl

import (
	"context"
	"encoding/json"
	"hmdp/ent"
	"hmdp/ent/shop"
	"hmdp/global"
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

func (p *pgImpl) QueryShopById(ctx context.Context, id string) *ent.Shop {
	client := global.App.Redis
	key := tools.CACHE_SHOP_KEY + id
	shopJson, err := client.Get(context.Background(), key).Result()
	if err != nil {
		log.Printf("Failed to get qShop from Redis: %v", err)
	}
	var qShop *ent.Shop
	if shopJson != "" {
		if err := json.Unmarshal([]byte(shopJson), qShop); err != nil {
			log.Printf("Failed to unmarshal qShop from Redis: %v", err)
		}
		return qShop
	}
	qId, _ := strconv.Atoi(id)
	qShop, err = p.dbClient.Shop.Query().
		Where(shop.IDEQ(int64(qId))).
		Only(context.Background())
	if err != nil {
		log.Printf("Failed to query qShop from database: %v", err)
	}

	client.Set(context.Background(), key, qShop, tools.CACHE_SHOP_TTL*time.Minute)

	return qShop
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
