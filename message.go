package openai

import (
	"context"
	"fmt"
	"net/http"
)

type Message struct {
	ID          string         `json:"id"`
	Object      string         `json:"object"`
	CreatedAt   int64          `json:"created_at"`
	ThreadID    string         `json:"thread_id"`
	Role        string         `json:"role"`
	Content     any            `json:"content"`
	AssistantId string         `json:"assistant_id,omitempty"`
	RunID       string         `json:"run_id,omitempty"`
	FileIDs     []string       `json:"file_ids,omitempty"`
	Metadata    map[string]any `json:"metadata"`

	httpHeader
}

type MessageRequest ThreadMessage

type ModifyThreadMessageRequest struct {
	Metadata map[string]any `json:"metadata"`
}

// CreateThreadMessage creates a new thread.
func (c *Client) CreateThreadMessage(ctx context.Context, threadID string, request MessageRequest) (response Message, err error) {
	urlSuffix := fmt.Sprintf("/threads/%s/messages", threadID)
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(urlSuffix), withBody(request),
		withBetaAssistantV1())
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

// RetrieveThreadThreadMessage retrieves a thread.
func (c *Client) RetrieveThreadThreadMessage(ctx context.Context, threadID, messageID string) (response Message, err error) {
	urlSuffix := fmt.Sprintf("/threads/%s/messages/%s", threadID, messageID)

	req, err := c.newRequest(ctx, http.MethodGet, c.fullURL(urlSuffix),
		withBetaAssistantV1())
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}

// ModifyThreadMessage modifies a thread.
func (c *Client) ModifyThreadMessage(
	ctx context.Context,
	threadID, messageID string,
	request ModifyThreadMessageRequest,
) (response Message, err error) {

	urlSuffix := fmt.Sprintf("/threads/%s/messages/%s", threadID, messageID)
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(urlSuffix), withBody(request),
		withBetaAssistantV1())
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)
	return
}
