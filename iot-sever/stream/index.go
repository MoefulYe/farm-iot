package stream

import (
	"github.com/google/uuid"
	ext "github.com/reugn/go-streams/extension"
	"time"
)

var (
	source = make(chan any, 1024)
)

type KeepAlive struct {
	Uuid      uuid.UUID
	Timestamp time.Time
	Weight    float64
	Health    float64
	Geo       struct {
		Latitude  float64
		Longitude float64
	}
}

func init() {
	ext.NewChanSource(source)
}

func StreamInput() chan<- any {
	return source
}
