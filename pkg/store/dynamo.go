package store

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/lisukdev/Plates/pkg/domain/workout"
)

const (
	UserWorkoutTemplates string = "UserWorkoutTemplates"
	WorkoutTemplates     string = "WorkoutTemplates"
)

type storedWorkoutLibrary struct {
	Workouts []workout.TemplateMetadata `json:"workouts"`
}

func userTemplateLibraryKey(userId string) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{"UserId": &types.AttributeValueMemberS{Value: userId}}
}

func workoutTemplatesKey(templateId uuid.UUID) map[string]types.AttributeValue {
	return map[string]types.AttributeValue{"id": &types.AttributeValueMemberS{Value: templateId.String()}}
}

type DynamoWorkoutLibrary struct {
	DynamoDbClient *dynamodb.Client
}

type TemplateRow struct {
	TemplateId string
	Workout    workout.TemplateWorkout
}

func (library DynamoWorkoutLibrary) ListWorkoutTemplates(userId string) ([]workout.TemplateMetadata, error) {
	out, err := library.DynamoDbClient.Scan(context.TODO(), &dynamodb.ScanInput{TableName: aws.String(WorkoutTemplates)})
	if err != nil {
		return nil, err
	}
	fmt.Println(out.Items)
	return []workout.TemplateMetadata{}, nil
}

func (library DynamoWorkoutLibrary) AddWorkoutTemplate(userId string, templateWorkout workout.TemplateWorkout) (*workout.TemplateWorkout, error) {
	templateRow := TemplateRow{TemplateId: templateWorkout.Id.String(), Workout: templateWorkout}
	marshaledTemplateWorkout, err := attributevalue.MarshalMap(templateRow)
	if err != nil {
		return nil, err
	}
	out, err := library.DynamoDbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(WorkoutTemplates),
		Item:      marshaledTemplateWorkout,
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(out)
	unMarshaledTemplateWorkout := workout.TemplateWorkout{}
	err = attributevalue.UnmarshalMap(out.Attributes, &unMarshaledTemplateWorkout)
	if err != nil {
		return nil, err
	}
	return &unMarshaledTemplateWorkout, nil
}

func (library DynamoWorkoutLibrary) GetWorkoutTemplate(templateId uuid.UUID) (*workout.TemplateWorkout, error) {
	out, err := library.DynamoDbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(WorkoutTemplates),
		Key:       workoutTemplatesKey(templateId),
	})

	if err != nil {
		return nil, err
	}
	if out.Item == nil {
		return nil, errors.New("Item not found, key: " + templateId.String())
	}
	fmt.Println(out)
	unMarshaledTemplateWorkout := TemplateRow{}

	err = attributevalue.UnmarshalMap(out.Item, &unMarshaledTemplateWorkout)
	if err != nil {
		return nil, err
	}
	return &unMarshaledTemplateWorkout.Workout, nil
}
