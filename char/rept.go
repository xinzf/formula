package char

import (
	"errors"
	"github.com/Knetic/govaluate"
	"github.com/spf13/cast"
)

type Rept struct {
	Base
}

func (this *Rept) Init(cfg interface{}) error {
	return nil
}

func (this *Rept) GetName() string {
	return "REPT"
}

func (this *Rept) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) != 2 {
			return nil, errors.New("REPT 参数'数量不足")
		}
		argument1, err := cast.ToStringE(arguments[0])
		if err != nil {
			return nil, err
		}
		argument2, err := cast.ToIntE(arguments[1])
		if err != nil {
			return nil, err
		}
		var str string

		for i := 0; i < argument2; i++ {
			str += argument1
		}
		
		return str, nil
	}
}

func (this *Rept) GetDescription() string {
	return `REPT函数可以将文本重复一定次数
用法：REPT(文本,重复次数)
示例：REPT("居润",3)返回"居润居润居润"`
}
