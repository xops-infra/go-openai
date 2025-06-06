package openai

import (
	"net/http"
	"regexp"
)

const (
	openaiAPIURLv1                 = "https://api.openai.com/v1"
	defaultEmptyMessagesLimit uint = 300

	azureAPIPrefix         = "openai"
	azureApiBaseURL        = ".openai.azure.com"
	azureDeploymentsPrefix = "deployments"
	azureDefaultApiVersion = "2023-08-01-preview"

	AnthropicAPIVersion = "2023-06-01"
)

type APIType string

const (
	APITypeOpenAI          APIType = "OPEN_AI"
	APITypeAzure           APIType = "AZURE"
	APITypeAzureAD         APIType = "AZURE_AD"
	APITypeCloudflareAzure APIType = "CLOUDFLARE_AZURE"
	APITypeAnthropic       APIType = "ANTHROPIC"
)

const AzureAPIKeyHeader = "api-key"

const defaultAssistantVersion = "v2" // upgrade to v2 to support vector store

type HTTPDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// ClientConfig is a configuration of a client.
type ClientConfig struct {
	authToken string

	BaseURL              string
	OrgID                string
	APIType              APIType
	APIVersion           string // required when APIType is APITypeAzure or APITypeAzureAD or APITypeAnthropic
	AssistantVersion     string
	AzureModelMapperFunc func(model string) string // replace model to azure deployment name func
	HTTPClient           HTTPDoer
	AzureResourceName    string

	EmptyMessagesLimit uint
}

func DefaultConfig(authToken string) ClientConfig {
	return ClientConfig{
		authToken:        authToken,
		BaseURL:          openaiAPIURLv1,
		APIType:          APITypeOpenAI,
		AssistantVersion: defaultAssistantVersion,
		OrgID:            "",

		HTTPClient: &http.Client{},

		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}
}

func DefaultAzureConfig(apiKey, baseURL string) ClientConfig {
	return ClientConfig{
		authToken:  apiKey,
		OrgID:      "",
		APIType:    APITypeAzure,
		APIVersion: "2023-05-15",
		AzureModelMapperFunc: func(model string) string {
			return regexp.MustCompile(`[.:]`).ReplaceAllString(model, "")
		},

		HTTPClient: &http.Client{},

		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}
}

func DefaultAzureConfigWithModelMapperFunc(apiKey, resource, apiVersion string, azureModelMapperFunc func(model string) string) ClientConfig {
	if apiVersion == "" {
		apiVersion = azureDefaultApiVersion
	}
	return ClientConfig{
		authToken:            apiKey,
		BaseURL:              azureApiBaseURL,
		AzureResourceName:    resource,
		APIType:              APITypeAzure,
		APIVersion:           apiVersion,
		AzureModelMapperFunc: azureModelMapperFunc,

		HTTPClient: &http.Client{},

		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}
}

func (c ClientConfig) WithHttpClientConfig(client *http.Client) ClientConfig {
	c.HTTPClient = client
	return c
}

func DefaultAnthropicConfig(apiKey, baseURL string) ClientConfig {
	if baseURL == "" {
		baseURL = "https://api.anthropic.com/v1"
	}
	return ClientConfig{
		authToken:  apiKey,
		BaseURL:    baseURL,
		OrgID:      "",
		APIType:    APITypeAnthropic,
		APIVersion: AnthropicAPIVersion,

		HTTPClient: &http.Client{},

		EmptyMessagesLimit: defaultEmptyMessagesLimit,
	}
}

func (ClientConfig) String() string {
	return "<OpenAI API ClientConfig>"
}

func (c ClientConfig) GetAzureDeploymentByModel(model string) string {
	if c.AzureModelMapperFunc != nil {
		return c.AzureModelMapperFunc(model)
	}

	return model
}
