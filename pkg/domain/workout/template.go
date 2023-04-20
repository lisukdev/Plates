package workout

import "github.com/google/uuid"

type TemplateMetadata struct {
	Id                uuid.UUID `json:"id"`
	Name              string    `json:"name"`
	Version           int       `json:"version"`
	Creator           string    `json:"creator"`
	CreationTimestamp string    `json:"creationTimestamp"`
	UpdatedTimestamp  string    `json:"updatedTimestamp"`
}

type TemplateWorkout struct {
	Id                uuid.UUID `json:"id"`
	Name              string    `json:"name"`
	Version           int       `json:"version"`
	Creator           string    `json:"creator"`
	CreationTimestamp string    `json:"creationTimestamp"`
	UpdatedTimestamp  string    `json:"updatedTimestamp"`

	Exercises []TemplateExercise `json:"exercises"`
}

type TemplateExercise struct {
	Name              string        `json:"name"`
	Note              string        `json:"note"`
	Tempo             Tempo         `json:"tempo"`
	TargetRestSeconds int           `json:"restSeconds"`
	Sets              []TemplateSet `json:"sets"`
}

type TemplateSet struct {
	Target TargetSet `json:"target"`
}
