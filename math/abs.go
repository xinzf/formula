package math

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"math"
	//"sort"
)

type Abs struct {
	Base
}

func (this *Abs) Init(cfg interface{}) error {
	return nil
}

func (this *Abs) GetName() string {
	return "ABS"
}

func (*Abs) GetDescription() string {
	return `ABS(number)
返回数字的绝对值。`
}

func (*Abs) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		val, flag := arguments[0].(float64)
		if !flag {
			return nil, fmt.Errorf("ABS: 参数不是有效的 float 类型")
		}
		return math.Abs(val), nil
	}
}
