package logic

import (
	"errors"
	"github.com/Knetic/govaluate"
	"reflect"
)

type Equal struct {
	Base
}

func (this *Equal) Init(cfg interface{}) error {
	return nil
}

func (this *Equal) GetName() string {
	return "EQUAL"
}

func (this *Equal) GetFunc() govaluate.ExpressionFunction {
	return func(args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return false, errors.New("EQUAL: 参数必须是两个")
		}

		return reflect.DeepEqual(args[0], args[1]), nil
	}
}

func (this *Equal) GetDescription() string {
	return `EQUAL函数是数据比较函数，若完全相当返回True，否则返回False，匹配时区分类型和大小写
用法：EQUAL(val1,val2) 或者 EQUAL(val1,[val1,val2])`
}
