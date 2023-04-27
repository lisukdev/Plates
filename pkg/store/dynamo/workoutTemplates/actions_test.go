package workoutTemplates

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"testing"
)

type mockDynamoDBClient struct {
	data []map[string]types.AttributeValue
}

func (c mockDynamoDBClient) Scan(_ context.Context, _ *dynamodb.ScanInput, _ ...func(*dynamodb.Options)) (
	*dynamodb.ScanOutput,
	error,
) {
	return &dynamodb.ScanOutput{
		Items: c.data,
	}, nil
}

func makeMockRow(
	Id string,
	ObjectVersion int,
	Name string,
	Creator string,
	CreationTimestamp string,
	UpdatedTimestamp string,
) map[string]types.AttributeValue {
	row := storedTemplateWorkout{
		Id:                Id,
		SchemaVersion:     currentSchemaVersion,
		ObjectVersion:     ObjectVersion,
		Name:              Name,
		Creator:           Creator,
		CreationTimestamp: CreationTimestamp,
		UpdatedTimestamp:  UpdatedTimestamp,
	}
	result, err := attributevalue.MarshalMap(row)
	if err != nil {
		panic("Error marshaling mock row: " + err.Error())
	}
	return result
}

func TestListAllItems(t *testing.T) {
	uuid1 := uuid.MustParse("3be34021-44bd-45b0-a3e0-c072f5a92f10")
	uuid2 := uuid.MustParse("36d39b9d-d77d-4854-a634-1036a74e6926")
	client := &mockDynamoDBClient{
		data: []map[string]types.AttributeValue{
			makeMockRow(uuid1.String(), 1, "Test 1", "Test User 1", "2021-01-01", "2021-01-01"),
			makeMockRow(uuid2.String(), 1, "Test 2", "Test User 2", "2021-01-02", "2021-01-02"),
		},
	}
	result, err := ListAllItems(client)
	if err != nil {
		t.Errorf("Error listing items: %v", err.Error())
	}
	if len(result) != 2 {
		t.Errorf("Expected 2 items, got: %v", len(result))
	}
	if result[0].Id != uuid1 {
		t.Errorf("Expected Id 1, got: %v", result[0].Id)
	}
	if result[0].Name != "Test 1" {
		t.Errorf("Expected Name 1, got: %v", result[0].Name)
	}
	if result[1].Id != uuid2 {
		t.Errorf("Expected Id 2, got: %v", result[1].Id)
	}
	if result[1].Name != "Test 2" {
		t.Errorf("Expected Name 2, got: %v", result[1].Name)
	}
}
