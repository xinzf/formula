package logic

import "github.com/Knetic/govaluate"

type True struct {
	Base
}

func (this *True) Init(cfg interface{}) error {
	return nil
}

func (this *True) GetName() string {
	return "TRUE"
}

func (*True) GetDescription() string {
	return `TRUE函数返回布尔值true
用法：TRUE()
示例：略`
}

func (*True) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		return true, nil
	}
}
