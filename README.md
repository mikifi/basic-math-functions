# Implementation:
* Reverse polish notation calculater from https://gist.github.com/lwiecek/952b6bca77c843e83c5d5577da9fd37e


## Input

The lamdba takes a reverse polish notation expression as input (as json), e.g:

```json
{"expression": "3 2 /"}
```
## Output

The output is either an error object or a json with the result, e.g:

```json
{"result":1.5}
```


# Prerequisites:
* go
* aws cli 
* aws ami role (take note of created arn):
```sh
aws iam create-role --role-name lambda-basic-execution--assume-role-policy-document file://src/build/lambda-trust-policy.json
aws iam attach-role-policy --role-name lambda-basic-execution --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
```


# Local build:
```sh
#go get "github.com/aws/aws-lambda-go/lambda"
export GOPATH=`pwd`
rm main.zip || true; GOOS=linux GOARCH=amd64 go build -o main cmd/main; zip main.zip main; rm main || true
aws s3 cp main.zip s3://mikifi-deploy/
```

# Local deploy and invoke

```sh
aws lambda create-function --function-name simple_calc --zip-file fileb://main.zip --handler main --runtime go1.x --role "arn:aws:iam::801927127646:role/lambda-basic-execution"

aws lambda invoke --function-name simple_calc --invocation-type "RequestResponse" --payload '{"expression": "3 3 *"}' --cli-binary-format raw-in-base64-out  response.txt

cat response.txt 
```

# CloudFormation deploy
```sh
aws cloudformation deploy --template-file src/build/cf_functions.yaml --stack-name simple-calc
```

# TODO
* Adding bucket, zip (copy) and iam role to cloudformation config
* Adding unit tests and functional tests
* ci/cd pipeline
* Either json or yaml config, not both
* go mod?
