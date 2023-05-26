package model

import (
	"encoding/json"
	"time"
)

type RedisData struct {
	ExpireTime time.Time
	Data       json.RawMessage
}
