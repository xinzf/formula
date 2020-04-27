package math

import (
    "github.com/xinzf/formula/utils"
	"errors"
	"github.com/Knetic/govaluate"
	"math"
)

type Power struct {
	values map[string]interface{}
}

func (this *Power) GetName() string {
	return "POWER"
}

func (*Power) GetCategory() string {
	return "数学函数"
}

func (*Power) GetDescription() string {
	return `POWER(number, power)
返回数字乘幂的结果。
number: 必需。 基数。 可为任意实数。
power: 必需。 基数乘幂运算的指数。`
}

func (this *Power) SetValues(values map[string]interface{}) {
	this.values = values
}

func (*Power) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) != 2 {
			return nil, errors.New("POWER: 参数数量不对")
		}

		x := utils.NewConvert(arguments[0]).Float64()
		y := utils.NewConvert(arguments[1]).Float64()

		return math.Pow(x, y), nil

		//number, n_flag := arguments[0].(float64)
		//num_digit, d_flag := arguments[1].(float64)
		//if !n_flag {
		//	return nil, fmt.Errorf("ROUND: 第一个参数不是有效的 float 类型")
		//}
		//if !d_flag {
		//	return nil, fmt.Errorf("ROUND: 第二个参数不是有效的 float 类型")
		//}
		//res := math.Pow(number, num_digit)

		//return res, nil
	}
}
