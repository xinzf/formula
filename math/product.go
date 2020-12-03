package math

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/spf13/cast"
)

type Product struct {
	Base
}

func (this *Product) Init(cfg interface{}) error {
	return nil
}

func (this *Product) GetName() string {
	return "PRODUCT"
}

func (*Product) GetDescription() string {
	return `PRODUCT(number1, [number2], …)
函数使所有以参数形式给出的数字相乘并返回乘积。`
}

func (*Product) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		res := 1.0
		if len(arguments) < 1 {
			return nil, errors.New("PRODUCT 参数数量不足")
		}

		for k, param := range arguments {
			val, err := cast.ToFloat64E(param)
			if err != nil {
				return nil, fmt.Errorf("PRODUCT: 第 %d 个参数不是有效的 float64 类型", k)
			}
			res = val * res
		}

		return res, nil
	}
}
