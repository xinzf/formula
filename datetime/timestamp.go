package datetime

import (
	"github.com/Knetic/govaluate"
)

type Timestamp struct {
	Base
}

func (this *Timestamp) Init(cfg interface{}) error {
	panic("implement me")
}

func (this *Timestamp) GetName() string {
	panic("implement me")
}

func (this *Timestamp) GetFunc() govaluate.ExpressionFunction {
	panic("implement me")
}

func (this *Timestamp) GetDescription() string {
	panic("implement me")
}
