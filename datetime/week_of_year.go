package datetime

import (
	"github.com/Knetic/govaluate"
)

type WeekOfYear struct {
	Base
}

func (this *WeekOfYear) Init(cfg interface{}) error {
	panic("implement me")
}

func (this *WeekOfYear) GetName() string {
	panic("implement me")
}

func (this *WeekOfYear) GetFunc() govaluate.ExpressionFunction {
	panic("implement me")
}

func (this *WeekOfYear) GetDescription() string {
	panic("implement me")
}
