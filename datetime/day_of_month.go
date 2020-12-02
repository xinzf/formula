package datetime

import (
	"github.com/Knetic/govaluate"
)

type DayOfMonth struct {
	Base
}

func (this *DayOfMonth) Init(cfg interface{}) error {
	panic("implement me")
}

func (this *DayOfMonth) GetName() string {
	panic("implement me")
}

func (this *DayOfMonth) GetFunc() govaluate.ExpressionFunction {
	panic("implement me")
}

func (this *DayOfMonth) GetDescription() string {
	panic("implement me")
}
