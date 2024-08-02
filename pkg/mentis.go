package mentis

import "github.com/sashabaranov/go-openai"

const (
	System    = openai.ChatMessageRoleSystem
	User      = openai.ChatMessageRoleUser
	Assistant = openai.ChatMessageRoleAssistant
	Function  = openai.ChatMessageRoleFunction
	Tool      = openai.ChatMessageRoleTool
)

type Mentis struct {
	BaseURL string
	token   string
	Client  *openai.Client
}

func New(url, token string) Mentis {
	config := openai.DefaultConfig(token)
	config.BaseURL = url
	client := openai.NewClientWithConfig(config)

	return Mentis{
		BaseURL: url,
		token:   token,
		Client:  client,
	}
}
