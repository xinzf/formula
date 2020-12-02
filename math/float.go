package math

import (
	"errors"
	"github.com/Knetic/govaluate"
	"github.com/xinzf/formula/utils"
)

type Float struct {
	Base
}

func (this *Float) Init(cfg interface{}) error {
	return nil
}

func (this *Float) GetName() string {
	return "FLOAT"
}

func (*Float) GetDescription() string {
	return `FLOAT(number)
将数据转换为浮点型。`
}

func (*Float) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) == 0 {
			return nil, errors.New("FLOAT: 参数不足")
		}

		val := utils.NewConvert(arguments[0])
		return val.Float64(), nil
	}
}
