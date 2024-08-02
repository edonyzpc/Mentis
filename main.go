package main

import (
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
	log.Infof("you got %d models", len(models))

	systemMessage := openai.ChatCompletionMessage{
		Role:    mentis.System,
		Content: "你是一个助手",
	}
	userMessage := openai.ChatCompletionMessage{
		Role:    mentis.User,
		Content: "你好，介绍一下你自己吧！",
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

	/*
		systemMessageTrans := openai.ChatCompletionMessage{
			Role:    "system",
			Content: "你是一个精通的中文翻译助手，请将我说的话翻译成英文",
		}
		chat.CreateChatBot("\n我是翻译机器人，请输入要翻译的内容: ", &MENTIS, systemMessageTrans)
	*/

	// init attack vector TTP incident
	attackVectorIncidentSystem := openai.ChatCompletionMessage{
		Role:    mentis.System,
		Content: "You are a cybersecurity expert. Your task is to produce a comprehensive incident response testing scenario based on the information provided.",
	}
	attackVectorIncidentUser := openai.ChatCompletionMessage{
		Role: mentis.User,
		Content: `**Background information:**
The company operates in the 'Technology / IT' industry and is of size 'Large Enterprise (10,000+ employees)'. 

**Threat actor information:**
Threat actor group 'APT16' is planning to target the company using the following kill chain
Resource Development: Server (T1584.004)
**Your task:**
Create an incident response testing scenario based on the information provided. The goal of the scenario is to test the company's incident response capabilities against the identified threat actor group. 
Your response should be well structured and formatted using Markdown. Write in British English.`,
	}
	attackMessages := []openai.ChatCompletionMessage{
		attackVectorIncidentSystem,
		attackVectorIncidentUser,
	}
	resp, err = chat.Create(&MENTIS, "qwen-max", attackMessages)
	if err != nil {
		log.Fatal("create chat error: ", err)
	}
	incident := resp.Choices[0].Message.Content
	log.Info(incident)

	// init attack vector TTP evaluate
	/*
		system := "You are an AI assistant that helps users update and ask questions about their incident response scenario. Only respond to questions or requests relating to the scenario, or incident response testing in general. Here is the scenario that the user previously generated:\n" + incident + "\n"
		attackVectorEvaluateSystem := openai.ChatCompletionMessage{
			Role:    mentis.System,
			Content: system,
		}
		chat.CreateChatBot("\nStart TTP evaluation: ", &MENTIS, attackVectorEvaluateSystem)
	*/
}
