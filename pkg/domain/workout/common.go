package workout

type RepRange struct {
	Min   int  `json:"min"`
	Max   int  `json:"max"`
	Amrap bool `json:"amrap"`
}

type Weight struct {
	Amount float64    `json:"amount"`
	Unit   WeightUnit `json:"unit"`
}

type WeightUnit string

const (
	Kg WeightUnit = "kg"
	Lb WeightUnit = "lb"
	Na WeightUnit = "na"
)

type Load struct {
	LoadingScheme LoadingScheme `json:"scheme"`
	Pct           float64       `json:"pct"`
	Rpe           float64       `json:"rpe"`
	Absolute      Weight        `json:"absolute"`
}

type LoadingScheme string

const (
	PercentOfMax LoadingScheme = "pct"
	RPE          LoadingScheme = "rpe"
	Absolute     LoadingScheme = "abs"
)

type TargetSet struct {
	Reps RepRange `json:"reps"`
	Load Load     `json:"load"`
}

type AchievedSet struct {
	Reps   int    `json:"reps"`
	Weight Weight `json:"weight"`
}

type Tempo struct {
	EccentricSeconds  int `json:"eccentricSeconds"`
	PauseSeconds      int `json:"pauseSeconds"`
	ConcentricSeconds int `json:"concentricSeconds"`
	RestSeconds       int `json:"restSeconds"`
}
