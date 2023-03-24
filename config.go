package openai

import (
	"net/http"
)

const (
	apiURLv1                       = "https://api.openai.com/v1"
	azureApiURLv1                  = ".openai.azure.com/openai/deployments/"
	defaultEmptyMessagesLimit uint = 300
	defaultApiVersion              = "2023-03-15-preview"
)

type ClientConfig struct {
	authToken string

	HTTPClient *http.Client

	BaseURL string
	OrgID   string

	EmptyMessagesLimit uint

	// azure
	ResourceName   string
	DeploymentName string
	ApiVersion     string
}

func DefaultConfig(authToken string) ClientConfig {
	return ClientConfig{
		HTTPClient: &http.Client{},
		BaseURL:    apiURLv1,
		OrgID:      "",
		authToken:  authToken,

		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}
}
