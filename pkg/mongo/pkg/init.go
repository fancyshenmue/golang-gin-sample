package pkg

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	formatter := &log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	log.SetFormatter(formatter)
}
