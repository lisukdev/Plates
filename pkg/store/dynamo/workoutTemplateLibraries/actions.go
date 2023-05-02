package workoutTemplateLibraries

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/lisukdev/Plates/pkg/domain/interfaces"
	"github.com/lisukdev/Plates/pkg/domain/workout"
	"log"
)

func ListLibraryItems(client interfaces.DbQueryer, libraryId uuid.UUID) ([]workout.TemplateMetadata, error) {
	queryExpression := "LibraryId = :libraryId"
	out, err := client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              tableName,
		KeyConditionExpression: &queryExpression,
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":libraryId": &types.AttributeValueMemberS{Value: libraryId.String()},
		},
	})
	if err != nil {
		return nil, err
	}
	var templateWorkouts = make([]workout.TemplateMetadata, 0)
	for _, item := range out.Items {
		unMarshaledTemplateWorkout := storedRow{}
		err = attributevalue.UnmarshalMap(item, &unMarshaledTemplateWorkout)
		if err != nil {
			return nil, err
		}
		templateWorkout, err := toDomain(&unMarshaledTemplateWorkout)
		if err != nil {
			return nil, err
		}
		templateWorkouts = append(templateWorkouts, *templateWorkout)
	}
	log.Printf("Final map: %v", templateWorkouts)
	return templateWorkouts, nil
}

func TransactionPutItem(libraryId uuid.UUID, item *workout.TemplateMetadata) (*types.TransactWriteItem, error) {
	row := toStored(libraryId, item)
	marshaledTemplateWorkout, err := attributevalue.MarshalMap(row)
	if err != nil {
		return nil, err
	}

	putRequest := types.Put{
		TableName: tableName,
		Item:      marshaledTemplateWorkout,
	}
	return &types.TransactWriteItem{Put: &putRequest}, nil
}
