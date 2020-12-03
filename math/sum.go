package math

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
)

type Sum struct {
	Base
}

func (this *Sum) Init(cfg interface{}) error {
	return nil
}

func (this *Sum) GetName() string {
	return "SUM"
}

func (*Sum) GetDescription() string {
	return `SUM(number1, [number2], …)
函数使所有以参数形式给出的数字相加并返回和。`
}

func (*Sum) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) < 1 {
			return nil, errors.New("SUM 参数不足")
		}
		res := 0.00
		for key, param := range arguments {
			val, flag := param.(float64)
			if !flag {
				return nil, fmt.Errorf("CONCAT: 第 %d 个参数不是有效的 float64 类型", key+1)
			}
			res = val + res
		}
		return res, nil
	}
}
