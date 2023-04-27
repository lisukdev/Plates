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
	"github.com/lisukdev/Plates/pkg/domain"
	"github.com/lisukdev/Plates/pkg/domain/workout"
	"github.com/lisukdev/Plates/pkg/store"
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
		WorkoutLibraryRepository: store.DynamoWorkoutLibrary{DynamoDbClient: client},
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

type ApiRequest struct {
	UserId   string
	Template api.WorkoutTemplate
}

func (request ApiRequest) GetName() string {
	return *request.Template.Name
}
func (request ApiRequest) GetCreator() string {
	return request.UserId
}
func (request ApiRequest) GetExercises() []workout.TemplateExercise {
	result := make([]workout.TemplateExercise, len(request.Template.Exercises))
	return result
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	lc, _ := lambdacontext.FromContext(ctx)

	requestBodyWorkout := api.WorkoutTemplate{}
	err := json.Unmarshal([]byte(request.Body), &requestBodyWorkout)
	if err != nil {
		return handleError(err)
	}

	apiRequest := ApiRequest{
		UserId:   lc.Identity.CognitoIdentityID,
		Template: requestBodyWorkout,
	}

	storedTemplate, err := service.CreateTemplateInLibrary(lc.Identity.CognitoIdentityID, apiRequest)
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
