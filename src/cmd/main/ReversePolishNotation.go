package main

import (
	"fmt"
	"strconv"
	"strings"
)

func pop2(stack []float64) ([]float64, float64, float64) {
	var ab []float64
	stack, ab = stack[:len(stack)-2], stack[len(stack)-2:]
	return stack, ab[0], ab[1]
}

func ReversePolishNotation(s string) (float64, error) {
	var stack []float64
	var tokens = strings.Fields(s)
	var value float64
	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			var a, b float64
			if len(stack) < 2 {
				return 0, fmt.Errorf("not enough elmenents on the stack to calculate %s", token)
			}
			stack, a, b = pop2(stack)
			switch token {
			case "+":
				value = a + b
			case "-":
				value = a - b
			case "*":
				value = a * b
			case "/":
				value = a / b
			}
		default:
			var err error
			value, err = strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid token %s", token)
			}
		}
		stack = append(stack, value)
	}
	if len(stack) != 1 {
		return 0, fmt.Errorf("incomplete expression %s", stack)
	}
	return stack[0], nil
}


