package dynamo

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/lisukdev/Plates/pkg/domain/workout"
	"github.com/lisukdev/Plates/pkg/store/dynamo/workoutTemplateLibraries"
	"github.com/lisukdev/Plates/pkg/store/dynamo/workoutTemplates"
	"log"
)

type DynamoWorkoutLibrary struct {
	DynamoDbClient *dynamodb.Client
}

func (library DynamoWorkoutLibrary) ListWorkoutTemplates(libraryId uuid.UUID) ([]workout.TemplateMetadata, error) {
	return workoutTemplateLibraries.ListLibraryItems(library.DynamoDbClient, libraryId)
}

func (library DynamoWorkoutLibrary) AddWorkoutTemplate(libraryId uuid.UUID, templateWorkout *workout.TemplateWorkout) (*workout.TemplateWorkout, error) {
	putTemplate, err := workoutTemplates.TransactionPutItem(templateWorkout)
	if err != nil {
		return nil, err
	}

	putMetadata, err := workoutTemplateLibraries.TransactionPutItem(libraryId, templateWorkout.Metadata())
	if err != nil {
		return nil, err
	}

	txn := []types.TransactWriteItem{*putTemplate, *putMetadata}
	log.Printf("Transaction: %+v", txn)

	_, err = library.DynamoDbClient.TransactWriteItems(context.TODO(), &dynamodb.TransactWriteItemsInput{TransactItems: txn})
	if err != nil {
		return nil, err
	}
	return templateWorkout, nil
}

func (library DynamoWorkoutLibrary) GetWorkoutTemplate(templateId uuid.UUID) (*workout.TemplateWorkout, error) {
	return workoutTemplates.GetItem(library.DynamoDbClient, templateId)
}
