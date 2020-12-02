package formula

import (
	"fmt"
	"strings"
	"testing"
	"text/scanner"
)

func init() {
}

func TestLogic(t *testing.T) {
	const src = `
SELECT name,age FROM MATCH("xxx","dfdf") WHERE uid = MATCH("xxx","ddd")
`

	//fnStart:=-1
	//sql:=""
	stack := make([]string, 0)
	var s scanner.Scanner
	s.Mode = scanner.GoTokens
	s.Init(strings.NewReader(src))
	s.Filename = "example"
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		str := s.TokenText()
		if str == "," {
			//fmt.Println(stack[len(stack)-1])
			str = stack[len(stack)-1] + ","
			stack[len(stack)-1] = str
		} else if str == "(" {
			//fnStart = len(stack)-1
			stack = append(stack, str)
		} else if str == ")" {
			stack = append(stack, str)
			//fn:=stack[fnStart:len(stack)]
			//fmt.Println(strings.Join(fn,""))
			//val,_:=NewCalculator(map[string]interface{}{}).Eval(strings.Join(fn,""))
			val, err := NewCalculator(map[string]interface{}{}).Eval(`MATCH("fdsf","fdsf")"fdsf"`)
			fmt.Println(val, err)
			//stack = append(stack[:fnStart],val.(bool))
			//fnStart=-1
		} else {
			stack = append(stack, str)
		}
		//fmt.Println(stack,str)
		//if str == "," {
		//    str = stack[len(stack)-2:len(stack)-1][0]+","
		//    stack[len(stack)-1] = str
		//}
	}
	fmt.Println(strings.Join(stack, " "))
}
