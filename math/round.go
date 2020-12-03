package math

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/spf13/cast"
	//"sort"
)

type Round struct {
	Base
}

func (this *Round) Init(cfg interface{}) error {
	return nil
}

func (this *Round) GetName() string {
	return "ROUND"
}

func (*Round) GetDescription() string {
	return `ROUND(number, num_digits)
将数字四舍五入到指定的位数。
number: 必需。 要四舍五入的数字。
num_digits: 必需。 要进行四舍五入运算的位数。`
}

func (*Round) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) != 2 {
			return nil, errors.New("ROUND 参数个数不足")
		}
		argument1, err := cast.ToFloat64E(arguments[0])
		if err != nil {
			return nil, errors.New("ROUND: 第一个参数不是有效的 float 类型")
		}

		argument2, err := cast.ToIntE(arguments[1])
		if err != nil {
			return nil, errors.New("ROUND: 第一个参数不是有效的 int 类型")
		}
		str := "%." + cast.ToString(argument2) + "f"

		return cast.ToFloat64(fmt.Sprintf(str, argument1)), nil
	}
}
