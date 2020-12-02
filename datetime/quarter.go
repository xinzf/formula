package datetime

import (
	"github.com/Knetic/govaluate"
)

type Quarter struct {
	Base
}

func (this *Quarter) Init(cfg interface{}) error {
	panic("implement me")
}

func (this *Quarter) GetName() string {
	panic("implement me")
}

func (this *Quarter) GetFunc() govaluate.ExpressionFunction {
	panic("implement me")
}

func (this *Quarter) GetDescription() string {
	panic("implement me")
}
