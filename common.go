package openai

// common.go defines common types used throughout the OpenAI API.

// Usage Represents the total token usage per request to OpenAI.
type Usage struct {
	PromptTokens            int                      `json:"prompt_tokens"`
	ImageTokens             int                      `json:"image_tokens,omitempty"`
	CompletionTokens        int                      `json:"completion_tokens"`
	ThinkingTokens          int                      `json:"thinking_tokens,omitempty"`
	TotalTokens             int                      `json:"total_tokens"`
	PromptTokensDetails     *PromptTokensDetails     `json:"prompt_tokens_details,omitempty"`
	CompletionTokensDetails *CompletionTokensDetails `json:"completion_tokens_details,omitempty"`
}

// CompletionTokensDetails Breakdown of tokens used in a completion.
type CompletionTokensDetails struct {
	AudioTokens              int `json:"audio_tokens,omitempty"`
	ReasoningTokens          int `json:"reasoning_tokens,omitempty"`
	AcceptedPredictionTokens int `json:"accepted_prediction_tokens,omitempty"`
	RejectedPredictionTokens int `json:"rejected_prediction_tokens,omitempty"`
}

// PromptTokensDetails Breakdown of tokens used in the prompt.
type PromptTokensDetails struct {
	AudioTokens          int `json:"audio_tokens,omitempty"`
	CachedTokens         int `json:"cached_tokens,omitempty"`
	CachedCreationTokens int `json:"cached_creation_tokens,omitempty"` // anthropic specific
	CacheReadTokens      int `json:"cache_read_tokens,omitempty"`      // anthropic specific
}
