package chat

import (
	"context"

	"github.com/sashabaranov/go-openai"

	mentis "github.com/edonyzpc/Mentis/pkg"
)

/*
// ChatCompletionRequest represents a request structure for chat completion API.
type ChatCompletionRequest struct {
	Model            string                        `json:"model"`
	Messages         []ChatCompletionMessage       `json:"messages"`
	MaxTokens        int                           `json:"max_tokens,omitempty"`
	Temperature      float32                       `json:"temperature,omitempty"`
	TopP             float32                       `json:"top_p,omitempty"`
	N                int                           `json:"n,omitempty"`
	Stream           bool                          `json:"stream,omitempty"`
	Stop             []string                      `json:"stop,omitempty"`
	PresencePenalty  float32                       `json:"presence_penalty,omitempty"`
	ResponseFormat   *ChatCompletionResponseFormat `json:"response_format,omitempty"`
	Seed             *int                          `json:"seed,omitempty"`
	FrequencyPenalty float32                       `json:"frequency_penalty,omitempty"`
	// LogitBias is must be a token id string (specified by their token ID in the tokenizer), not a word string.
	// incorrect: `"logit_bias":{"You": 6}`, correct: `"logit_bias":{"1639": 6}`
	// refs: https://platform.openai.com/docs/api-reference/chat/create#chat/create-logit_bias
	LogitBias map[string]int `json:"logit_bias,omitempty"`
	// LogProbs indicates whether to return log probabilities of the output tokens or not.
	// If true, returns the log probabilities of each output token returned in the content of message.
	// This option is currently not available on the gpt-4-vision-preview model.
	LogProbs bool `json:"logprobs,omitempty"`
	// TopLogProbs is an integer between 0 and 5 specifying the number of most likely tokens to return at each
	// token position, each with an associated log probability.
	// logprobs must be set to true if this parameter is used.
	TopLogProbs int    `json:"top_logprobs,omitempty"`
	User        string `json:"user,omitempty"`
	// Deprecated: use Tools instead.
	Functions []FunctionDefinition `json:"functions,omitempty"`
	// Deprecated: use ToolChoice instead.
	FunctionCall any    `json:"function_call,omitempty"`
	Tools        []Tool `json:"tools,omitempty"`
	// This can be either a string or an ToolChoice object.
	ToolChoice any `json:"tool_choice,omitempty"`
	// Options for streaming response. Only set this when you set stream: true.
	StreamOptions *StreamOptions `json:"stream_options,omitempty"`
	// Disable the default behavior of parallel tool calls by setting it: false.
	ParallelToolCalls any `json:"parallel_tool_calls,omitempty"`
}
*/

const (
	// according to [LLM Parameters](./parameters.md), set the default value of Mentis
	Model = "qwen-max"
	// qwen-max support 8k, but range of max_token should be [1, 2000]
	MaxTokens   = 2000
	Temperature = 0.3
	TopN        = 20
	TopP        = 0.8
	//PresencePenalty  = 1.0
	FrequencyPenalty = 1.0
)

func Create(mentis *mentis.Mentis, model string, messages []openai.ChatCompletionMessage) (openai.ChatCompletionResponse, error) {
	return mentis.Client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:       model,
		Messages:    messages,
		MaxTokens:   MaxTokens,
		Temperature: Temperature,
		TopP:        TopP,
		N:           TopN,
		//PresencePenalty: PresencePenalty,
	})
}

func CreateStream(mentis *mentis.Mentis, model string, messages []openai.ChatCompletionMessage) (*openai.ChatCompletionStream, error) {
	return mentis.Client.CreateChatCompletionStream(context.Background(), openai.ChatCompletionRequest{
		Model:    model,
		Messages: messages,
		Stream:   true,
		StreamOptions: &openai.StreamOptions{
			IncludeUsage: true,
		},
		MaxTokens:   MaxTokens,
		Temperature: Temperature,
		TopP:        TopP,
		N:           TopN,
		//PresencePenalty: PresencePenalty,
	})
}

func CustomeCreate(mentis *mentis.Mentis, req openai.ChatCompletionRequest) (openai.ChatCompletionResponse, error) {
	return mentis.Client.CreateChatCompletion(context.Background(), req)
}
