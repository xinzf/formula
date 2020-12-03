package char

import (
	"errors"
	"github.com/Knetic/govaluate"
	"github.com/spf13/cast"
	"strings"
)

type Upper struct {
	Base
}

func (this *Upper) Init(cfg interface{}) error {
	return nil
}

func (this *Upper) GetName() string {
	return "UPPER"
}

func (this *Upper) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) != 1 {
			return nil, errors.New("UPPER 参数不足")
		}

		str, err := cast.ToStringE(arguments[0])
		if err != nil {
			return nil, errors.New("UPPER 传入类型有误")
		}

		return strings.ToTitle(str), nil
	}
}

func (this *Upper) GetDescription() string {
	return `UPPER函数可以将一个文本中的所有小写字母转换为大写字母
用法：UPPER(文本)
示例：UPPER("jayz")返回"JAYZ"`
}
