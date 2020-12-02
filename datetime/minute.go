package datetime

import (
	"github.com/Knetic/govaluate"
)

type Minute struct {
	Base
}

func (this *Minute) Init(cfg interface{}) error {
	panic("implement me")
}

func (this *Minute) GetName() string {
	panic("implement me")
}

func (this *Minute) GetFunc() govaluate.ExpressionFunction {
	panic("implement me")
}

func (this *Minute) GetDescription() string {
	panic("implement me")
}
