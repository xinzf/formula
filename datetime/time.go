package datetime

import (
	"github.com/Knetic/govaluate"
)

type Time struct {
	Base
}

func (this *Time) Init(cfg interface{}) error {
	panic("implement me")
}

func (this *Time) GetName() string {
	panic("implement me")
}

func (this *Time) GetFunc() govaluate.ExpressionFunction {
	panic("implement me")
}

func (this *Time) GetDescription() string {
	panic("implement me")
}
