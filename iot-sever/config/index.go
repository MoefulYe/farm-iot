package config

import mqtt "github.com/eclipse/paho.mqtt.golang"

type Config struct {
	Broker    string
	User      string
	Passwd    string
	PgConnStr string
}

func (c *Config) NewServerOpts() *mqtt.ClientOptions {
	return mqtt.NewClientOptions().AddBroker(c.Broker).SetUsername(c.User).
		SetPassword(c.Passwd)
}

var Conf Config

func init() {
	Conf.Broker = "tcp://0.0.0.0:1883"
	Conf.User = "admin"
	Conf.Passwd = "admin"
	Conf.PgConnStr = "host=127.0.0.1 port=5432 user=postgres dbname=farm-iot password=123456"
}