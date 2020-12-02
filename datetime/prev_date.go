package datetime

import (
	"github.com/Knetic/govaluate"
)

type PrevDate struct {
	Base
}

func (this *PrevDate) Init(cfg interface{}) error {
	panic("implement me")
}

func (this *PrevDate) GetName() string {
	panic("implement me")
}

func (this *PrevDate) GetFunc() govaluate.ExpressionFunction {
	panic("implement me")
}

func (this *PrevDate) GetDescription() string {
	panic("implement me")
}
