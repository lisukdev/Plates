package store

import (
	"github.com/google/uuid"
	"github.com/lisukdev/Plates/pkg/domain/workout"
)

type WorkoutLibrary interface {
	ListWorkoutTemplates(userId string) ([]workout.TemplateMetadata, error)
	AddWorkoutTemplate(userId string, templateWorkout *workout.TemplateWorkout) (*workout.TemplateWorkout, error)
	GetWorkoutTemplate(templateId uuid.UUID) (*workout.TemplateWorkout, error)
}
