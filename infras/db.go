package infras

import (
	"database/sql"
	"time"

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

func (d *DB) RunMysql() *sql.DB {
	log.Info("running mysql db...")
	cfg := mysql.Config{
		User:      d.config.Mysql.UserName,
		Passwd:    d.config.Mysql.Passwd,
		Net:       "tcp",
		Addr:      d.config.Addr(d.config.Mysql.Host, d.config.Mysql.Port),
		DBName:    d.config.Mysql.DbName,
		ParseTime: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Error("db could not open with err=%v", err)
		return nil
	}
	err = db.Ping()
	if err != nil {
		log.Error("db could not ping with err=%v", err)
		return nil
	}
	db.SetMaxIdleConns(d.config.Mysql.MaxIdleConns)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Duration(d.config.Mysql.ConnMaxLifeTimeMiliseconds))

	log.Info("connected")
	return db
}
