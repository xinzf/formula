package math

import (
	"errors"
	"github.com/Knetic/govaluate"
	"github.com/spf13/cast"
	"math"
)

type Mod struct {
	Base
}

func (this *Mod) Init(cfg interface{}) error {
	return nil
}

func (this *Mod) GetName() string {
	return "MOD"
}

func (*Mod) GetDescription() string {
	return `返回两数相除的余数。 结果的符号与被除数相同。
number: 必需。 要计算余数的被除数。
divisor: 必需。 除数。`
}

func (*Mod) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) != 2 {
			return nil, errors.New("MOD: 参数数量不对")
		}

		x, err := cast.ToFloat64E(arguments[0])
		if err != nil {
			return nil, errors.New("MOD 第一个参数类型不匹配")
		}
		y, err := cast.ToFloat64E(arguments[1])
		if err != nil {
			return nil, errors.New("MOD 第二个参数类型不匹配")
		}

		if y == 0 {
			return nil, errors.New("MOD: 第 2 个参数不能为 0 ，或者其类型错误")
		}

		return math.Mod(x, y), nil

		//number, n_flag := arguments[0].(float64)
		//num_digit, d_flag := arguments[1].(float64)
		//if !n_flag {
		//	return nil, fmt.Errorf("ROUND: 第一个参数不是有效的 float 类型")
		//}
		//if !d_flag {
		//	return nil, fmt.Errorf("ROUND: 第二个参数不是有效的 float 类型")
		//}
		//res := math.Mod(number, num_digit)
		//return res, nil
	}
}
