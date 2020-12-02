package math

import (
	"errors"
	"github.com/Knetic/govaluate"
	"math/rand"
)

type Rand struct {
	Base
}

func (this *Rand) Init(cfg interface{}) error {
	return nil
}

func (this *Rand) GetName() string {
	return "RAND"
}

func (*Rand) GetDescription() string {
	return `RAND()
返回大于等于 0 且小于 1 的均匀分布随机实数。每一次触发计算都会变化。`
}

func (*Rand) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) > 0 {
			return nil, errors.New("RAND: 参数数量不对")
		}
		return rand.Float64(), nil
	}
}
