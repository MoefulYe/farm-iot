package main

import (
	_ "github.com/MoefulYe/farm-iot/iot-server/grpc"
	"github.com/MoefulYe/farm-iot/iot-server/router"
	. "github.com/MoefulYe/farm-iot/iot-server/server"
	"os"
)

func main() {
	signal := make(chan os.Signal, 1)
	done := make(chan struct{}, 1)
	router.RegisterRouter()
	go func() {
		<-signal
		Server.Disconnect(250)
		done <- struct{}{}
	}()
	<-done
	close(signal)
	close(done)
}
