package char

import (
	"errors"
	"github.com/Knetic/govaluate"
	"github.com/spf13/cast"
	"strings"
)

type Split struct {
	Base
}

func (this *Split) Init(cfg interface{}) error {
	return nil
}

func (this *Split) GetName() string {
	return "SPLIT"
}

func (this *Split) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) != 2 {
			return nil, errors.New("SPLIT 参数数量不足")
		}
		argument1, err := cast.ToStringE(arguments[0])
		if err != nil {
			return nil, errors.New("SPLIT 第一个参数类型不匹配")
		}
		argument2, err := cast.ToStringE(arguments[1])
		if err != nil {
			return nil, errors.New("SPLIT 第二个参数类型不匹配")
		}

		return strings.Join(strings.Split(argument1, argument2), ","), nil
	}
}

func (this *Split) GetDescription() string {
	return `SPLIT函数可以将文本按指定分割符分割成数组
用法：SPLIT(文本,分隔符_文本)
示例：SPLIT("居润-应用定制工具","-")返回"居润，应用定制工具"`
}
