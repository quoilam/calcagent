package tools

import (
	"context"
	"errors"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/sirupsen/logrus"
)

type Formula struct {
	Operand1  int    `json:"Operand1" jsonschema:"required,description=the first operand"`
	Operand2  int    `json:"Operand2" jsonschema:"required,description=the second operand"`
	Operation string `json:"Operation" jsonschema:"required,enum=add,enum=sub,enum=mul,enum=div"`
}

type Result struct {
	Result int `json:"result"`
}

func AddService(ctx context.Context, form *Formula) (*Result, error) {
	logrus.Infof("%d %s %d = %d\n", form.Operand1, form.Operation, form.Operand2, form.Operand1+form.Operand2)
	return &Result{Result: form.Operand1 + form.Operand2}, nil
}
func SubService(ctx context.Context, form *Formula) (*Result, error) {
	logrus.Infof("%d %s %d = %d\n", form.Operand1, form.Operation, form.Operand2, form.Operand1-form.Operand2)
	return &Result{Result: form.Operand1 - form.Operand2}, nil
}
func MulService(ctx context.Context, form *Formula) (*Result, error) {
	logrus.Infof("%d %s %d = %d\n", form.Operand1, form.Operation, form.Operand2, form.Operand1*form.Operand2)
	return &Result{Result: form.Operand1 * form.Operand2}, nil
}
func DivService(ctx context.Context, form *Formula) (*Result, error) {
	logrus.Infof("%d %s %d = %d\n", form.Operand1, form.Operation, form.Operand2, form.Operand1/form.Operand2)
	if form.Operand2 == 0 {
		return nil, errors.New("division by zero")
	}
	return &Result{Result: form.Operand1 / form.Operand2}, nil
}

func NewCalcTool() ([]tool.BaseTool) {
	addtool,err := utils.InferTool("add_tool", "calcs addition", AddService)
	if err != nil {
		logrus.Errorf("failed to create add tool: %v", err)
		return nil
	}
	
	subtool,err := utils.InferTool("sub_tool", "calcs subtraction", SubService)
	if err != nil {
		logrus.Errorf("failed to create sub tool: %v", err)
		return nil
	}

	multool,err := utils.InferTool("mul_tool", "calcs multiplication", MulService)
	if err != nil {
		logrus.Errorf("failed to create mul tool: %v", err)
		return nil
	}

	divtool,err := utils.InferTool("div_tool", "calcs division", DivService)
	if err != nil {
		logrus.Errorf("failed to create div tool: %v", err)
		return nil
	}
	return []tool.BaseTool{addtool, subtool, multool, divtool}	
}
