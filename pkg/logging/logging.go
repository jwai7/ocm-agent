package logging

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

type level log.Level

var (
	defaultLogLevel = log.InfoLevel.String()
	logLevel        level
)

func (l *level) String() string {
	return log.Level(*l).String()
}

func (l *level) Set(value string) error {
	lvl, err := log.ParseLevel(strings.TrimSpace(value))
	if err == nil {
		*l = level(lvl)
	}
	return err
}

func initLogging() {
	log.SetLevel(log.Level(logLevel))
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		PadLevelText:  false,
	})
}

func init() {
	// Set default log level
	_ = logLevel.Set(defaultLogLevel)
}
