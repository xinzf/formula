package formula

import (
	"github.com/Knetic/govaluate"
	"github.com/spf13/cast"
	"github.com/xinzf/formula/logic"
	"github.com/xinzf/formula/utils"
)

type Description interface {
	GetExpression() string
	GetArgDescribes() []map[string]string
	GetExample() string
}

type Function interface {
	Init(cfg interface{}) error
	GetName() string
	GetFunc() govaluate.ExpressionFunction
	GetCategory() string
	GetDescription() string
}

var _functions = make([]Function, 0)

func init() {
	// 高级函数
	//_ = RegisterFunction(new(advance.Json), nil)
	// 文本函数
	//_ = RegisterFunction(new(char.Concat), nil)
	//_ = RegisterFunction(new(char.Exact), nil)
	//_ = RegisterFunction(new(char.Left), nil)
	//_ = RegisterFunction(new(char.Len), nil)
	//_ = RegisterFunction(new(char.Lower), nil)
	//_ = RegisterFunction(new(char.Replace), nil)
	//_ = RegisterFunction(new(char.Rept), nil)
	//_ = RegisterFunction(new(char.Right), nil)
	//_ = RegisterFunction(new(char.Search), nil)
	//_ = RegisterFunction(new(char.Split), nil)
	//_ = RegisterFunction(new(char.Trim), nil)
	//_ = RegisterFunction(new(char.Upper), nil)
	// 日期函数
	//_ = RegisterFunction(new(datetime.Date), nil)
	//_ = RegisterFunction(new(datetime.DateAdd), nil)
	//_ = RegisterFunction(new(datetime.DateTime), nil)
	//_ = RegisterFunction(new(datetime.DayOfMonth), nil)
	//_ = RegisterFunction(new(datetime.DayOfYear), nil)
	//_ = RegisterFunction(new(datetime.Hour), nil)
	//_ = RegisterFunction(new(datetime.IsLeapYear), nil)
	//_ = RegisterFunction(new(datetime.Minute), nil)
	//_ = RegisterFunction(new(datetime.Month), nil)
	//_ = RegisterFunction(new(datetime.MonthBegin), nil)
	//_ = RegisterFunction(new(datetime.MonthEnd), nil)
	//_ = RegisterFunction(new(datetime.NextDate), nil)
	//_ = RegisterFunction(new(datetime.PrevDate), nil)
	//_ = RegisterFunction(new(datetime.Quarter), nil)
	//_ = RegisterFunction(new(datetime.QuarterBegin), nil)
	//_ = RegisterFunction(new(datetime.QuarterEnd), nil)
	//_ = RegisterFunction(new(datetime.Second), nil)
	//_ = RegisterFunction(new(datetime.Time), nil)
	//_ = RegisterFunction(new(datetime.Timestamp), nil)
	//_ = RegisterFunction(new(datetime.WeekDay), nil)
	//_ = RegisterFunction(new(datetime.WeekOfYear), nil)
	//_ = RegisterFunction(new(datetime.Year), nil)
	// 逻辑函数
	_ = RegisterFunction(new(logic.And), nil)
	//_ = RegisterFunction(new(logic.Equal), nil)
	_ = RegisterFunction(new(logic.False), nil)
	_ = RegisterFunction(new(logic.If), nil)
	_ = RegisterFunction(new(logic.Ifs), nil)
	_ = RegisterFunction(new(logic.IsEmpty), nil)
	//_ = RegisterFunction(new(logic.Match), nil)
	_ = RegisterFunction(new(logic.Not), nil)
	_ = RegisterFunction(new(logic.Or), nil)
	_ = RegisterFunction(new(logic.True), nil)
	//_ = RegisterFunction(new(logic.Xor), nil)
	// 数学函数
	//_ = RegisterFunction(new(math.Abs), nil)
	//_ = RegisterFunction(new(math.Average), nil)
	//_ = RegisterFunction(new(math.Count), nil)
	//_ = RegisterFunction(new(math.Float), nil)
	//_ = RegisterFunction(new(math.Int), nil)
	//_ = RegisterFunction(new(math.Max), nil)
	//_ = RegisterFunction(new(math.Min), nil)
	//_ = RegisterFunction(new(math.Mod), nil)
	//_ = RegisterFunction(new(math.Power), nil)
	//_ = RegisterFunction(new(math.Product), nil)
	//_ = RegisterFunction(new(math.Rand), nil)
	//_ = RegisterFunction(new(math.Round), nil)
	//_ = RegisterFunction(new(math.Sum), nil)
	//_ = RegisterFunction(new(math.Sumproduct), nil)
}

