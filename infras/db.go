package infras

import (
	"database/sql"

	"github.com/finnpn/workout-tracker/config"
	"github.com/finnpn/workout-tracker/pkg/log"
	"github.com/go-sql-driver/mysql"
)

type DB struct {
	config *config.Config
}

func NewDB(config *config.Config) *DB {
	return &DB{
		config: config,
	}
}

func (d *DB) RunMysql() (*sql.DB, error) {
	log.Info("running mysql db...")
	cfg := mysql.Config{
		User:   d.config.Mysql.User,
		Passwd: d.config.Mysql.Passwd,
		Net:    "tcp",
		Addr:   d.config.Addr(d.config.Mysql.Host, d.config.Mysql.Port),
		DBName: d.config.Mysql.DbName,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Error("db could not open with err=%v", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		log.Error("db could not ping with err=%v", err)
		return nil, err
	}
	log.Info("connected")
	return db, nil
}
