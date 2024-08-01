package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/edony-ink/log"
	"github.com/sashabaranov/go-openai"

	mentis "github.com/edonyzpc/Mentis/pkg"
	"github.com/edonyzpc/Mentis/pkg/chat"
	"github.com/edonyzpc/Mentis/pkg/models"
)

const (
	QWEN_URL = "https://dashscope.aliyuncs.com/compatible-mode/v1"
)

var (
	MENTIS mentis.Mentis
)

func init() {
	log.SWLogger.Init("/tmp/mentis/mentis.log", log.DebugLevel, true)

	t := os.Getenv("QWEN_APIKEY")
	if t == "" {
		buf, err := os.ReadFile("./env/mentis.env")
		if err != nil {
			log.Fatal("read env file error: ", err)
		}
		t = string(buf)
	}

	MENTIS = mentis.New(QWEN_URL, t)
}

func main() {
	models, err := models.ListModels(&MENTIS)
	if err != nil {
		log.Fatal("list models error: ", err)
	}

	for _, model := range models {
		data, _ := json.MarshalIndent(model, "", " ")
		log.Infof("%+s", data)
	}

	systemMessage := openai.ChatCompletionMessage{
		Role:    "system",
		Content: "你是一个助手",
	}
	userMessage := openai.ChatCompletionMessage{
		Role:    "user",
		Content: "你好",
	}
	messages := []openai.ChatCompletionMessage{
		systemMessage,
		userMessage,
	}
	resp, err := chat.Create(&MENTIS, "qwen-max", messages)
	if err != nil {
		log.Fatal("create chat error: ", err)
	}

	log.Info(resp.Choices[0].Message.Content)

	systemMessageTrans := openai.ChatCompletionMessage{
		Role:    "system",
		Content: "你是一个精通的中文翻译助手，请将我说的话翻译成英文",
	}

	var sentence string
	fmt.Println("请输入要翻译的内容: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		messagesTrans := []openai.ChatCompletionMessage{
			systemMessageTrans,
		}
		sentence = scanner.Text()
		userMessageTrans := openai.ChatCompletionMessage{
			Role:    "user",
			Content: sentence,
		}
		messagesTrans = append(messagesTrans, userMessageTrans)
		stream, err := chat.CreateStream(&MENTIS, "qwen-max", messagesTrans)
		if err != nil {
			log.Fatal("create stream chat error: ", err)
		}
		defer stream.Close()

		fmt.Println("Stream response: ")
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Println("\nStream finished")
				break
			}

			if err != nil {
				fmt.Printf("\nStream error: %v\n", err)
				break
			}

			if len(response.Choices) > 0 {
				fmt.Printf(response.Choices[0].Delta.Content)
			}
		}
		fmt.Print("\n> ")
	}

}
