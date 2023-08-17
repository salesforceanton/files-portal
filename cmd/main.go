package main

import (
	"os"

	"github.com/salesforceanton/files-portal/internal/config"
	"github.com/salesforceanton/files-portal/internal/logger"
	"github.com/salesforceanton/files-portal/internal/repository"
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
		logger.LogExecutionIssue(err)
		return
	}

	// Connect to DB
	_, err = repository.NewPostgresDB(&cfg.DB)
	if err != nil {
		logger.LogExecutionIssue(err)
		return
	}
}
