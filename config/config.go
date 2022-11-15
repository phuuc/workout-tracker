package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/finnpn/workout-tracker/pkg/helpers"
	"github.com/finnpn/workout-tracker/pkg/log"
	"github.com/joho/godotenv"
)

type Config struct {
	Server *Server
	App    *App
	Mysql  *Mysql
}

type App struct {
	IsProduction bool
}

type Server struct {
	Host string `env:"host"`
	Port int    `env:"port"`
}

type Mysql struct {
	User   string `env:"user"`
	Passwd string `env:"passwd"`
	Net    string `env:"net"`
	Host   string `env:"host"`
	Port   int    `env:"port"`
	DbName string `env:"db_name"`
}

func NewConfig(app *App) *Config {
	cfg := &Config{
		App: app,
	}
	err := cfg.parseEnv()
	if err != nil {
		log.Error("shut down the program with err =%v", err)
		os.Exit(1)
	}
	return cfg
}

func (c *Config) parseEnv() (err error) {
	log.Info("loading env file ...")
	dir := helpers.RootDir()
	if c.App.IsProduction {
		err = godotenv.Load(fmt.Sprintf("%s/prod.env", dir))
	} else {
		err = godotenv.Load(fmt.Sprintf("%s/local.env", dir))
	}
	if err != nil {
		return errors.New("could not parse server env : " + err.Error())
	}
	c.Server.Port, err = c.stringToInt("API_PORT")
	if err != nil {
		return err
	}
	c.Mysql.Port, err = c.stringToInt("MYSQL_PORT")
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) stringToInt(port string) (int, error) {
	p, err := strconv.Atoi(os.Getenv(port))
	if err != nil {
		log.Error("could not parse int to string with err=%v", err)
		return 0, errors.New("could not parse port : " + err.Error())
	}
	return p, nil
}

func (c *Config) Addr(host string, port int) string {
	return fmt.Sprintf("%s:%d", host, port)
}
