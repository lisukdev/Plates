package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

type MyRequest struct {
	Name string `json:"name"`
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	requestBody := MyRequest{}
	json.Unmarshal([]byte(request.Body), &requestBody)
	lc, _ := lambdacontext.FromContext(ctx)
	responseBody := fmt.Sprintf("Hello %s, yourd id is %s", requestBody.Name, lc.Identity.CognitoIdentityID)
	return events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           nil,
		MultiValueHeaders: nil,
		Body:              responseBody,
		IsBase64Encoded:   false,
	}, nil
}

func main() {
	runtime.Start(handleRequest)
}
