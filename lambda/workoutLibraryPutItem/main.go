package main

import (
	"context"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/lisukdev/Plates/pkg/domain/workout"
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

func handleRequest(ctx context.Context, workout workout.TemplateWorkout) (*workout.TemplateWorkout, error) {
	lc, _ := lambdacontext.FromContext(ctx)
	library, err := buildClient(ctx)
	if err != nil {
		return nil, err
	}

	return library.AddWorkoutTemplate(lc.Identity.CognitoIdentityID, workout)
}

func main() {
	runtime.Start(handleRequest)
}
