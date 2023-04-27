package workoutTemplateLibraries

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/lisukdev/Plates/pkg/domain/workout"
)

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
