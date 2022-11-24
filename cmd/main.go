package main

import (
	"flag"

	"github.com/finnpn/workout-tracker/config"
	"github.com/finnpn/workout-tracker/infras"
	"github.com/finnpn/workout-tracker/infras/repository"
	"github.com/finnpn/workout-tracker/interfaces/restapi/handler"
	"github.com/finnpn/workout-tracker/server"
	usecases "github.com/finnpn/workout-tracker/usecases/user"
)

func main() {
	configFile := flag.String("config-file", "", "config file path")
	flag.Parse()

	cfg, err := config.Load(*configFile)
	if err != nil {
		panic(err)
	}

	db := infras.NewDB(cfg)
	sqlDb := db.SetupMysql()
	db.RunMigration()

	userRepo := repository.NewUserRepository(sqlDb)

	userUc := usecases.NewAuthUserUc(userRepo)

	handler := handler.NewHandler(userUc)

	server := server.NewRouter(cfg, handler)

	server.Run()
}
