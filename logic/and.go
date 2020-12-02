package logic

import (
	"errors"
	"github.com/Knetic/govaluate"
)

type And struct {
	Base
}

func (this *And) Init(cfg interface{}) error {
	return nil
}

func (this *And) GetName() string {
	return "AND"
}

func (*And) GetDescription() string {
	return `如果所有参数都为真，AND函数返回布尔值true，否则返回布尔值 false
用法：AND(逻辑表达式1,逻辑表达式2,...)
示例：AND(语文成绩>90,数学成绩>90,英语成绩>90)，如果三门课成绩都> 90，返回true，否则返回false`
}

func (this *And) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) == 0 {
			return nil, errors.New("AND: 参数数量不足")
		}

		data := true
		for _, a := range arguments {
			val, flag := a.(bool)
			if !flag {
				return nil, errors.New("AND: 不是有效的 Boolean 数据")
			}
			if val == false {
				data = false
				break
			}
		}

		return data, nil
	}
}
