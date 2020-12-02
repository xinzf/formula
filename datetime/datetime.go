package datetime

import (
	"github.com/Knetic/govaluate"
)

type DateTime struct {
	Base
}

func (this *DateTime) Init(cfg interface{}) error {
	panic("implement me")
}

func (this *DateTime) GetName() string {
	panic("implement me")
}

func (this *DateTime) GetFunc() govaluate.ExpressionFunction {
	panic("implement me")
}

func (this *DateTime) GetDescription() string {
	panic("implement me")
}
