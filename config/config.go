package config

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"gopkg.in/ini.v1"
	"sync"
)

type Config struct {
	App
	Mysql
	Wx
}

type (
	App struct {
		Addr     string `ini:"addr"`
		LogLevel string `ini:"log_level"`
		PageSize int    `ini:"page_size"`
	}

	Mysql struct {
		Connect string `ini:"connect"`
		MaxIdle int    `ini:"max_idle"`
		MaxOpen int    `ini:"max_open"`
	}

	Wx struct {
		AppId     string `ini:"app_id"`
		AppSecret string `ini:"app_secret"`
	}
)

var (
	Conf *Config
	once sync.Once
)

func NewConfig(env string) {
	once.Do(func() {
		cfg, err := ini.ShadowLoad(fmt.Sprintf("config/%s.ini", env))
		if err != nil {
			log.Fatal(err)
		}

		c := new(Config)
		err = cfg.MapTo(c)
		Conf = c
	})
}
