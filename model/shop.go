package model

import "time"

type Shop struct {
	id         int64     `json:"id"`
	name       string    `json:"name"`
	typeId     uint64    `json:"type_id"`
	images     string    `json:"images"`
	area       string    `json:"area"`
	address    string    `json:"address"`
	x          float64   `json:"x"`
	y          float64   `json:"y"`
	avgPrice   uint16    `json:"avg_price"`
	sold       uint64    `json:"sold"`
	comments   uint64    `json:"comments"`
	score      int8      `json:"score"`
	openHours  string    `json:"open_hours"`
	createTime time.Time `json:"create_time"`
	updateTime time.Time `json:"update_time"`
	distance   float64   `json:"distance"`
}
