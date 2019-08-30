package global

import (
	"github.com/BurntSushi/toml"
	"jephy.cn/lib/db"
	"path/filepath"
	"sync"
)

var (
	global *GlobalConfig
	once   sync.Once
)

func Config() *GlobalConfig {
	return global
}

func LoadConfig(path string) error {
	once.Do(func() {
		global = new(GlobalConfig)
		filePath, err := filepath.Abs(path)
		if err != nil {
			panic(err)
		}

		if _, err := toml.DecodeFile(filePath, global); err != nil {
			panic(err)
		}

		global.DB, err = db.Open(global.DataBase.Driver, global.DataBase.Host, global.DataBase.Port, global.DataBase.User, global.DataBase.Password, global.DataBase.DBName, global.DataBase.MaxIdleConns)
		if err != nil {
			panic(err)
		}
	})
	return nil
}