func RegisterFunction(fn Function, cfg interface{}) error {
	if err := fn.Init(cfg); err != nil {
		return err
	}
	_functions = append(_functions, fn)
	return nil
}

type Out struct {
	Name string              `json:"name"`
	Funs []map[string]string `json:"funs"`
}

func NewCalculator(values map[string]interface{}) *Calculator {
	c := &Calculator{
		values: values,
	}
	return c
}

type Calculator struct {
	values     map[string]interface{}
	expression string
}

func (this *Calculator) build() map[string]govaluate.ExpressionFunction {
	funcs := make(map[string]govaluate.ExpressionFunction)
	for _, f := range _functions {
		funcs[f.GetName()] = f.GetFunc()
	}
	return funcs
}

func (this *Calculator) Functions() []Out {
	outs := make([]Out, 0)
	indexs := make(map[string]int, 0)
	for _, f := range _functions {
		index, flag := indexs[f.GetCategory()]
		if !flag {
			outs = append(outs, Out{
				Name: f.GetCategory(),
				Funs: []map[string]string{},
			})
			indexs[f.GetCategory()] = len(outs) - 1
			index = len(outs) - 1
		}
		outs[index].Funs = append(outs[index].Funs, map[string]string{
			"name":        f.GetName(),
			"description": f.GetDescription(),
		})
	}

	return outs
}

//func (this *Calculator) EvalBool(expression string) (bool, error) {
//	this.expression = expression
//	val, err := this.Eval(expression)
//	if err != nil {
//		return false, err
//	}
//
//	v, flag := val.(bool)
//	if !flag {
//		return false, errors.New("返回值不是有效 Boolean 数据")
//	}
//	return v, nil
//}

func (this *Calculator) eval(expression string) (val interface{}, err error) {
	this.expression = expression
	defer func() {
		if er := recover(); er != nil {
			switch er.(type) {
			case error:
				err = er.(error)
			}
		}
	}()

	funcs := this.build()

	var exp *govaluate.EvaluableExpression
	exp, err = govaluate.NewEvaluableExpressionWithFunctions(expression, funcs)
	if err != nil {
		return nil, err
	}

	var data interface{}
	data, err = exp.Evaluate(this.values)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (this *Calculator) GetExpression() string {
	return this.expression
}

type Expression string

func (e Expression) String() string {
	return string(e)
}

func (e Expression) Eval(values map[string]interface{}) (interface{}, error) {
	return NewCalculator(values).eval(e.String())
}

func (e Expression) EvalBool(values map[string]interface{}) (bool, error) {
	val, err := e.Eval(values)
	if err != nil {
		return false, err
	}

	return cast.ToBoolE(val)
}

func (e Expression) EvalString(values map[string]interface{}) (string, error) {
	val, err := e.Eval(values)
	if err != nil {
		return "", err
	}

	return cast.ToStringE(val)
}

func (e Expression) EvalInt(values map[string]interface{}) (int64, error) {
	val, err := e.Eval(values)
	if err != nil {
		return 0, err
	}

	return cast.ToInt64E(val)
}

func (e Expression) EvalFloat(values map[string]interface{}) (float64, error) {
	val, err := e.Eval(values)
	if err != nil {
		return 0, err
	}
	return cast.ToFloat64E(val)
}

func (e Expression) EvalSliceString(values map[string]interface{}) ([]string, error) {
	val, err := e.Eval(values)
	if err != nil {
		return []string{}, err
	}

	return cast.ToStringSliceE(val)
}

func (e Expression) EvalSliceInt(values map[string]interface{}) ([]int, error) {
	val, err := e.Eval(values)
	if err != nil {
		return []int{}, err
	}

	return cast.ToIntSliceE(val)
}

func (e Expression) EvalBind(values map[string]interface{}, target interface{}) error {
	val, err := e.Eval(values)
	if err != nil {
		return err
	}

	return utils.NewConvert(val).Bind(target)
}
