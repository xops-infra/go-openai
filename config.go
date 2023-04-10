package openai

import (
	"net/http"
)

const (
	openaiAPIURLv1                 = "https://api.openai.com/v1"
	defaultEmptyMessagesLimit uint = 300

	azureAPIPrefix         = "openai"
	azureApiBaseURL        = ".openai.azure.com"
	azureDeploymentsPrefix = "deployments"
	azureDefaultApiVersion = "2023-03-15-preview"
)

type APIType string

const (
	APITypeOpenAI  APIType = "OPEN_AI"
	APITypeAzure   APIType = "AZURE"
	APITypeAzureAD APIType = "AZURE_AD"
)

const AzureAPIKeyHeader = "api-key"

// ClientConfig is a configuration of a client.
type ClientConfig struct {
	authToken string

	BaseURL    string
	OrgID      string
	APIType    APIType
	APIVersion string // required when APIType is APITypeAzure or APITypeAzureAD
	EngineName string
	Engine     string // required when APIType is APITypeAzure or APITypeAzureAD

	HTTPClient *http.Client

	EmptyMessagesLimit uint
}

func DefaultConfig(authToken string) ClientConfig {
	return ClientConfig{
		authToken: authToken,
		BaseURL:   openaiAPIURLv1,
		APIType:   APITypeOpenAI,
		OrgID:     "",

		HTTPClient: &http.Client{},

		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}
}

func DefaultAzureConfig(apiKey, engineName, engine, apiVersion string) ClientConfig {

	if apiVersion == "" {
		apiVersion = azureDefaultApiVersion
	}

	return ClientConfig{
		authToken:  apiKey,
		OrgID:      "",
		APIType:    APITypeAzure,
		EngineName: engineName,
		APIVersion: apiVersion,
		Engine:     engine,

		HTTPClient: &http.Client{},

		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}
}

func (c ClientConfig) WithHttpClientConfig(client *http.Client) ClientConfig {
	c.HTTPClient = client
	return c
}
