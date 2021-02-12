# Implementation:
* Reverse polish notation calculater from https://gist.github.com/lwiecek/952b6bca77c843e83c5d5577da9fd37e

# Prerequisites:
* go
* aws cli 
* aws ami role (take note of created arn):
```sh
aws iam create-role --role-name lambda-basic-execution--assume-role-policy-document file://lambda-trust-policy.json
aws iam attach-role-policy --role-name lambda-basic-execution --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
```


# Local build:
```sh
#go mod init github.com/mikifi/basic-math-functions        
go get "github.com/aws/aws-lambda-go/lambda"
rm main.zip || true; GOOS=linux GOARCH=amd64 go build -o main .; zip main.zip main

```

# Local deploy and invoke

```sh
aws lambda create-function --function-name simple_calc --zip-file fileb://main.zip --handler main --runtime go1.x --role "arn:aws:iam::801927127646:role/lambda-basic-execution"

aws lambda invoke --function-name simple_calc --invocation-type "RequestResponse" --payload '{"name": "af3"}' --cli-binary-format raw-in-base64-out  response.txt

cat response.txt 
```



