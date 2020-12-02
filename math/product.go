package math

import (
	"github.com/Knetic/govaluate"
	"github.com/xinzf/formula/utils"
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
		for _, param := range arguments {
			val := utils.NewConvert(param).Float64()
			//val, flag := param.(float64)
			//if !flag {
			//	return nil, fmt.Errorf("CONCAT: 第 %d 个参数不是有效的 float64 类型", key)
			//}
			res = val * res
		}
		return res, nil
	}
}
