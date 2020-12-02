package datetime

import (
	"github.com/Knetic/govaluate"
)

type IsLeapYear struct {
	Base
}

func (this *IsLeapYear) Init(cfg interface{}) error {
	panic("implement me")
}

func (this *IsLeapYear) GetName() string {
	panic("implement me")
}

func (this *IsLeapYear) GetFunc() govaluate.ExpressionFunction {
	panic("implement me")
}

func (this *IsLeapYear) GetDescription() string {
	panic("implement me")
}
