package workout

import "github.com/google/uuid"

type Workout struct {
	Id                uuid.UUID  `json:"id"`
	TemplateId        uuid.UUID  `json:"templateId"`
	Name              string     `json:"name"`
	AchievedTimestamp string     `json:"achievedTimestamp"`
	Exercises         []Exercise `json:"exercises"`
}

type Exercise struct {
	Name              string        `json:"name"`
	Note              string        `json:"note"`
	Tempo             Tempo         `json:"tempo"`
	TargetRestSeconds int           `json:"restSeconds"`
	Sets              []TemplateSet `json:"sets"`
}

type Set struct {
	Target            TargetSet   `json:"target"`
	Achieved          AchievedSet `json:"achieved"`
	AchievedTimestamp string      `json:"achieved_timestamp"`
}
