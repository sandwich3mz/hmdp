package model

import "time"

type Voucher struct {
	Id          int64     `json:"id"`
	ShopId      uint64    `json:"shop_id"`
	Title       string    `json:"title"`
	SubTitle    string    `json:"sub_title"`
	Rules       string    `json:"rules"`
	PayValue    uint64    `json:"pay_value"`
	ActualValue int64     `json:"actual_value"`
	Type        int8      `json:"type"`
	Status      int8      `json:"status"`
	CreateTime  time.Time `json:"create_time"`
	UpdateTime  time.Time `json:"update_time"`
	Stock       uint64    `json:"stock"`
	BeginTime   time.Time `json:"begin_time"`
	EndTime     time.Time `json:"end_time"`
}
