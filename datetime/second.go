package datetime

import (
	"github.com/Knetic/govaluate"
)

type Second struct {
	Base
}

func (this *Second) Init(cfg interface{}) error {
	panic("implement me")
}

func (this *Second) GetName() string {
	panic("implement me")
}

func (this *Second) GetFunc() govaluate.ExpressionFunction {
	panic("implement me")
}

func (this *Second) GetDescription() string {
	panic("implement me")
}
