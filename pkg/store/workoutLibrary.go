package store

import (
	"github.com/lisukdev/Plates/pkg/domain/workout"
)

type WorkoutLibrary interface {
	listWorkoutTemplates(userId string) ([]workout.TemplateMetadata, error)
}
