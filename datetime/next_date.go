package datetime

import (
	"github.com/Knetic/govaluate"
)

type NextDate struct {
	Base
}

func (this *NextDate) Init(cfg interface{}) error {
	panic("implement me")
}

func (this *NextDate) GetName() string {
	panic("implement me")
}

func (this *NextDate) GetFunc() govaluate.ExpressionFunction {
	panic("implement me")
}

func (this *NextDate) GetDescription() string {
	panic("implement me")
}
