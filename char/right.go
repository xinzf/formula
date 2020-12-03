package char

import (
	"errors"
	"github.com/Knetic/govaluate"
	"github.com/spf13/cast"
)

type Right struct {
	Base
}

func (this *Right) Init(cfg interface{}) error {
	return nil
}

func (this *Right) GetName() string {
	return "RIGHT"
}

func (this *Right) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) != 2 {
			return nil, errors.New("RIGHT 参数数量不足")
		}
		argument1, err := cast.ToStringE(arguments[0])
		if err != nil {
			return nil, errors.New("RIGHT 第一个参数类型不匹配")
		}
		argument2, err := cast.ToIntE(arguments[1])

		if err != nil {
			return nil, errors.New("RIGHT 第二个参数不匹配")
		}

		if argument2 <= 0 {
			return argument1, nil
		}

		rune_ := []rune(argument1)

		if len(rune_)-1 < argument2 {
			return argument1, nil
		}

		return string(rune_[len(rune_) - argument2:]), nil
	}
}

func (this *Right) GetDescription() string {
	return `RIGHT函数可以获取由给定文本右端指定数量的字符构成的文本值
用法：RIGHT(文本,文本长度)
示例：RIGHT("三年二班周杰伦",3)返回"周杰伦"，也就是"三年二班周杰伦"从右往左的前3个字符`
}
