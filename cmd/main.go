package main

import (
	"fmt"
	"github.com/xinzf/formula"
)

func main() {
	args := make(map[string]interface{})
	args["args"] = ""
	//args["sfd"]=3
	//args=append(args,true)
	var exp formula.Expression = `ISEMPTY(args)`
	ret, err := exp.EvalBool(args)
	if err != nil {
		panic(err)
	}

	fmt.Println(ret)
}
