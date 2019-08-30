package db

import (
	"errors"
)

type DB interface {
	Query(dest interface{}, sql string) error
	Exec(sql string) error
	Close() error
}

func Open(driver, host string, port int, user string, password string, dbName string, maxIdleConns int) (DB, error) {
	switch driver {
	case "mysql":
		return newMySql(host, port, user, password, dbName, maxIdleConns)
	}
	return nil, errors.New("not find driver")
}
