package main

import (
        "context"
        "github.com/aws/aws-lambda-go/lambda"
)

type RpnExpression struct {
        Expression string `json:"expression"`
}

type CalcResult struct {
	Result float64 `json:"result"`
}

func HandleRequest(ctx context.Context, input RpnExpression) (CalcResult, error) {
	var result, err = ReversePolishNotation(input.Expression)
	return CalcResult{Result: result}, err
}

func main() {
        lambda.Start(HandleRequest)
}
