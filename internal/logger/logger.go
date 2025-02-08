package logger

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
)

const (
	localEnv = "local"
	devEnv   = "dev"
	prodEnv  = "prod"
)

func InitLogger(env, logFilePath string) (*os.File, error) {
	var handler slog.Handler
	var logFile *os.File
	var err error

	switch env {
	case localEnv:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	case devEnv:
		logFile, err = setupLog(logFilePath)
		if err != nil {
			return nil, err
		}
		handler = slog.NewTextHandler(logFile, &slog.HandlerOptions{Level: slog.LevelInfo})
	case prodEnv:
		logFile, err = setupLog(logFilePath)
		if err != nil {
			return nil, err
		}
		handler = slog.NewTextHandler(logFile, &slog.HandlerOptions{Level: slog.LevelWarn})
	default:
		return nil, fmt.Errorf("неизвестный env: %s", env)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	return logFile, nil
}

func setupLog(logFilePath string) (*os.File, error) {
	dir := filepath.Dir(logFilePath)
	if dir != "" && dir != "." {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			if err = os.MkdirAll(dir, 0755); err != nil {
				return nil, fmt.Errorf("ошибка создания директории %s: %v", logFilePath, err)
			}
		}
	}
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии лог файла %s: %v", logFilePath, err)
	}

	return logFile, nil
}
