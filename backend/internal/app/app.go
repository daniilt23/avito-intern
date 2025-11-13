package app

import (
	"avito/internal/config"
	"avito/internal/handlers"
	"avito/internal/repository/postgres"
	"avito/internal/repository/postgres/team"
	"avito/internal/repository/postgres/user"
	"avito/internal/service"
	mylog "avito/log"
	"log"
	"os"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Start() {
	log.Println("START APP")
	cfg := config.MustLoadConfig()

	logger := mylog.SetupLogger()
	logger.Info("logger init")

	db, err := postgres.ConnectDB(*cfg.Database)
	if err != nil {
		logger.Error("Cannot connect to db")
		os.Exit(1)
	}
	logger.Info("Connect to db")

	teamRepoSql := team.NewTeamRepoSQL(db)

	userRepoSQL := user.NewUserRepoSQL(db)
	logger.Debug("Create user repo")

	service := service.NewService(teamRepoSql, userRepoSQL, logger)
	logger.Debug("Create service")

	handler := handlers.NewHandler(service)
	logger.Debug("Create handler")

	srv := NewServer(*cfg, handler.InitRoutes())
	logger.Info("server data", "host", cfg.Server.Host, "port", cfg.Server.Port)

	if err = srv.Run(); err != nil {
		logger.Error("Cannot start server")
		os.Exit(1)
	}
}
