package chat

import (
	"context"

	"github.com/sashabaranov/go-openai"

	mentis "github.com/edonyzpc/Mentis/pkg"
)

/*
from openai import OpenAI
import os

def get_response():
    client = OpenAI(
        api_key=os.getenv("DASHSCOPE_API_KEY"), # 如果您没有配置环境变量，请在此处用您的API Key进行替换
        base_url="https://dashscope.aliyuncs.com/compatible-mode/v1",  # 填写DashScope服务的base_url
    )
    completion = client.chat.completions.create(
        model="qwen-plus",
        messages=[{'role': 'system', 'content': 'You are a helpful assistant.'},
                  {'role': 'user', 'content': '你是谁？'}]
        )
    print(completion.model_dump_json())

if __name__ == '__main__':
    get_response()
*/

func Create(mentis *mentis.Mentis, model string, messages []openai.ChatCompletionMessage) (openai.ChatCompletionResponse, error) {
	return mentis.Client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:    model,
		Messages: messages,
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
	})
}
