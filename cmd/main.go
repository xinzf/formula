package main

import (
	"fmt"
	"github.com/xinzf/formula"
)

func main() {
	var exp formula.Expression = `DATE('fsdf','')`
	ret, err := exp.EvalString(nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(ret)
}
