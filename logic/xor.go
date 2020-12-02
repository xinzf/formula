package logic

import (
	"github.com/Knetic/govaluate"
)

type Xor struct {
	Base
}

func (this *Xor) Init(cfg interface{}) error {
	return nil
}

func (this *Xor) GetName() string {
	return "XOR"
}

func (this *Xor) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		return nil, nil
	}
}

func (this *Xor) SetValues(values map[string]interface{}) {
}

func (this *Xor) GetDescription() string {
	return `返回所有参数的异或值。异或的含义是，两个值相同，返回false，两个值不同，返回true。
用法：XOR(logical1,[logical2], …)`
}
