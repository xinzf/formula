package char

import "github.com/Knetic/govaluate"

type Search struct {
	Base
}

func (this *Search) Init(cfg interface{}) error {
	return nil
}

func (this *Search) GetName() string {
	return "SEARCH"
}

func (this *Search) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		return nil, nil
	}
}

func (this *Search) GetDescription() string {
	return `SEARCH函数可以获取文本1在文本2中的开始位置
用法：SEARCH(文本1,文本2)
示例：SEARCH("2016","居润2016")返回4`
}
