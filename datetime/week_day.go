package datetime

import (
	"github.com/Knetic/govaluate"
)

type WeekDay struct {
	Base
}

func (this *WeekDay) Init(cfg interface{}) error {
	panic("implement me")
}

func (this *WeekDay) GetName() string {
	panic("implement me")
}

func (this *WeekDay) GetFunc() govaluate.ExpressionFunction {
	panic("implement me")
}

func (this *WeekDay) GetDescription() string {
	panic("implement me")
}
