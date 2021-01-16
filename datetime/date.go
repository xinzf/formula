package datetime

import (
	"errors"
	"github.com/Knetic/govaluate"
	"github.com/spf13/cast"
	"github.com/uniplaces/carbon"
	"github.com/xinzf/formula/utils"
	"time"
)

type Date struct {
	Base
}

func (this *Date) Init(cfg interface{}) error {
	return nil
}

func (this *Date) GetName() string {
	return "DATE"
}

func (this *Date) GetFunc() govaluate.ExpressionFunction {
	return func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) != 2 {
			return nil, errors.New("DATE 参数不足，需要2个参数")
		}

		t := time.Now()

		timestamp, err := cast.ToInt64E(arguments[0])
		if err == nil {
			str := cast.ToString(timestamp)
			if len(str) >= 13 {
				t = utils.UnixMilliToTime(timestamp)
			} else {
				t = time.Unix(timestamp, 0)
			}
		}

		format := cast.ToString(arguments[1])
		switch format {
		case "date":
			format = carbon.DateFormat
		case "datetime":
			format = carbon.DefaultFormat
		default:
			format = carbon.DefaultFormat
		}

		return carbon.NewCarbon(t).Format(format), nil
	}
}

func (this *Date) GetDescription() string {
	return `DATE函数格式化日期
语法：DATE(timestamp,format)
参数：
timestamp: 时间戳（整型），如果是不为时间戳的任意数据，取当前时间
format: 可选值：date、datetime，默认为 datetime
用法：DATE(1610776543084,'date')、DATE('NOW','datetime')`
}
