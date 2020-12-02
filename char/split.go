package char

import "github.com/Knetic/govaluate"

type Split struct {
	Base
}

func (this *Split) Init(cfg interface{}) error {
	return nil
}

func (this *Split) GetName() string {
	return "SPLIT"
}

func (this *Split) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		return nil, nil
	}
}

func (this *Split) GetDescription() string {
	return `SPLIT函数可以将文本按指定分割符分割成数组
用法：SPLIT(文本,分隔符_文本)
示例：SPLIT("居润-应用定制工具","-")返回"居润，应用定制工具"`
}
