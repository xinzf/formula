package char

import (
	"errors"
	"github.com/Knetic/govaluate"
	"github.com/spf13/cast"
)

type Replace struct {
	Base
}

func (this *Replace) Init(cfg interface{}) error {
	return nil
}

func (this *Replace) GetName() string {
	return "REPLACE"
}

//GetFunc
/*
	1、验证参数个数
	2、验证参数类型
	3、验证参数是否合法
	4、startStr + 新文本 + endStr
*/
func (this *Replace) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {

		if len(arguments) != 4 {
			return nil, errors.New("REPLACE 参数数量不足")
		}

		argument1, err := cast.ToStringE(arguments[0])
		if err != nil {
			return nil, errors.New("REPLACE 第一个参数类型错误")
		}

		argument2, err := cast.ToIntE(arguments[1])
		if err != nil {
			return nil, errors.New("REPLACE 第二个参数类型错误")
		}

		argument3, err := cast.ToIntE(arguments[2])
		if err != nil {
			return nil, errors.New("REPLACE 第三个参数类型错误")
		}

		argument4, err := cast.ToStringE(arguments[3])
		if err != nil {
			return nil, errors.New("REPLACE 第四个参数类型错误")
		}

		if argument2 <= 0 {
			return argument1, nil
		}

		rune_ := []rune(argument1)

		if len(rune_)-1 < argument2 || len(rune_)-1 < argument3 {
			return argument1, nil
		}

		if argument2-1+argument3 > len(rune_) {
			return argument1, nil
		}

		startStr := string(rune_[:argument2-1])
		endStr := string(rune_[argument2-1+argument3:])

		return startStr + argument4 + endStr, nil
	}
}

func (this *Replace) GetDescription() string {
	return `REPLACE函数可以根据指定的字符数，将部分文本替换为不同的文本
用法：REPLACE(文本,开始位置,替换长度,新文本)
示例：REPLACE("居润应用定制工具",2,6,"企业PAAS管理平台")返回"居润企业PAAS管理平台"`
}
