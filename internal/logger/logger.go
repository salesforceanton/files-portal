package logger

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func LogHandlerIssue(handler string, err error) {
	logrus.WithFields(logrus.Fields{
		"handler": handler,
		"problem": fmt.Sprintf("Error appeared in handler [%s]: %s", handler, err.Error()),
	}).Error(err)
}

func LogServerIssue(err error) {
	logrus.WithFields(logrus.Fields{
		"point":   "server",
		"problem": fmt.Sprintf("Server running error: %s", err.Error()),
	}).Error(err)
}
