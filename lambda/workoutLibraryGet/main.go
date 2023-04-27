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

	domainList, err := service.ListTemplates(lc.Identity.CognitoIdentityID)
	if err != nil {
		return handleError(err)
	}

	var responseList []api.WorkoutMetadata
	for _, metadata := range domainList {
		responseList = append(responseList, adapters.TemplateMetadataToApi(&metadata))
	}
	responseBody, err := json.Marshal(responseList)
	if err != nil {
		return handleError(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode:        200,
		Headers:           nil,
		MultiValueHeaders: nil,
		Body:              string(responseBody),
		IsBase64Encoded:   false,
	}, nil
}

func main() {
	runtime.Start(handleRequest)
}
