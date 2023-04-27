package adapters

import (
	"github.com/google/uuid"
	"github.com/lisukdev/Plates/pkg/domain/workout"
	"log"
	"testing"
)

func TestTemplateMetadataToApi(t *testing.T) {
	metadata := workout.TemplateMetadata{
		Id:                uuid.Nil,
		Name:              "Test Name",
		Version:           0,
		Creator:           "Test User",
		CreationTimestamp: "Test Creation Date",
		UpdatedTimestamp:  "Test Updated Date",
	}

	apiMetadata := TemplateMetadataToApi(&metadata)
	if *apiMetadata.Id != "00000000-0000-0000-0000-000000000000" {
		t.Errorf("Expected Id to be 00000000-0000-0000-0000-000000000000, got %s", *apiMetadata.Id)
	}
	if *apiMetadata.Name != "Test Name" {
		t.Errorf("Expected Name to be Test Name, got %s", *apiMetadata.Name)
	}
}

func TestTemplateMetadataListToApiList(t *testing.T) {
	metadataList := []workout.TemplateMetadata{
		{
			Id:                uuid.MustParse("94d04d64-224c-455f-a6b0-61dfb71a223a"),
			Name:              "Test Name 1",
			Version:           0,
			Creator:           "Test User 1",
			CreationTimestamp: "Test Creation Date 1",
			UpdatedTimestamp:  "Test Updated Date 1",
		},
		{
			Id:                uuid.MustParse("039b9176-235f-4393-826f-8435bb7ad3b8"),
			Name:              "Test Name 2",
			Version:           1,
			Creator:           "Test User 2",
			CreationTimestamp: "Test Creation Date 2",
			UpdatedTimestamp:  "Test Updated Date 2",
		},
	}
	apiList := TemplateMetadataListToApiList(metadataList)
	log.Printf("apiList: %v", apiList)
	if *apiList[0].Id != "94d04d64-224c-455f-a6b0-61dfb71a223a" {
		t.Errorf("Expected Id to be 94d04d64-224c-455f-a6b0-61dfb71a223a, got %s", *apiList[0].Id)
	}
	if *apiList[0].Name != "Test Name 1" {
		t.Errorf("Expected Name to be Test Name 1, got %s", *apiList[0].Name)
	}
	if *apiList[1].Id != "039b9176-235f-4393-826f-8435bb7ad3b8" {
		t.Errorf("Expected Id to be 039b9176-235f-4393-826f-8435bb7ad3b8, got %s", *apiList[1].Id)
	}
	if *apiList[1].Name != "Test Name 2" {
		t.Errorf("Expected Name to be Test Name 2, got %s", *apiList[1].Name)
	}
}
