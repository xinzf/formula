package logic

import (
	"errors"
	"github.com/Knetic/govaluate"
	"github.com/spf13/cast"
)

type IsEmpty struct {
	Base
}

func (this *IsEmpty) Init(cfg interface{}) error {
	return nil
}

func (this *IsEmpty) GetName() string {
	return "ISEMPTY"
}

func (this *IsEmpty) GetFunc() govaluate.ExpressionFunction {
	return func(arg ...interface{}) (interface{}, error) {

		if len(arg) != 1 {
			return false, errors.New("ISEMPTY: 参数只能是一个")
		}

		str, err := cast.ToStringE(arg[0])
		if err == nil && str == "" {
			return true, nil
		}

		slice, err := cast.ToSliceE(arg[0])
		if err == nil && len(slice) == 0 {
			return true, nil
		}
		strs, err := cast.ToStringSliceE(arg[0])
		if err == nil && len(strs) == 0 {
			return true, nil
		}

		ints, err := cast.ToIntSliceE(arg[0])
		if err == nil && len(ints) == 0 {
			return true, nil
		}

		booleans, err := cast.ToBoolSliceE(arg[0])
		if err == nil && len(booleans) == 0 {
			return true, nil
		}

		durs, err := cast.ToDurationSliceE(arg[0])
		if err == nil && len(durs) == 0 {
			return true, nil
		}

		if mp, err := cast.ToStringMapE(arg[0]); err == nil && len(mp) == 0 {
			return true, nil
		}

		if mp, err := cast.ToStringMapBoolE(arg[0]); err == nil && len(mp) == 0 {
			return true, nil
		}

		if mp, err := cast.ToStringMapStringSliceE(arg[0]); err == nil && len(mp) == 0 {
			return true, nil
		}

		if mp, err := cast.ToStringMapIntE(arg[0]); err == nil && len(mp) == 0 {
			return true, nil
		}

		if mp, err := cast.ToStringMapInt64E(arg[0]); err == nil && len(mp) == 0 {
			return true, nil
		}

		if mp, err := cast.ToStringMapStringE(arg[0]); err == nil && len(mp) == 0 {
			return true, nil
		}

		return false, nil
	}
}

func (this *IsEmpty) GetDescription() string {
	return `ISEMPTY函数可以用来判断给定数据是否为空文本、空数组、空对象，如果为空返回TRUE，否则返回FALSE
用法：ISEMPTY(数据)`
}
