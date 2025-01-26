package config

import "github.com/bigkevmcd/go-configparser"

type redisConfig struct {
	Host     string
	Port     string
	Password string
	DB       string
}

func getRedisConfig(p *configparser.ConfigParser) *redisConfig {
	rc := &redisConfig{}
	rc.Host, _ = p.Get("redis", "host")
	rc.Port, _ = p.Get("redis", "port")
	rc.Password, _ = p.Get("redis", "password")
	rc.DB, _ = p.Get("redis", "db")
	return rc
}
