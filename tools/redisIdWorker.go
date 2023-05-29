package tools

import (
	"context"
	"hmdp/global"
	"time"
)

const (
	CountBits      = 32
	BeginTimestamp = 1640995200
)

// NextId 生成订单id
func NextId(keyPrefix string) int64 {
	client := global.App.Redis
	// 生成时间戳
	now := time.Now().UTC()
	nowSecond := now.Unix()
	timeStamp := nowSecond - BeginTimestamp

	// 生成序列号
	// 获取日期
	date := now.Format("2006:01:02")
	// 自增
	ctx := context.Background()
	count, _ := client.Incr(ctx, "icr:"+keyPrefix+":"+date).Uint64()

	return (timeStamp << CountBits) | int64(count)
}
