package model

import "time"

type Voucher struct {
	Id          int64     `json:"id,omitempty"`
	ShopId      uint64    `json:"shopId,omitempty"`
	Title       string    `json:"title,omitempty"`
	SubTitle    string    `json:"subTitle,omitempty"`
	Rules       string    `json:"rules,omitempty"`
	PayValue    uint64    `json:"payValue,omitempty"`
	ActualValue int64     `json:"actualValue,omitempty"`
	Type        int8      `json:"type"`
	Stock       uint64    `json:"stock,omitempty"`
	Status      int8      `json:"status,omitempty"`
	BeginTime   time.Time `json:"beginTime,omitempty"`
	EndTime     time.Time `json:"endTime,omitempty"`
}
