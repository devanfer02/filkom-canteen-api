package log

import (
	"sync"
	"time"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type LogInfo map[string]interface{}

var (
	logger      *logrus.Logger
	once        sync.Once
	currLogDate string
)

func initLog() {
	logger = logrus.New()
}

func updateHook() {
	currDate := time.Now().Format("2006-01-02")

	if currDate == currLogDate {
		return
	}

	currLogDate = currDate
	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "./internal/data/logs/" + currDate + ".log",
		logrus.WarnLevel:  "./internal/data/logs/" + currDate + ".log",
		logrus.ErrorLevel: "./internal/data/logs/" + currDate + ".log",
	}

	logger.Hooks = make(logrus.LevelHooks)
	logger.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))
}

func getLogger() *logrus.Logger {
	once.Do(initLog)

	updateHook()

	return logger
}

func Info(fields LogInfo, info string) {
	log := getLogger()

	log.WithFields(logrus.Fields(fields)).Info(info)
}

func Warn(fields LogInfo, info string) {
	log := getLogger()

	log.WithFields(logrus.Fields(fields)).Warn(info)
}

func Fatal(fields LogInfo, info string) {
	log := getLogger()

	log.WithFields(logrus.Fields(fields)).Fatal(info)
}

func Error(fields LogInfo, info string) {
	log := getLogger()

	log.WithFields(logrus.Fields(fields)).Error(info)
}
