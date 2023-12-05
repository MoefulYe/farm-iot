package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Postgres   string `toml:"postgres"`
	Port       string `toml:"port"`
	GrpcServer string `toml:"grpc_server"`
	Influxdb   struct {
		Url      string `toml:"url"`
		Auth     string `toml:"auth"`
		Username string `toml:"username"`
		Passwd   string `toml:"passwd"`
	} `toml:"influxdb"`
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
}
