package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/lisukdev/Plates/api"
	"github.com/lisukdev/Plates/pkg/store"
)

func buildClient(ctx context.Context) (*store.DynamoWorkoutLibrary, error) {
	myConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	client := dynamodb.NewFromConfig(myConfig)
	return &store.DynamoWorkoutLibrary{DynamoDbClient: client}, nil
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
	fmt.Println(request.PathParameters)
	library, err := buildClient(ctx)
	if err != nil {
		return handleError(err)
	}
	templateId := request.PathParameters["templateId"]
	templateUuid, err := uuid.Parse(templateId)
	if err != nil {
		return handleError(err)
	}
	template, err := library.GetWorkoutTemplate(templateUuid)
	if err != nil {
		return handleError(err)
	}
	templateIdString := template.Id.String()
	templateVersionInt := int32(template.Version)
	response := api.WorkoutTemplate{
		Id:      &templateIdString,
		Name:    &template.Name,
		Version: &templateVersionInt,
	}

	marshaledResponse, err := response.MarshalJSON()
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
