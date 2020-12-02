package char

import "github.com/Knetic/govaluate"

type Right struct {
	Base
}

func (this *Right) Init(cfg interface{}) error {
	return nil
}

func (this *Right) GetName() string {
	return "RIGHT"
}

func (this *Right) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		return nil, nil
	}
}

func (this *Right) GetDescription() string {
	return `RIGHT函数可以获取由给定文本右端指定数量的字符构成的文本值
用法：RIGHT(文本,文本长度)
示例：RIGHT("三年二班周杰伦",3)返回"周杰伦"，也就是"三年二班周杰伦"从右往左的前3个字符`
}
