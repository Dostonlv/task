package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Dostonlv/task/config"
	"github.com/Dostonlv/task/internal/server"
	"github.com/Dostonlv/task/pkg/db/postgres"
	"github.com/Dostonlv/task/pkg/logger"
	"github.com/Dostonlv/task/pkg/utils"
)

// @title Blog and News API.
// @version 1.0
// @description Blog and News API Server.
// @contact.name Doston Nematov (kei)
// @contact.url  https://github.com/Dostonlv
// @contact.telegram https://t.me/dostonlv
// @contact.email dostonlv@icloud.com
// @BasePath /v1
func main() {
	log.Println("Starting api server")

	configPath := utils.GetConfigPath(os.Getenv("config"))

	fmt.Println(configPath)

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	appLogger := logger.NewApiLogger(cfg)

	appLogger.InitLogger()
	appLogger.Infof("AppVersion: %s, LogLevel: %s, Mode: %s", cfg.Server.AppVersion, cfg.Logger.Level, cfg.Server.Mode)

	psqlDB, err := postgres.NewPsqlDB(cfg)
	if err != nil {
		appLogger.Fatalf("Postgresql init: %s", err)
	} else {
		appLogger.Infof("Postgres connected, Status: %#v", psqlDB.Stats())
	}
	defer psqlDB.Close()

	s := server.NewServer(cfg, psqlDB, appLogger)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
