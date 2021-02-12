package main

import (
        "fmt"
        "context"
        "github.com/aws/aws-lambda-go/lambda"
)

type RpnExpression struct {
        Expression string `json:"expression"`
}

func HandleRequest(ctx context.Context, input RpnExpression) (string, error) {
	var result, _ = ReversePolishNotation(input.Expression)
	return fmt.Sprintf("%f", result), nil
}

func main() {
        lambda.Start(HandleRequest)
}
