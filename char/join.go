package char

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/spf13/cast"
	"strings"
)

type Join struct {
	Base
}

func (j *Join) Init(cfg interface{}) error {
	return nil
}

func (j *Join) GetName() string {
	return "JOIN"
}

func (j *Join) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		fmt.Println("join", len(arguments), arguments)
		if len(arguments) < 2 {
			return nil, errors.New("JOIN 参数数量不足")
		}
		//if _, ok := arguments[0].([]string); !ok {
		//	return nil, errors.New("JOIN 第一个参数类型不正确")
		//}
		if len(arguments) == 2 {
			if _, ok := arguments[1].(string); !ok {
				return nil, errors.New("JOIN 第二个参数类型不正确")
			}
			return strings.Join(cast.ToStringSlice(arguments[0]), cast.ToString(arguments[1])), nil
		}
		sep := arguments[len(arguments)-1]
		arr := arguments[:len(arguments)-1]
		return strings.Join(cast.ToStringSlice(arr), cast.ToString(sep)), nil
	}
}

func (j *Join) GetDescription() string {
	return `JOIN函数用来把数组转化字符串`
}
