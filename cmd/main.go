package main

import (
	"fmt"
	"github.com/xinzf/formula"
)

func main() {
	args := make(map[string]interface{})
	args["aaa"] = []float64{1,2}
	args["bbb"] = []float64{3,2}
	//args=append(args,true)
	var exp formula.Expression = `SUMPRODUCT(aaa,bbb)`
	ret, err := exp.EvalFloat(args)
	if err != nil {
		panic(err)
	}

	fmt.Println(ret)
}
