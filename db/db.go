package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

type Config struct {
	Host        string
	Port        uint16
	Username    string
	Password    string
	DbName      string
	DriverName  string
	MaxIdle     int
	MaxOpen     int
	MaxIdleTime int
	MaxLifeTime int
}

func NewDb(config *Config) (*sqlx.DB, error) {

	info := fmt.Sprintf("%s:%s@(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.DbName)
	db, err := sqlx.Open(config.DriverName, info)
	if err != nil {
		return nil, err
	}

	if config.MaxIdle != 0 {
		db.SetMaxIdleConns(config.MaxIdle)
	}
	if config.MaxIdleTime != 0 {
		db.SetConnMaxIdleTime(time.Duration(config.MaxIdleTime))
	}
	if config.MaxOpen != 0 {
		db.SetMaxOpenConns(config.MaxOpen)
	}
	if config.MaxLifeTime != 0 {
		db.SetConnMaxLifetime(time.Duration(config.MaxLifeTime))
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
