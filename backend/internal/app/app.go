package app

import (
	"avito/internal/config"
	"avito/internal/handlers"
	"avito/internal/repository/postgres"
	pullrequest "avito/internal/repository/postgres/pull_request"
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

	teamRepoSql := team.NewTeamRepoSQL(db, logger)
	logger.Info("Create user repo")

	userRepoSQL := user.NewUserRepoSQL(db, logger)
	logger.Info("Create user repo")

	pullRequestRepoSQL := pullrequest.NewPullRequestRepoSQL(db, logger)

	service := service.NewService(teamRepoSql, userRepoSQL, pullRequestRepoSQL, logger)
	logger.Info("Create service")

	handler := handlers.NewHandler(service)
	logger.Info("Create handler")

	srv := NewServer(*cfg, handler.InitRoutes())
	logger.Info("server data", "host", cfg.Server.Host, "port", cfg.Server.Port)

	err = srv.Run()
	if err != nil {
		logger.Error("Cannot start server")
		os.Exit(1)
	}
}
