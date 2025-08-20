package agent

import (
	"os"

	"github.com/cloudwego/eino-ext/components/model/openai"
)

func (a *Agent) newChatModel() {
	config := &openai.ChatModelConfig{
		Model:   os.Getenv("OPENAI_MODEL"),
		BaseURL: os.Getenv("OPENAI_API_ENDPOINT"),
		APIKey:  os.Getenv("OPENAI_API_KEY"),
	}
	cm, err := openai.NewChatModel(a.ctx, config)
	if err != nil {
		a.err = err
		return
	}

	a.chatModel = cm
}
