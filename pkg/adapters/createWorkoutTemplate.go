package adapters

import (
	"github.com/lisukdev/Plates/api"
	"github.com/lisukdev/Plates/pkg/domain/user"
	"github.com/lisukdev/Plates/pkg/domain/workout"
)

type CreateWorkoutTemplateRequestAdapter struct {
	User    user.AuthorizedUser
	Request *api.CreateWorkoutTemplateRequest
}

func (request CreateWorkoutTemplateRequestAdapter) GetName() string {
	return *request.Request.Name
}
func (request CreateWorkoutTemplateRequestAdapter) GetCreator() string {
	return request.User.Id.String()
}
func (request CreateWorkoutTemplateRequestAdapter) GetExercises() []workout.TemplateExercise {
	result := make([]workout.TemplateExercise, len(request.Request.Exercises))
	return result
}
