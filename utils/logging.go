package utils

import (
	logger "github.com/sirupsen/logrus"
)

func InitLogging() {
	logger.SetFormatter(&logger.JSONFormatter{})

	logger.Info("People API is live and running")
}
