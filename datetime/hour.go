package datetime

import (
	"github.com/Knetic/govaluate"
)

type Hour struct {
	Base
}

func (this *Hour) Init(cfg interface{}) error {
	panic("implement me")
}

func (this *Hour) GetName() string {
	panic("implement me")
}

func (this *Hour) GetFunc() govaluate.ExpressionFunction {
	panic("implement me")
}

func (this *Hour) GetDescription() string {
	panic("implement me")
}
