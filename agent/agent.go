package agent

import (
	"context"
	// "errors"
	// "fmt"
	// "io"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"

	"calcagent/tools"
)

type Agent struct {
	ctx       context.Context
	chatModel model.ToolCallingChatModel
	Agent     *react.Agent

	err error
}

func NewAgent(ctx context.Context) (*Agent, error) {
	a := &Agent{
		ctx: ctx,
	}

	a.newChatModel()

	ins, err := react.NewAgent(a.ctx, &react.AgentConfig{
		ToolCallingModel: a.chatModel,
		ToolsConfig: compose.ToolsNodeConfig{
			Tools: tools.NewCalcTool(),
		},
		MessageModifier: func(ctx context.Context, input []*schema.Message) []*schema.Message {
			res := make([]*schema.Message, 0, len(input)+1)

			res = append(res, schema.SystemMessage(systemPrompt))
			res = append(res, input...)
			return res
		},
		// StreamToolCallChecker: func(ctx context.Context, sr *schema.StreamReader[*schema.Message]) (bool, error) {
		// 	defer sr.Close()
		// 	for {
		// 		msg, err := sr.Recv()
		// 		if err != nil {
		// 			if errors.Is(err, io.EOF) {
		// 				break
		// 			}
		// 			return false, err
		// 		}
		// 		if msg.Content != "" {
		// 			fmt.Printf("|->%v<-|", msg.Content)
		// 		}
		// 		if len(msg.ToolCalls) > 0 {
		// 			return true, nil
		// 		}
		// 	}
		// 	return false, nil
		// },
		MaxStep: 40,
	})
	if err != nil {
		return nil, err
	}

	a.Agent = ins
	return a, nil

}
