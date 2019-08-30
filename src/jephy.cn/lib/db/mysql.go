package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"reflect"
	"time"
)

var (
	DefaultMaxOpenConns      = 800
	DefaultConnectWaiTimeout = 15 * time.Second
	DefaultCharset           = "utf8"
)
var (
	ErrNoRows = sql.ErrNoRows
)

type MySql struct {
	*sqlx.DB
}

func newMySql(host string, port int, user string, password string, dbName string, maxIdleConns int) (*MySql, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=%s&charset=%s",
		user, password, host, port, dbName, DefaultConnectWaiTimeout.String(), DefaultCharset)
	sqlxDB, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}

	sqlxDB.SetMaxIdleConns(maxIdleConns)
	sqlxDB.SetMaxOpenConns(DefaultMaxOpenConns)

	return &MySql{
		DB: sqlxDB,
	}, nil
}

func (db *MySql) Query(dest interface{}, sql string) error {
	var err error
	t := reflect.TypeOf(dest)
	if reflectx.Deref(t).Kind() == reflect.Slice {
		err = db.DB.Select(dest, sql)
	} else {
		err = db.DB.Get(dest, sql)
	}
	return err
}

func (db *MySql) Exec(sql string) error {
	_, err := db.DB.Exec(sql)
	return err
}

func (db *MySql) Close() error {
	return db.DB.Close()
}
