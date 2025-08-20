package main

import (
	"calcagent/agent"
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/cloudwego/eino/schema"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	godotenv.Load()
	ctx := context.Background()
	agent, err := agent.NewAgent(ctx)
	if err != nil {
		logrus.Fatalf("failed to create agent: %v", err)
	}

	historyMessages := make([]*schema.Message, 0)
	stream := false

	
	for {
		var userinput string
		fmt.Printf("\nUser input : ")
		_, _ = fmt.Scan(&userinput)
		historyMessages = append(historyMessages, schema.UserMessage(userinput))

		if !stream {
			// 非流式输出
			modelOutputs, err := agent.Agent.Generate(ctx, historyMessages)
			if err != nil {
				logrus.Errorf("failed to generate messages: %v", err)
				continue
			}
			fmt.Printf("Model Outputs: %s\n", modelOutputs.Content)
			historyMessages = append(historyMessages, modelOutputs)
		} else {
			// 流式输出
			sr, err := agent.Agent.Stream(ctx, historyMessages)
			if err != nil {
				logrus.Fatalf("failed to stream messages: %v", err)
			}

			modelOutputs := schema.Message{
				Role: schema.Assistant,
			}
			for {
				msg, err := sr.Recv()
				if err != nil {
					if errors.Is(err, io.EOF) {
						logrus.Infoln("EOF")
						break
					} else {
						logrus.Errorf("failed to receive message: %v", err)
					}
				}

				fmt.Print(msg.Content)
				modelOutputs.Content += msg.Content
			}

			historyMessages = append(historyMessages, &modelOutputs)

		}
	}

}
