package datetime

import (
	"github.com/Knetic/govaluate"
)

type Month struct {
	Base
}

func (this *Month) Init(cfg interface{}) error {
	panic("implement me")
}

func (this *Month) GetName() string {
	panic("implement me")
}

func (this *Month) GetFunc() govaluate.ExpressionFunction {
	panic("implement me")
}

func (this *Month) GetDescription() string {
	panic("implement me")
}
