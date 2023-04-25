#!/bin/bash
set -eo pipefail
ARTIFACT_BUCKET=$(cat bucket-name.txt)
cd lambda
cd function
GOARCH=amd64 GOOS=linux go build main.go
cd ../
cd workoutLibraryGet
GOARCH=amd64 GOOS=linux go build main.go
cd ../
cd workoutLibraryGetItem
GOARCH=amd64 GOOS=linux go build main.go
cd ../
cd workoutLibraryPutItem
GOARCH=amd64 GOOS=linux go build main.go
cd ../
cd ../
aws cloudformation package --template-file template.yml --s3-bucket $ARTIFACT_BUCKET --output-template-file out.yml
aws cloudformation deploy --template-file out.yml --stack-name blank-go --capabilities CAPABILITY_NAMED_IAM
