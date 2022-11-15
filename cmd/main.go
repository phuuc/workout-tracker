package main

import (
	"flag"

	"github.com/finnpn/workout-tracker/config"
	"github.com/finnpn/workout-tracker/infras"
	"github.com/finnpn/workout-tracker/server"
)

func main() {
	isProd := flag.Bool("is production", false, "to specify environment variable")
	flag.Parse()

	cfg := config.NewConfig(&config.App{
		IsProduction: *isProd,
	})
	server := server.NewRouter(cfg)
	server.Run()

	db := infras.NewDB(cfg)
	db.RunMysql()
}
