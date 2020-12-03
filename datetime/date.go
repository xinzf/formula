package datetime

import (
	"github.com/Knetic/govaluate"
	"github.com/uniplaces/carbon"
)

type Date struct {
	Base
}

func (this *Date) Init(cfg interface{}) error {
	return nil
}

func (this *Date) GetName() string {
	return "DATE"
}

func (this *Date) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		return carbon.Now().Format("2006-01-02"), nil
	}
}

func (this *Date) GetDescription() string {
	return `DATE函数获取当前日期，格式为 yyyy-MM-dd
用法：DATE()`
}
