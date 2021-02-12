package main

import (
        "fmt"
        "context"
        "github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
        Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	var result, _ = ReversePolishNotation("2 2 +")
	return fmt.Sprintf("Calc: %f!", result), nil
}

func main() {
        lambda.Start(HandleRequest)
}
