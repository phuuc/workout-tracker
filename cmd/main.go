package main

import (
	"flag"

	"github.com/finnpn/workout-tracker/config"
	"github.com/finnpn/workout-tracker/infras"
	"github.com/finnpn/workout-tracker/server"
)

func main() {
	configFile := flag.String("config-file", "", "config file path")
	flag.Parse()

	cfg, err := config.Load(*configFile)
	if err != nil {
		panic(err)
	}

	server := server.NewRouter(cfg)
	db := infras.NewDB(cfg)

	db.RunMysql()
	db.RunMigration()
	server.Run()
}
