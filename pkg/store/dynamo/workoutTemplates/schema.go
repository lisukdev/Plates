package workoutTemplates

import "github.com/aws/aws-sdk-go-v2/aws"

var tableName = aws.String("WorkoutTemplates")

const currentSchemaVersion = 0

type storedTemplateWorkout struct {
	Id                string
	SchemaVersion     int
	ObjectVersion     int
	Name              string
	Creator           string
	CreationTimestamp string
	UpdatedTimestamp  string

	Exercises []storedTemplateWorkout
}

type storedTemplateExercise struct {
}
