package openai

// common.go defines common types used throughout the OpenAI API.

// Usage Represents the total token usage per request to OpenAI.
type Usage struct {
	PromptTokens            int                      `json:"prompt_tokens"`
	ImageTokens             int                      `json:"image_tokens"`
	CompletionTokens        int                      `json:"completion_tokens"`
	ThinkingTokens          int                      `json:"thinking_tokens"`
	TotalTokens             int                      `json:"total_tokens"`
	PromptTokensDetails     *PromptTokensDetails     `json:"prompt_tokens_details"`
	CompletionTokensDetails *CompletionTokensDetails `json:"completion_tokens_details"`
}

// CompletionTokensDetails Breakdown of tokens used in a completion.
type CompletionTokensDetails struct {
	AudioTokens              int `json:"audio_tokens"`
	ReasoningTokens          int `json:"reasoning_tokens"`
	AcceptedPredictionTokens int `json:"accepted_prediction_tokens"`
	RejectedPredictionTokens int `json:"rejected_prediction_tokens"`
}

// PromptTokensDetails Breakdown of tokens used in the prompt.
type PromptTokensDetails struct {
	AudioTokens          int `json:"audio_tokens"`
	CachedTokens         int `json:"cached_tokens"`
	CachedCreationTokens int `json:"cached_creation_tokens"` // anthropic specific
	CacheReadTokens      int `json:"cache_read_tokens"`      // anthropic specific
}
