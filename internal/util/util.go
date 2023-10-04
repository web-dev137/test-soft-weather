package util

import (
	"strconv"
	"strings"
)

func ArithmeticSum(inp string) (int, error) {

	var arithmeticOperator string
	var res int
	inpExpression := strings.Fields(inp)
	res, _ = strconv.Atoi(string(inp[0]))
	for _, elExp := range inpExpression {
		switch elExp {
		case "+":
			arithmeticOperator = "+"
		case "-":
			arithmeticOperator = "-"
		default:
			number, err := strconv.Atoi(elExp)
			if err != nil {
				return res, err
			}
			if arithmeticOperator == "+" {
				res += number
			} else if arithmeticOperator == "-" {
				res -= number
			}
		}

	}
	return res, nil
}
