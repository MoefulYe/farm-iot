package config

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Broker   string `toml:"broker"`
	User     string `toml:"user"`
	Passwd   string `toml:"passwd"`
	Postgres string `toml:"postgres"`
	Influxdb struct {
		Url  string `toml:"url"`
		Auth string `toml:"auth"`
	} `toml:"influxdb"`
}

func (c *Config) NewServerOpts() *mqtt.ClientOptions {
	return mqtt.NewClientOptions().AddBroker(c.Broker).SetUsername(c.User).
		SetPassword(c.Passwd)
}

var Conf Config

func init() {
	viper.AddConfigPath("./")
	viper.SetConfigName("conf")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(Conf)
}
