package datetime

import (
	"github.com/Knetic/govaluate"
)

type DateAdd struct {
	Base
}

func (this *DateAdd) Init(cfg interface{}) error {
	return nil
}

func (this *DateAdd) GetName() string {
	return "DATEADD"
}

func (this *DateAdd) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		return nil, nil
	}
}

func (this *DateAdd) GetDescription() string {
	return `DATEADD函数用于计算指定日期之后或之前的日期
用法：DATEADD(datepart,number,date)
参数：
   datepart: （必选）可以是year、month、week、day
   number: （必选）间隔，一个有符号整数，支持负号（表示时间往前算起）
   date: （必选）从该日期开始算起，格式为yyyy-MM-dd HH:mm:ss或yyyy-MM-dd
`
	panic("implement me")
}
