package datetime

import (
	"github.com/Knetic/govaluate"
)

type DayOfYear struct {
	Base
}

func (this *DayOfYear) Init(cfg interface{}) error {
	panic("implement me")
}

func (this *DayOfYear) GetName() string {
	panic("implement me")
}

func (this *DayOfYear) GetFunc() govaluate.ExpressionFunction {
	panic("implement me")
}

func (this *DayOfYear) GetDescription() string {
	panic("implement me")
}
