package main

import (
	"context"
	"log/slog"

	"github.com/vermacodes/actlabs-policy-monster/internal/auth"
	"github.com/vermacodes/actlabs-policy-monster/internal/azure"
	"github.com/vermacodes/actlabs-policy-monster/internal/config"
	"github.com/vermacodes/actlabs-policy-monster/logger"
)

func main() {
	logger.SetupLogger()
	appConfig, err := config.NewConfig()
	if err != nil {
		slog.Error("Error loading configuration", slog.String("error", err.Error()))
		panic(err)
	}

	authClient, err := auth.NewAuth(appConfig)
	if err != nil {
		slog.Error("Error creating auth", slog.String("error", err.Error()))
		panic(err)
	}

	azure, err := azure.NewAzure(appConfig, authClient)
	if err != nil {
		slog.Error("Error creating azure client", slog.String("error", err.Error()))
		panic(err)
	}
	ctx := context.Background()
	networkAccessDisabled, err := azure.IsNetworkAccessDisabled(ctx)
	if err != nil {
		slog.Error("Error checking network access", slog.String("error", err.Error()))
		panic(err)
	}
	if networkAccessDisabled {
		slog.Info("Network access is disabled")
		err = azure.EnablePublicNetworkAccess(ctx)
		if err != nil {
			slog.Error("Error enabling network access", slog.String("error", err.Error()))
			panic(err)
		}
		slog.Info("Network access is enabled")
	} else {
		slog.Info("Network access is enabled")
	}

	slog.Info("Executed Successfully")
}
