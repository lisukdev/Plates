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
	lc, _ := lambdacontext.FromContext(ctx)
	library, err := buildClient(ctx)
	if err != nil {
		return handleError(err)
	}
	workouts, err := library.ListWorkoutTemplates(lc.Identity.CognitoIdentityID)
	if err != nil {
		return handleError(err)
	}
	var responseList []api.WorkoutMetadata
	for _, workout := range workouts {
		workoutIdString := workout.Id.String()
		workoutVersion := int32(workout.Version)
		responseList = append(responseList, api.WorkoutMetadata{
			Id:      &workoutIdString,
			Name:    &workout.Name,
			Version: &workoutVersion,
		})
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
