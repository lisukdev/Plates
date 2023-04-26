package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/lisukdev/Plates/api"
	"github.com/lisukdev/Plates/pkg/domain/workout"
	"github.com/lisukdev/Plates/pkg/store"
	"time"
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
	requestBodyWorkout := api.WorkoutTemplate{}
	err := json.Unmarshal([]byte(request.Body), &requestBodyWorkout)
	if err != nil {
		return handleError(err)
	}

	library, err := buildClient(ctx)
	if err != nil {
		return nil, err
	}

	lc, _ := lambdacontext.FromContext(ctx)

	timestamp := time.Now().Format(time.RFC3339)
	workout := workout.TemplateWorkout{
		Id:                uuid.UUID{},
		Name:              *requestBodyWorkout.Name,
		Version:           int(*requestBodyWorkout.Version),
		Creator:           lc.Identity.CognitoIdentityID,
		CreationTimestamp: timestamp,
		UpdatedTimestamp:  timestamp,
	}
	storedTemplate, err := library.AddWorkoutTemplate(lc.Identity.CognitoIdentityID, workout)
	if err != nil {
		return handleError(err)
	}

	responseTemplateId := storedTemplate.Id.String()
	responseTemplateVersion := int32(storedTemplate.Version)
	responseTemplate := api.WorkoutTemplate{
		Id:      &responseTemplateId,
		Name:    &storedTemplate.Name,
		Version: &responseTemplateVersion,
	}

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
