# Implementation:
* Reverse polish notation calculator from https://gist.github.com/lwiecek/952b6bca77c843e83c5d5577da9fd37e


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


# Local build
```sh
#go get "github.com/aws/aws-lambda-go/lambda"
export GOPATH=`pwd`
rm main.zip || true; GOOS=linux GOARCH=amd64 go build -o main cmd/main; zip main.zip main; rm main || true
aws s3 cp main.zip s3://mikifi-deploy/
```
# Deploy

## Local deploy

```sh
aws lambda create-function --function-name simple_calc --zip-file fileb://main.zip --handler main --runtime go1.x --role "<an existing role>"
```
### Delete

```sh
aws lambda delete-function --function-name simple_calc                
```

## CloudFormation deploy

This [deploy](src/build/cf_functions.yaml) assumes the lambda artifact is available as `main.zip` in the `mikifi-deploy` bucket (see local build above).

```sh
aws cloudformation deploy --template-file src/build/cf_functions.yaml --stack-name simple-calc --capabilities CAPABILITY_NAMED_IAM
```

### Delete
```sh
aws cloudformation delete-stack --stack-name simple-calc
```


# Invoke

```sh
aws lambda invoke --function-name simple_calc --invocation-type "RequestResponse" --payload '{"expression": "3 3 *"}' --cli-binary-format raw-in-base64-out response.json; cat response.json 
```

# TODO
* Adding unit tests and functional tests
* ci/cd pipeline
