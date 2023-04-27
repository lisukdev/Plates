package workoutTemplates

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/lisukdev/Plates/pkg/domain/interfaces"
	"github.com/lisukdev/Plates/pkg/domain/workout"
	"log"
)

func ListAllItems(client interfaces.DbScanner) ([]workout.TemplateMetadata, error) {
	out, err := client.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: tableName,
	})
	if err != nil {
		return nil, err
	}
	var templateWorkouts = make([]workout.TemplateMetadata, 0)
	for _, item := range out.Items {
		log.Printf("Found item: %v", item)
		unMarshaledTemplateWorkout := storedTemplateWorkout{}
		log.Printf("Pre unmarshalled item: %v", unMarshaledTemplateWorkout)
		err = attributevalue.UnmarshalMap(item, &unMarshaledTemplateWorkout)
		log.Printf("Unmarshalled item: %v", unMarshaledTemplateWorkout)
		if err != nil {
			return nil, err
		}
		templateWorkout, err := toDomain(&unMarshaledTemplateWorkout)
		if err != nil {
			return nil, err
		}
		templateWorkouts = append(templateWorkouts, *templateWorkout.Metadata())
	}
	log.Printf("Final map: %v", templateWorkouts)
	return templateWorkouts, nil
}

func GetItem(client interfaces.DbGetter, templateId uuid.UUID) (*workout.TemplateWorkout, error) {
	key := map[string]types.AttributeValue{"Id": &types.AttributeValueMemberS{Value: templateId.String()}}

	out, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: tableName,
		Key:       key,
	})

	if err != nil {
		return nil, err
	}

	if out.Item == nil {
		return nil, errors.New("Item not found, key: " + templateId.String())
	}

	unMarshaledTemplateWorkout := storedTemplateWorkout{}
	err = attributevalue.UnmarshalMap(out.Item, &unMarshaledTemplateWorkout)
	if err != nil {
		return nil, err
	}

	return toDomain(&unMarshaledTemplateWorkout)
}

func TransactionPutItem(item *workout.TemplateWorkout) (*types.TransactWriteItem, error) {
	row := toStored(item)
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
