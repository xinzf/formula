package datetime

import (
	"github.com/Knetic/govaluate"
)

type MonthEnd struct {
	Base
}

func (this *MonthEnd) Init(cfg interface{}) error {
	panic("implement me")
}

func (this *MonthEnd) GetName() string {
	panic("implement me")
}

func (this *MonthEnd) GetFunc() govaluate.ExpressionFunction {
	panic("implement me")
}

func (this *MonthEnd) GetDescription() string {
	panic("implement me")
}
