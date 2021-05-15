package util

import (
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)

	Logger = log
}
