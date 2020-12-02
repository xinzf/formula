package logic

import (
	"errors"
	"fmt"
	"github.com/Knetic/govaluate"
	"reflect"
)

type Match struct {
	Base
}

func (this *Match) Init(cfg interface{}) error {
	return nil
}

func (this *Match) GetName() string {
	return "MATCH"
}

func (*Match) GetDescription() string {
	return `用于匹配参数是数组的时候，是否符合被匹配数据相等
用法：MATCH(value,"要匹配的值","...")
示例：MATCH(李晨,"李**","陈**")`
}

func (this *Match) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {

		//return "向志", nil

		if len(arguments) < 2 {
			return nil, errors.New("MATCH: 参数数量不足")
		}

		targets := make([]interface{}, 0)

		for _, target := range arguments[1:] {
			if reflect.ValueOf(target).Kind() == reflect.Slice {
				vv := target.([]interface{})
				targets = append(targets, vv...)
			} else {
				targets = append(targets, target)
			}
		}

		for _, v := range targets {
			if reflect.DeepEqual(arguments[0], v) {
				fmt.Println(111)
				return true, nil
			}
		}
		return false, nil
	}
}
