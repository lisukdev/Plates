package main

import (
	"context"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/google/uuid"
	"github.com/lisukdev/Plates/pkg/domain/workout"
	"time"
)

type LocalWorkoutLibrary struct {
	Workouts []workout.TemplateMetadata
}

func (library LocalWorkoutLibrary) listWorkoutTemplates(userId string) ([]workout.TemplateMetadata, error) {
	return library.Workouts, nil
}

func handleRequest(ctx context.Context) ([]workout.TemplateMetadata, error) {
	lc, _ := lambdacontext.FromContext(ctx)
	now := time.Now().Format(time.RFC3339)
	library := LocalWorkoutLibrary{
		Workouts: []workout.TemplateMetadata{
			{
				Id:                uuid.New(),
				Name:              "My Template",
				Version:           1,
				Creator:           lc.Identity.CognitoIdentityID,
				CreationTimestamp: now,
				UpdatedTimestamp:  now,
			},
			{
				Id:                uuid.New(),
				Name:              "My Second Template",
				Version:           1,
				Creator:           lc.Identity.CognitoIdentityID,
				CreationTimestamp: now,
				UpdatedTimestamp:  now,
			},
		},
	}
	return library.listWorkoutTemplates(lc.Identity.CognitoIdentityID)
}

func main() {
	runtime.Start(handleRequest)
}
