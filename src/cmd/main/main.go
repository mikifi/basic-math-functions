package main

import (
        "context"
        rpn "internal/reverspolishnotation"
        "github.com/aws/aws-lambda-go/lambda"
)

type RpnExpression struct {
        Expression string `json:"expression"`
}

type CalcResult struct {
	Result float64 `json:"result"`
}

func HandleRequest(ctx context.Context, input RpnExpression) (CalcResult, error) {
	var result, err = rpn.ReversePolishNotation(input.Expression)
	return CalcResult{Result: result}, err
}

func main() {
        lambda.Start(HandleRequest)
}
