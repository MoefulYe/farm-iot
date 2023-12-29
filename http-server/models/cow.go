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

type KillCowReq struct {
	Cows []string `json:"cows"`
}
