package domain

import (
	"github.com/google/uuid"
	"github.com/lisukdev/Plates/pkg/domain/user"
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

func (service *WorkoutLibraryService) CreateTemplateInLibrary(user *user.AuthorizedUser, request CreateTemplateRequest) (*workout.TemplateWorkout, error) {
	newTemplate, err := workout.NewTemplate(request.GetName(), request.GetCreator(), request.GetExercises())
	if err != nil {
		return nil, err
	}
	return service.WorkoutLibraryRepository.AddWorkoutTemplate(user.Id, newTemplate)
}

func (service *WorkoutLibraryService) GetTemplate(user *user.AuthorizedUser, templateId uuid.UUID) (*workout.TemplateWorkout, error) {
	return service.WorkoutLibraryRepository.GetWorkoutTemplate(templateId)
}

func (service *WorkoutLibraryService) ListTemplates(user *user.AuthorizedUser) ([]workout.TemplateMetadata, error) {
	return service.WorkoutLibraryRepository.ListWorkoutTemplates(user.Id)
}
