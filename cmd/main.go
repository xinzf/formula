package main

import (
	"fmt"
	"github.com/xinzf/formula"
)

func main() {
	args := make(map[string]int)
	//args["sfd"]=3
	//args=append(args,true)
	exp := `ISEMPTY(args)`
	ret, err := formula.NewCalculator(map[string]interface{}{
		"args": args,
	}).Eval(exp)
	if err != nil {
		panic(err)
	}

	fmt.Println(ret)
}
