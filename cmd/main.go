package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/salesforceanton/files-portal/internal/config"
	"github.com/salesforceanton/files-portal/internal/logger"
	"github.com/salesforceanton/files-portal/internal/repository"
	"github.com/salesforceanton/files-portal/internal/server"
	"github.com/salesforceanton/files-portal/internal/service"
	handler "github.com/salesforceanton/files-portal/internal/transport/rest"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {

	// Initialize app configuration
	cfg, err := config.InitConfig()
	if err != nil {
		logger.LogServerIssue(err)
		return
	}

	// Connect to DB
	db, err := repository.NewPostgresDB(&cfg.DB)
	if err != nil {
		logger.LogServerIssue(err)
		return
	}

	// Set dependenties
	repos := repository.NewRepository(db)
	service := service.NewService(repos, cfg)
	handler := handler.NewHandler(service)
	server := new(server.Server)

	// Run server
	logrus.Info(fmt.Sprintf("SERVER STARTED: %s", time.Now().Local().String()))
	go func() {
		if err := server.Run(cfg.ServerPort, handler.InitRoutes()); err != nil {
			logger.LogServerIssue(err)
			return
		}
	}()

	// Gracefull shutdown
	exit := make(chan os.Signal)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	<-exit

	if err := server.Shutdown(context.Background()); err != nil {
		logger.LogServerIssue(err)
		return
	}

	if err := db.Close(); err != nil {
		logger.LogServerIssue(err)
		return
	}

	log.Info("Server shutdown successfully")
}
