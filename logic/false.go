package logic

import "github.com/Knetic/govaluate"

type False struct {
	Base
}

func (this *False) Init(cfg interface{}) error {
	return nil
}

func (this *False) GetName() string {
	return "FALSE"
}

func (*False) GetDescription() string {
	return `FALSE函数返回布尔值false
用法：FALSE()
示例：略`
}

func (*False) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		return false, nil
	}
}
