package utils

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func InitLogger() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)
}

func LogError(message string, err error) {
	log.WithFields(logrus.Fields{
		"error": err.Error(),
	}).Error(message)
}

func LogInfo(message string) {
	log.Info(message)
}
