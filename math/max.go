package math

import (
	"github.com/Knetic/govaluate"
	"sort"

	"fmt"
)

type Max struct {
	Base
}

func (this *Max) Init(cfg interface{}) error {
	return nil
}

func (this *Max) GetName() string {
	return "MAX"
}

func (*Max) GetDescription() string {
	return `MAX(number1, [number2], …)
返回一组值中的最大值。`
}

func (*Max) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		var paramSlice []float64
		for key, param := range arguments {
			val, flag := param.(float64)
			if !flag {
				return nil, fmt.Errorf("MAX: 第 %d 个参数不是有效数字格式", key)
			}
			paramSlice = append(paramSlice, val)
			//paramSlice = append(paramSlice, param.(float64))
		}
		sort.Float64s(paramSlice)
		max := paramSlice[len(arguments)-1]
		return max, nil
	}
}
