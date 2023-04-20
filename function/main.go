package main

import (
	"context"
	"fmt"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

type MyRequest struct {
	Name string `json:"name"`
}

func handleRequest(ctx context.Context, request MyRequest) (string, error) {
	lc, _ := lambdacontext.FromContext(ctx)
	return fmt.Sprintf("Hello %s, yourd id is %s", request.Name, lc.Identity.CognitoIdentityID), nil
}

func main() {
	runtime.Start(handleRequest)
}
