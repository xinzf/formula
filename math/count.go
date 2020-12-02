package math

import (
	"github.com/Knetic/govaluate"
)

type Count struct {
	Base
}

func (this *Count) Init(cfg interface{}) error {
	return nil
}

func (this *Count) GetName() string {
	return "COUNT"
}

func (*Count) GetDescription() string {
	return `COUNT(value1, [value2], …)
统计参数个数。`
}

func (this *Count) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		return len(arguments), nil
	}
}
