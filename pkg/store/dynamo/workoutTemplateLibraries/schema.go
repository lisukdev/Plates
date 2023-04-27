package workoutTemplateLibraries

import "github.com/aws/aws-sdk-go-v2/aws"

var tableName = aws.String("WorkoutTemplateLibraries")

const currentSchemaVersion = 0

type storedRow struct {
	LibraryId     string
	TemplateId    string
	SchemaVersion int
	ObjectVersion int

	Name              string
	Creator           string
	CreationTimestamp string
	UpdatedTimestamp  string
}
