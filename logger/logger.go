package logger

import (
	"log/slog"
	"os"
	"strconv"
)

func SetupLogger() {
	logLevel := os.Getenv("ACTLABS_HUB_LOG_LEVEL")
	if logLevel == "" {
		slog.Info("ACTLABS_HUB_LOG_LEVEL not set will default to 0")
		logLevel = "0"
	}

	logLevelInt, err := strconv.Atoi(logLevel)
	if err != nil {
		slog.Error("Error converting ACTLABS_HUB_LOG_LEVEL to int will default to 0",
			slog.String("error", err.Error()),
		)
		logLevelInt = 0
	}

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.Level(logLevelInt),
	}

	slogHandler := slog.NewTextHandler(os.Stdout, opts)
	slog.SetDefault(slog.New(slogHandler))
}
