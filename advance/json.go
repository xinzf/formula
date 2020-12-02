package advance

import (
	"errors"
	"github.com/Knetic/govaluate"
	"github.com/buger/jsonparser"
	"github.com/cstockton/go-conv"
	"github.com/json-iterator/go"
)

type Json struct {
	values map[string]interface{}
}

func (this *Json) Init(cfg interface{}) error {
	return nil
}

func (this *Json) GetName() string {
	return "JSON"
}

func (this *Json) SetValues(values map[string]interface{}) {
	this.values = values
}

func (this *Json) GetCategory() string {
	return "高级函数"
}

func (this *Json) GetDescription() string {
	return `JSON函数可以将Map或者List中的数据提取出来以便于后期使用和计算
用法：JSON(map/list,路径1,路径2,...)
示例：JSON({"person":{"name":"xxx"}},person","name")可以取出xxx
示例：JSON([{"age":33}],"[0]","age"),可以取出33`
}

func (this *Json) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) < 2 {
			return nil, errors.New("JSON: 参数数量不足")
		}

		target := arguments[0]
		paths := make([]string, 0)

		for _, a := range arguments[1:] {
			str, ok := a.(string)
			if !ok {
				return nil, errors.New("JSON: 路径参数必须是字符串")
			}

			paths = append(paths, str)
		}

		data, err := jsoniter.Marshal(target)
		if err != nil {
			return nil, errors.New("JSON: 解析第一个参数为JSON失败,Err: " + err.Error())
		}

		val, valType, _, err := jsonparser.Get(data, paths...)
		if err != nil {
			return nil, errors.New("JSON: 提取数据失败,Err: " + err.Error())
		}

		switch valType {
		case jsonparser.String:
			return string(val), nil
		case jsonparser.Number:
			number, err := conv.Float64(string(val))
			if err != nil {
				return nil, errors.New("JSON: 提取数据失败,Err: " + err.Error())
			}
			return number, nil
		case jsonparser.Object:
			mp := map[string]interface{}{}
			err = jsoniter.Unmarshal(val, &mp)
			if err != nil {
				return nil, errors.New("JSON: 提取数据失败,Err: " + err.Error())
			}
			return mp, nil
		case jsonparser.Array:
			list := make([]interface{}, 0)
			err = jsoniter.Unmarshal(val, &list)
			if err != nil {
				return nil, errors.New("JSON: 提取数据失败,Err: " + err.Error())
			}
			return list, nil
		case jsonparser.Boolean:
			v, err := conv.Bool(string(val))
			if err != nil {
				return nil, errors.New("JSON: 提取数据失败,Err: " + err.Error())
			}
			return v, nil
		case jsonparser.Null:
			return nil, nil
		}

		return nil, nil
	}
}
