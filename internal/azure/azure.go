package azure

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storage/armstorage"
	"github.com/vermacodes/actlabs-policy-monster/internal/auth"
	"github.com/vermacodes/actlabs-policy-monster/internal/config"
)

type Azure interface {
	IsNetworkAccessDisabled(ctx context.Context) (bool, error)
	EnablePublicNetworkAccess(ctx context.Context) error
}

type azure struct {
	auth      *auth.Auth
	appConfig *config.Config
}

func NewAzure(appConfig *config.Config, auth *auth.Auth) (Azure, error) {
	return &azure{
		appConfig: appConfig,
		auth:      auth,
	}, nil
}

// IsNetworkAccessDisabled checks if network access is disabled for a given storage account.
func (a *azure) IsNetworkAccessDisabled(ctx context.Context) (bool, error) {

	// Create a storage account client
	client, err := armstorage.NewAccountsClient(a.appConfig.ActlabsHubSubscriptionID, a.auth.Credentials, nil)
	if err != nil {
		return false, fmt.Errorf("failed to create storage accounts client: %w", err)
	}

	// Get the storage account properties
	account, err := client.GetProperties(ctx, a.appConfig.ActlabsHubResourceGroupName, a.appConfig.ActlabsHubStorageAccountName, nil)
	if err != nil {
		return false, fmt.Errorf("failed to get storage account: %w", err)
	}

	if account.Properties != nil && account.Properties.PublicNetworkAccess != nil {
		if string(*account.Properties.PublicNetworkAccess) == "Enabled" {
			slog.Info("PublicNetworkAccess is enabled")
			return false, nil
		}
	}

	slog.Info("PublicNetworkAccess is not enabled")
	return true, nil
}

// Enable public network access for a storage account.
func (a *azure) EnablePublicNetworkAccess(ctx context.Context) error {
	// Create a storage account client
	client, err := armstorage.NewAccountsClient(a.appConfig.ActlabsHubSubscriptionID, a.auth.Credentials, nil)
	if err != nil {
		return fmt.Errorf("failed to create storage accounts client: %w", err)
	}

	// Update the storage account properties
	_, err = client.Update(ctx, a.appConfig.ActlabsHubResourceGroupName, a.appConfig.ActlabsHubStorageAccountName, armstorage.AccountUpdateParameters{
		Properties: &armstorage.AccountPropertiesUpdateParameters{
			PublicNetworkAccess: to.Ptr(armstorage.PublicNetworkAccessEnabled),
		},
	}, nil)
	if err != nil {
		return fmt.Errorf("failed to update storage account: %w", err)
	}

	slog.Info("PublicNetworkAccess is enabled")
	return nil
}
