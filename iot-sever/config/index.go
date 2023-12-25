package config

import (
	. "github.com/MoefulYe/farm-iot/iot-server/logger"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
)

type Config struct {
	Broker   string `toml:"broker"`
	User     string `toml:"user"`
	Passwd   string `toml:"passwd"`
	Postgres string `toml:"postgres"`
	GrpcAddr string `toml:"grpc_addr"`
	Influxdb struct {
		Url      string `toml:"url"`
		Auth     string `toml:"auth"`
		Username string `toml:"username"`
		Passwd   string `toml:"passwd"`
	} `toml:"influxdb"`
}

func (c *Config) NewServerOpts() *mqtt.ClientOptions {
	return mqtt.NewClientOptions().AddBroker(c.Broker).SetUsername(c.User).
		SetPassword(c.Passwd).SetClientID("iot-server")
}

var Conf Config

func init() {
	viper.AddConfigPath("./")
	viper.SetConfigName("conf")
	viper.SetConfigType("toml")
	Logger.Infow("read config")
	err := viper.ReadInConfig()
	if err != nil {
		Logger.Fatal(err.Error())
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		Logger.Fatal(err.Error())
	}
}
