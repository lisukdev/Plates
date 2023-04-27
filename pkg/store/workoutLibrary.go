package store

import (
	"github.com/google/uuid"
	"github.com/lisukdev/Plates/pkg/domain/workout"
)

type WorkoutLibrary interface {
	ListWorkoutTemplates(libraryId uuid.UUID) ([]workout.TemplateMetadata, error)
	AddWorkoutTemplate(libraryId uuid.UUID, templateWorkout *workout.TemplateWorkout) (*workout.TemplateWorkout, error)
	GetWorkoutTemplate(templateId uuid.UUID) (*workout.TemplateWorkout, error)
}
