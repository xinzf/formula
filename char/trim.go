package char

import "github.com/Knetic/govaluate"

type Trim struct {
	Base
}

func (this *Trim) Init(cfg interface{}) error {
	return nil
}

func (this *Trim) GetName() string {
	return "TRIM"
}

func (this *Trim) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		return nil, nil
	}
}

func (this *Trim) GetDescription() string {
	return `TRIM函数可以删除文本首尾的空格
用法：TRIM(文本)
示例：TRIM(" 居润 ")返回"居润"`
}
