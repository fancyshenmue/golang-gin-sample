package loghandle

import (
	log "github.com/sirupsen/logrus"
)

// LogInfo | Log Info
func LogInfo(function, action, info string) {
	log.WithFields(log.Fields{
		"Function": function,
		"Action":   action,
	}).Info(info)
}

// ErrorHandle | Error Handle
func ErrorHandle(function, action, info string, err error) {
	if err != nil {
		log.WithFields(log.Fields{
			"Function": function,
			"Action":   action,
		}).Infof("%s: %v", info, err)
	}
}
