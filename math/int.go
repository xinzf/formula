package math

import (
	"errors"
	"github.com/Knetic/govaluate"
	"github.com/xinzf/formula/utils"
)

type Int struct {
	Base
}

func (this *Int) Init(cfg interface{}) error {
	return nil
}

func (this *Int) GetName() string {
	return "INT"
}

func (*Int) GetDescription() string {
	return `INT(number)
将数据转换为整形。`
}

func (*Int) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) == 0 {
			return nil, errors.New("INT: 参数不足")
		}

		return utils.NewConvert(arguments[0]).Int(), nil
	}
}
