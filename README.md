
go get "github.com/aws/aws-lambda-go/lambda"
GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main




{
    "Role": {
        "Path": "/",
        "RoleName": "lambda-basic-execution",
        "RoleId": "AROA3VNUU4JPFUQ2H5FXC",
        "Arn": "arn:aws:iam::801927127646:role/lambda-basic-execution",
        "CreateDate": "2021-02-11T22:11:42+00:00",
        "AssumeRolePolicyDocument": {
            "Version": "2012-10-17",
            "Statement": {
                "Effect": "Allow",
                "Principal": {
                    "Service": "lambda.amazonaws.com"
                },
                "Action": "sts:AssumeRole"
            }
        }
    }
}


aws lambda create-function --function-name helloworld_go --zip-file fileb://main.zip --handler main --runtime go1.x --role "arn:aws:iam::801927127646:role/lambda-basic-execution"


aws lambda invoke --function-name helloworld_go --invocation-type "RequestResponse" --payload '{"name": "af3"}' --cli-binary-format raw-in-base64-out  response.txt
mikael@xps:~/dev/dnb$ cat response.txt 




