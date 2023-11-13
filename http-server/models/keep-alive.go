package models

type KeepAlive struct {
	Id        string  `json:"id"`
	Time      string  `json:"time"`
	Health    float64 `json:"health"`
	Weight    float64 `json:"weight"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type RangeQuery struct {
	Start string `form:"start"`
	Stop  string `form:"stop"`
}
