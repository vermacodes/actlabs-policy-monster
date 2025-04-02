package auth

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/vermacodes/actlabs-policy-monster/internal/config"
)

type Auth struct {
	Credentials azcore.TokenCredential
}

func NewAuth(appConfig *config.Config) (*Auth, error) {
	var cred azcore.TokenCredential
	var err error

	if appConfig.UseMSI {
		cred, err = azidentity.NewManagedIdentityCredential(&azidentity.ManagedIdentityCredentialOptions{
			ID: azidentity.ClientID(appConfig.ActlabsHubMSIClientID),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to create managed identity credential: %v", err)
		}
	} else {
		cred, err = azidentity.NewDefaultAzureCredential(nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create default azure credential: %v", err)
		}
	}

	return &Auth{
		Credentials: cred,
	}, nil
}
