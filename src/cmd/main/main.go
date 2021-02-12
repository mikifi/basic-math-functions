package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	rpn "internal/reverspolishnotation"
)

// RpnExpression is the input DTO
type RpnExpression struct {
	Expression string `json:"expression"`
}

// CalcResult is the output DTO
type CalcResult struct {
	Result float64 `json:"result"`
}

// HandleRequest is the AWS lamdba handler
func HandleRequest(ctx context.Context, input RpnExpression) (CalcResult, error) {
	var result, err = rpn.ReversePolishNotation(input.Expression)
	return CalcResult{Result: result}, err
}

func main() {
	lambda.Start(HandleRequest)
}
