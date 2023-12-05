package models

import (
	"time"
)

type DeviceInfo struct {
	Id     string     `json:"id"`
	BornAt time.Time  `json:"born_at"`
	DeadAt *time.Time `json:"dead_at,omitempty"`
	Reason *string    `json:"reason,omitempty"`
}

type KillReq struct {
	Uuid string
}

type KillResp struct{}
