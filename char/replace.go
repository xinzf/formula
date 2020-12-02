package char

import "github.com/Knetic/govaluate"

type Replace struct {
	Base
}

func (this *Replace) Init(cfg interface{}) error {
	return nil
}

func (this *Replace) GetName() string {
	return "REPLACE"
}

func (this *Replace) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		return nil, nil
	}
}

func (this *Replace) GetDescription() string {
	return `REPLACE函数可以根据指定的字符数，将部分文本替换为不同的文本
用法：REPLACE(文本,开始位置,替换长度,新文本)
示例：REPLACE("居润应用定制工具",3,6,"企业PAAS管理平台")返回"居润企业PAAS管理平台"`
}
