package math

import (
	"github.com/Knetic/govaluate"
	"sort"

	"fmt"
)

type Min struct {
	Base
}

func (this *Min) Init(cfg interface{}) error {
	return nil
}

func (this *Min) GetName() string {
	return "MIN"
}

func (*Min) GetDescription() string {
	return `Min(number1, [number2], …)
返回一组值中的最大值。`
}

func (*Min) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		//var paramSlice = []float64{}
		if len(arguments) < 1 {
			return nil, fmt.Errorf("MIN 参数个数不足")
		}
		paramSlice := make([]float64, 0)
		for key, param := range arguments {
			val, flag := param.(float64)
			if !flag {
				return nil, fmt.Errorf("MIN: 第 %d 个参数不是有效数字", key)
			}
			//val := utils.NewConvert(param).Float64()
			paramSlice = append(paramSlice, val)
			//paramSlice = append(paramSlice, param.(float64))
		}
		sort.Float64s(paramSlice)
		Min := paramSlice[0]
		return Min, nil
	}
}
