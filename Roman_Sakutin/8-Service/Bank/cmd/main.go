package main

import (
	"context"
	"log"

	"go.uber.org/zap"

	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/config"
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/controller"
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/logic"
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/repo"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err.Error())
	}

	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	pool, err := repo.CreateDatabasePoolConnections(ctx, cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}

	rep := repo.NewRepo(pool)
	lgc := logic.NewService(logger, rep)
	api := controller.NewApp(cfg, lgc, logger, controller.NewUser(lgc), controller.NewAccount(lgc), controller.NewTransfer(lgc))
	logger.Info("start server")
	api.StartServe()
}
