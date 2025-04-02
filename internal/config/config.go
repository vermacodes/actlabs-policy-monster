package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	UseMSI                       bool
	ActlabsHubMSIClientID        string
	ActlabsHubSubscriptionID     string
	ActlabsHubResourceGroupName  string
	ActlabsHubStorageAccountName string
}

func NewConfig() (*Config, error) {
	// Load configuration from environment variables
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}

	useMSI, err := strconv.ParseBool(getEnv("USE_MSI"))
	if err != nil {
		return nil, fmt.Errorf("required environment variable use_msi is not set")
	}

	actlabsHubMSIClientID := getEnv("ACTLABS_HUB_MSI_CLIENT_ID")
	if actlabsHubMSIClientID == "" {
		return nil, fmt.Errorf("required environment variable ACTLABS_HUB_MSI_CLIENT_ID is not set")
	}

	actlabsHubSubscriptionID := getEnv("ACTLABS_HUB_SUBSCRIPTION_ID")
	if actlabsHubSubscriptionID == "" {
		return nil, fmt.Errorf("required environment variable ACTLABS_HUB_SUBSCRIPTION_ID is not set")
	}
	actlabsHubResourceGroupName := getEnv("ACTLABS_HUB_RESOURCE_GROUP_NAME")
	if actlabsHubResourceGroupName == "" {
		return nil, fmt.Errorf("required environment variable ACTLABS_HUB_RESOURCE_GROUP_NAME is not set")
	}
	actlabsHubStorageAccountName := getEnv("ACTLABS_HUB_STORAGE_ACCOUNT_NAME")
	if actlabsHubStorageAccountName == "" {
		return nil, fmt.Errorf("required environment variable ACTLABS_HUB_STORAGE_ACCOUNT_NAME is not set")
	}

	slog.Info("Loaded configuration from environment variables",
		slog.String("use_msi", strconv.FormatBool(useMSI)),
		slog.String("actlabs_hub_msi_client_id", actlabsHubMSIClientID),
		slog.String("actlabs_hub_subscription_id", actlabsHubSubscriptionID),
		slog.String("actlabs_hub_resource_group_name", actlabsHubResourceGroupName),
		slog.String("actlabs_hub_storage_account_name", actlabsHubStorageAccountName),
	)

	return &Config{
		UseMSI:                       useMSI,
		ActlabsHubMSIClientID:        actlabsHubMSIClientID,
		ActlabsHubSubscriptionID:     actlabsHubSubscriptionID,
		ActlabsHubResourceGroupName:  actlabsHubResourceGroupName,
		ActlabsHubStorageAccountName: actlabsHubStorageAccountName,
	}, nil
}

// Helper function to retrieve the value and log it
func getEnv(env string) string {
	value := os.Getenv(env)
	slog.Info("environment variable", slog.String("name", env), slog.String("value", value))
	return value
}

// Helper function to retrieve the value, if none found, set default and log it
// func getEnvWithDefault(env string, defaultValue string) string {
// 	value := os.Getenv(env)
// 	if value == "" {
// 		value = defaultValue
// 	}
// 	slog.Info("environment variable", slog.String("name", env), slog.String("value", value))
// 	return value
// }
