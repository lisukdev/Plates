package domain

import (
	"github.com/google/uuid"
	"github.com/lisukdev/Plates/pkg/domain/workout"
	"github.com/lisukdev/Plates/pkg/store"
)

type WorkoutLibraryService struct {
	WorkoutLibraryRepository store.WorkoutLibrary
}

type CreateTemplateRequest interface {
	GetName() string
	GetCreator() string
	GetExercises() []workout.TemplateExercise
}

func (service *WorkoutLibraryService) CreateTemplateInLibrary(libraryId string, request CreateTemplateRequest) (*workout.TemplateWorkout, error) {
	newTemplate, err := workout.NewTemplate(request.GetName(), request.GetCreator(), request.GetExercises())
	if err != nil {
		return nil, err
	}
	return service.WorkoutLibraryRepository.AddWorkoutTemplate(uuid.Nil, newTemplate)
}

func (service *WorkoutLibraryService) GetTemplate(templateId uuid.UUID) (*workout.TemplateWorkout, error) {
	return service.WorkoutLibraryRepository.GetWorkoutTemplate(templateId)
}

func (service *WorkoutLibraryService) ListTemplates(libraryId string) ([]workout.TemplateMetadata, error) {
	return service.WorkoutLibraryRepository.ListWorkoutTemplates(uuid.Nil)
}
