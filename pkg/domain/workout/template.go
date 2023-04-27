package workout

import (
	"github.com/google/uuid"
	"time"
)

func NewTemplate(name string, creator string, exercises []TemplateExercise) (*TemplateWorkout, error) {
	id, err := uuid.NewRandom()
	timestamp := time.Now().Format(time.RFC3339)
	if err != nil {
		return nil, err
	}
	return &TemplateWorkout{
		Id:                id,
		Name:              name,
		Version:           0,
		Creator:           creator,
		CreationTimestamp: timestamp,
		UpdatedTimestamp:  timestamp,
		Exercises:         exercises,
	}, nil
}

func (t TemplateWorkout) Metadata() TemplateMetadata {
	return TemplateMetadata{
		Id:                t.Id,
		Name:              t.Name,
		Version:           t.Version,
		Creator:           t.Creator,
		CreationTimestamp: t.CreationTimestamp,
		UpdatedTimestamp:  t.UpdatedTimestamp,
	}
}

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
