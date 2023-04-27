package workoutTemplateLibraries

import (
	"errors"
	"github.com/google/uuid"
	"github.com/lisukdev/Plates/pkg/domain/workout"
)

func toStored(libraryId uuid.UUID, metadata *workout.TemplateMetadata) *storedRow {
	return &storedRow{
		LibraryId:         libraryId.String(),
		TemplateId:        metadata.Id.String(),
		SchemaVersion:     currentSchemaVersion,
		ObjectVersion:     metadata.Version,
		Creator:           metadata.Creator,
		CreationTimestamp: metadata.CreationTimestamp,
		UpdatedTimestamp:  metadata.UpdatedTimestamp,
	}
}

func toDomain(stored storedRow) (*workout.TemplateWorkout, error) {
	if stored.SchemaVersion != currentSchemaVersion {
		return nil, errors.New("Schema version mismatch")
	}
	domainId, err := uuid.Parse(stored.TemplateId)
	if err != nil {
		return nil, err
	}
	return &workout.TemplateWorkout{
		Id:                domainId,
		Name:              stored.Name,
		Version:           stored.ObjectVersion,
		Creator:           stored.Creator,
		CreationTimestamp: stored.CreationTimestamp,
		UpdatedTimestamp:  stored.UpdatedTimestamp,
	}, nil
}
