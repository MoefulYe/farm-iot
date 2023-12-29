package models

import (
	"time"
)

type CowInfo struct {
	Id     string     `json:"id"`
	BornAt time.Time  `json:"born_at"`
	DeadAt *time.Time `json:"dead_at,omitempty"`
	Reason string     `json:"reason,omitempty"`
	Parent string     `json:"parent,omitempty"`
}

type CowQuery struct {
	Page   int    `json:"page" form:"page" binding:"min=1"`
	Size   int    `json:"size" form:"size" binding:"min=0,max=40"`
	Filter string `json:"filter" form:"filter"`
}

type KillCowReq struct {
	Cows []string `json:"cows"`
}
