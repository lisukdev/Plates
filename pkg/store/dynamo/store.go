package dynamo

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/lisukdev/Plates/pkg/domain/workout"
	"github.com/lisukdev/Plates/pkg/store/dynamo/workoutTemplates"
)

type DynamoWorkoutLibrary struct {
	DynamoDbClient *dynamodb.Client
}

func (library DynamoWorkoutLibrary) ListWorkoutTemplates(userId string) ([]workout.TemplateMetadata, error) {
	return workoutTemplates.ListAllItems(library.DynamoDbClient)
}

func (library DynamoWorkoutLibrary) AddWorkoutTemplate(userId string, templateWorkout *workout.TemplateWorkout) (*workout.TemplateWorkout, error) {
	txn := make([]types.TransactWriteItem, 1)

	putItem, err := workoutTemplates.TransactionPutItem(templateWorkout)
	if err != nil {
		return nil, err
	}
	txn = append(txn, *putItem)
	_, err = library.DynamoDbClient.TransactWriteItems(context.TODO(), &dynamodb.TransactWriteItemsInput{TransactItems: txn})
	if err != nil {
		return nil, err
	}
	return templateWorkout, nil
}

func (library DynamoWorkoutLibrary) GetWorkoutTemplate(templateId uuid.UUID) (*workout.TemplateWorkout, error) {
	return workoutTemplates.GetItem(library.DynamoDbClient, templateId)
}
