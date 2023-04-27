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

func TemplateMetadataToApi(metadata *workout.TemplateMetadata) api.WorkoutMetadata {
	idString := metadata.Id.String()
	nameString := metadata.Name
	versionInt := int32(metadata.Version)
	return api.WorkoutMetadata{
		Id:      &idString,
		Name:    &nameString,
		Version: &versionInt,
	}
}

func TemplateMetadataListToApiList(metadataList []workout.TemplateMetadata) []api.WorkoutMetadata {
	var apiList []api.WorkoutMetadata
	for _, metadata := range metadataList {
		apiList = append(apiList, TemplateMetadataToApi(&metadata))
	}
	return apiList
}
