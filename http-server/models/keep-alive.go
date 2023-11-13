package models

import "time"

type KeepAlive struct {
	Id        string    `json:"id"`
	Time      time.Time `json:"time"`
	Health    float64   `json:"health,omitempty"`
	Weight    float64   `json:"weight,omitempty"`
	Latitude  float64   `json:"latitude,omitempty"`
	Longitude float64   `json:"longitude,omitempty"`
}

type RangeQuery struct {
	Start  string `form:"start"`
	Stop   string `form:"stop"`
	Fields string `form:"fields"`
}
