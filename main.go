package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/edony-ink/log"

	mentis "github.com/edonyzpc/Mentis/pkg"
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
	models, err := MENTIS.Client.ListModels(context.Background())
	if err != nil {
		fmt.Printf("list models error: %v\n", err)
		return
	}

	for _, model := range models.Models {
		data, _ := json.MarshalIndent(model, "", " ")
		log.Infof("%+s", data)
	}
}
