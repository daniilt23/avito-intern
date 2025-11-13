package log

import (
	"log/slog"
	"os"
)

func SetupLogger() *slog.Logger {
	loggerType := os.Getenv("ENV")

	var logger *slog.Logger

	switch loggerType {
	case "prod":
		logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelInfo}))
	default:
		logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelDebug}))
	}

	return logger
}
