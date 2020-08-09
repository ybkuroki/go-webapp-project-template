package main

import (
	"github.com/labstack/echo/v4"
	"github.com/ybkuroki/go-webapp-project-template/config"
	"github.com/ybkuroki/go-webapp-project-template/logger"
	"github.com/ybkuroki/go-webapp-project-template/migration"
	"github.com/ybkuroki/go-webapp-project-template/repository"
	"github.com/ybkuroki/go-webapp-project-template/router"
	"github.com/ybkuroki/go-webapp-project-template/session"
)

func main() {
	e := echo.New()

	config.Load()
	logger.InitLogger(e, config.GetConfig())
	e.Logger.Info("Loaded this configuration : application." + *config.GetEnv() + ".yml")

	repository.InitDB()
	db := repository.GetDB()

	migration.CreateDatabase(config.GetConfig())
	migration.InitMasterData(config.GetConfig())

	router.Init(e, config.GetConfig())
	session.Init(e, config.GetConfig())
	if err := e.Start(":8080"); err != nil {
		e.Logger.Error(err)
	}

	defer db.Close()
}
