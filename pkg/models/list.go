package models

import (
	"context"
	"encoding/json"

	"github.com/edony-ink/log"
	"github.com/sashabaranov/go-openai"

	mentis "github.com/edonyzpc/Mentis/pkg"
)

func ListModels(mentis *mentis.Mentis) ([]openai.Model, error) {
	models, err := mentis.Client.ListModels(context.Background())
	if err != nil {
		log.Errorf("list models error: %v\n", err)
		return nil, err
	}

	for _, model := range models.Models {
		data, _ := json.MarshalIndent(model, "", " ")
		log.Infof("%+s", data)
	}

	return models.Models, nil
}
