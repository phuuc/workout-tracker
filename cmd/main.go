package main

import (
	"github.com/finnpn/workout-tracker/config"
	"github.com/finnpn/workout-tracker/server"
)

func main() {
	cfg := config.NewConfig()
	server := server.NewRouter(cfg)
	server.Run()
}
