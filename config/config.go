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
	Log    *log.Log
	App    *App
}

type App struct {
	IsProduction bool
}

type Server struct {
	Host string `env:"HOST"`
	Port int    `env:"PORT"`
}

func NewConfig(app *App) *Config {
	l := log.NewLog()
	cfg := &Config{
		Log: l,
		App: app,
	}
	s, err := cfg.parseServerEnv()
	if err != nil {
		l.Error("shut down the program with err =%v", err)
		os.Exit(1)
	}
	cfg.Server = s
	return cfg
}

func (c *Config) parseServerEnv() (server *Server, err error) {
	c.Log.Info("loading server env file ...")
	dir := helpers.RootDir()
	if c.App.IsProduction {
		err = godotenv.Load(fmt.Sprintf("%s/prod.env", dir))
	} else {
		err = godotenv.Load(fmt.Sprintf("%s/local.env", dir))
	}
	if err != nil {
		return nil, errors.New("could not parse server env : " + err.Error())
	}
	port, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		c.Log.Error("could not parse int to string with err=%v", err)
		return nil, errors.New("could not parse port : " + err.Error())
	}
	return &Server{
		Host: os.Getenv("API_HOST"),
		Port: port,
	}, nil
}

func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}
