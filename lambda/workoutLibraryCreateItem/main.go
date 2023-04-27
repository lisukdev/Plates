package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/lisukdev/Plates/api"
	"github.com/lisukdev/Plates/pkg/adapters"
	"github.com/lisukdev/Plates/pkg/domain"
	"github.com/lisukdev/Plates/pkg/store/dynamo"
	"log"
)

var service domain.WorkoutLibraryService

func init() {
	myConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	client := dynamodb.NewFromConfig(myConfig)
	service = domain.WorkoutLibraryService{
		WorkoutLibraryRepository: dynamo.DynamoWorkoutLibrary{DynamoDbClient: client},
	}
}

func handleError(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode:        500,
		Headers:           nil,
		MultiValueHeaders: nil,
		Body:              err.Error(),
		IsBase64Encoded:   false,
	}, err
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	lc, _ := lambdacontext.FromContext(ctx)

	requestBody := api.CreateWorkoutTemplateRequest{}
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		return handleError(err)
	}
	apiRequest := adapters.CreateWorkoutTemplateRequestAdapter{
		UserId:  lc.Identity.CognitoIdentityID,
		Request: &requestBody,
	}

	createdTemplate, err := service.CreateTemplateInLibrary(lc.Identity.CognitoIdentityID, apiRequest)

	if err != nil {
		return handleError(err)
	}
	responseTemplate := adapters.TemplateWorkoutToApi(createdTemplate)
	marshaledResponse, err := responseTemplate.MarshalJSON()
	if err != nil {
		return handleError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           nil,
		MultiValueHeaders: nil,
		Body:              string(marshaledResponse),
		IsBase64Encoded:   false,
	}, nil
}

func main() {
	runtime.Start(handleRequest)
}
