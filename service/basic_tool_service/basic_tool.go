package basic_tool_service

import (
	"chatai/models"
	"chatai/pkg/logging"
	"chatai/pkg/setting"
	"context"
	"github.com/sashabaranov/go-openai"
	"strings"
	"time"
)

type BasicTool struct {
	ID            int
	Name         string
	Desc          string
	Prompt       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type BasicToolCardInfo struct {
	Name         string `json:"name"`
	Desc          string `json:"description"`
}

func GetBasicToolInfoList() ([]BasicToolCardInfo, error) {
	res, err := models.GetAllBasicTools()
	if err != nil {
		logging.Error("error: %v", err)
		return nil, err
	}
	basicCardInfoList := make([]BasicToolCardInfo, 0)
	for i := range res {
		basicCardInfoList = append(basicCardInfoList, BasicToolCardInfo{
			Name: res[i].Name,
			Desc: res[i].Desc,
		})
	}
	return basicCardInfoList, nil
}

func GetToolInfo(name string) (*models.BasicTool, error) {
	res, err := models.GetBasicToolByName(name)
	if err != nil {
		logging.Error("error: %v", err)
		return nil, err
	}
	return res, nil
}

func PostChatGPT(prompt, content string) (string, error) {
	APIKey := setting.OpenAISetting.APIKey
	client := openai.NewClient(APIKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: prompt+content,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}
	return strings.TrimLeft(resp.Choices[0].Message.Content, "\n"), nil
}