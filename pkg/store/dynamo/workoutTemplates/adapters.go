package workoutTemplates

import (
	"errors"
	"github.com/google/uuid"
	"github.com/lisukdev/Plates/pkg/domain/workout"
)

func toStored(workout *workout.TemplateWorkout) *storedTemplateWorkout {
	return &storedTemplateWorkout{
		Id:                workout.Id.String(),
		SchemaVersion:     currentSchemaVersion,
		ObjectVersion:     workout.Version,
		Name:              workout.Name,
		Creator:           workout.Creator,
		CreationTimestamp: workout.CreationTimestamp,
		UpdatedTimestamp:  workout.UpdatedTimestamp,
	}
}

func toDomain(stored *storedTemplateWorkout) (*workout.TemplateWorkout, error) {
	if stored.SchemaVersion != currentSchemaVersion {
		return nil, errors.New("Schema version mismatch")
	}
	domainId, err := uuid.Parse(stored.Id)
	if err != nil {
		return nil, err
	}
	return &workout.TemplateWorkout{
		Id:                domainId,
		Version:           stored.ObjectVersion,
		Name:              stored.Name,
		Creator:           stored.Creator,
		CreationTimestamp: stored.CreationTimestamp,
		UpdatedTimestamp:  stored.UpdatedTimestamp,
	}, nil
}
