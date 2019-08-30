package global

import (
	"jephy.cn/lib/db"
)

type GlobalConfig struct {
	HTTPAddr string    `toml:"http_addr"`
	LogLevel string    `toml:"log_level"`
	DataBase *DataBase `toml:"database"`

	DB db.DB
}

type DataBase struct {
	Driver       string `toml:"driver"`
	Host         string `toml:"host"`
	Port         int    `toml:"port"`
	User         string `toml:"user"`
	Password     string `toml:"password"`
	DBName       string `toml:"db_name"`
	MaxIdleConns int    `toml:"max_idle_conns"`
}
