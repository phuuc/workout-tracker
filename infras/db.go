package infras

import (
	"database/sql"
	"time"

	"github.com/finnpn/workout-tracker/config"
	"github.com/finnpn/workout-tracker/pkg/helpers"
	"github.com/finnpn/workout-tracker/pkg/log"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type DB struct {
	config *config.Config
}

func NewDB(config *config.Config) *DB {
	return &DB{
		config: config,
	}
}

func (d *DB) SetupMysql() *sql.DB {
	log.Info("running mysql db...")
	var err error
	db, err := sql.Open(d.config.Mysql.DriverName, d.mysqlConfig())
	if err != nil {
		log.Error("db could not open with err=%v", err)
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Error("db could not ping with err=%v", err)
		panic(err)
	}
	db.SetMaxIdleConns(d.config.Mysql.MaxIdleConns)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(time.Duration(d.config.Mysql.ConnMaxLifeTimeMiliseconds))

	log.Info("connected...")
	return db
}

func (d *DB) RunMigration() {
	m, err := migrate.New("file://"+helpers.RootDir()+"/db/migration", "mysql://wt:wt@tcp(127.0.0.1:5436)/workout-tracker?multiStatements=true&&parseTime=true")
	if err != nil {
		log.Error("failed run migration with err=%v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Error("failed run migration with err=%v", err)
	}
	log.Info("migration completed!")
}

func (d *DB) mysqlConfig() string {
	cfg := &mysql.Config{
		User:            d.config.Mysql.UserName,
		Passwd:          d.config.Mysql.Passwd,
		Net:             "tcp",
		Addr:            d.config.Addr(d.config.Mysql.Host, d.config.Mysql.Port),
		DBName:          d.config.Mysql.Name,
		ParseTime:       true,
		MultiStatements: true,
	}
	return cfg.FormatDSN()
}
