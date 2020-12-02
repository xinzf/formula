package datetime

import (
	"github.com/Knetic/govaluate"
)

type Year struct {
	Base
}

func (this *Year) Init(cfg interface{}) error {
	panic("implement me")
}

func (this *Year) GetName() string {
	panic("implement me")
}

func (this *Year) GetFunc() govaluate.ExpressionFunction {
	panic("implement me")
}

func (this *Year) GetDescription() string {
	panic("implement me")
}
