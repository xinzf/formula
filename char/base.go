package char

type Base struct {
}

func (Base) GetCategory() string {
	return "文本函数"
}

type Description struct {
	expression string
	args       []map[string]string
	example    string
}

func (d Description) GetExpression() string {
	return d.expression
}

func (d Description) GetArgDescribes() []map[string]string {
	return d.args
}

func (d Description) GetExample() string {
	return d.example
}
