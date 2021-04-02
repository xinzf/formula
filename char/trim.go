package char

import (
	"errors"
	"github.com/Knetic/govaluate"
	"github.com/spf13/cast"
	"strings"
)

type Trim struct {
	Base
}

func (this *Trim) Init(cfg interface{}) error {
	return nil
}

func (this *Trim) GetName() string {
	return "TRIM"
}

func (this *Trim) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) != 1 {
			return nil, errors.New("TRIM 参数不足")
		}
		str, err := cast.ToStringE(arguments[0])

		if err != nil {
			return nil, errors.New("TRIM 传入参数有误")
		}

		return strings.Replace(str, " ", "", -1), nil
	}
}

func (this *Trim) GetDescription() string {
	return `TRIM函数可以删除文本首尾的空格
用法：TRIM(文本)
示例：TRIM(" 居润 ")返回"居润"`
}
