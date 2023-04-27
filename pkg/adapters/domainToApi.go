package adapters

import (
	"github.com/lisukdev/Plates/api"
	"github.com/lisukdev/Plates/pkg/domain/workout"
)

func TemplateWorkoutToApi(workout *workout.TemplateWorkout) api.WorkoutTemplate {
	workoutIdString := workout.Id.String()
	versionInt := int32(workout.Version)
	return api.WorkoutTemplate{
		Id:      &workoutIdString,
		Name:    &workout.Name,
		Version: &versionInt,
	}
}
