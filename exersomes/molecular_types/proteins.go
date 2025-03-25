package proteins

// ExerciseProtein represents a protein released during exercise
type ExerciseProtein struct {
	Name                 string
	UniprotID            string
	MolecularWeightKDa   float64
	Classification       string   // "Myokine", "Hepatokine", "Adipokine", etc.
	SourceTissues        []string // Tissues that produce this protein
	Receptors            []string // Receptor(s) that bind this protein
	ExerciseRegulation   string   // "Acute up", "Chronic up", "Down", "Biphasic"
	TimeToRelease        string   // When released during/after exercise
	CirculationHalfLife  string   // Half-life in circulation
	SignalingPathways    []string // Pathways activated by this protein
	TargetTissues        []string // Tissues affected by this protein
	PhysiologicalEffects []string // Effects when increased by exercise
	ClinicalSignificance []string // Related disease contexts
	ExerciseTypes        []string // Exercise types with strongest evidence
}

// Key exercise-responsive proteins
var (
	IL6 = ExerciseProtein{
		Name:                "Interleukin 6",
		UniprotID:           "P05231",
		MolecularWeightKDa:  21.0,
		Classification:      "Myokine/Cytokine",
		SourceTissues:       []string{"Skeletal muscle", "Immune cells", "Adipose tissue"},
		Receptors:           []string{"IL-6R", "gp130"},
		ExerciseRegulation:  "Acute up, Chronic down",
		TimeToRelease:       "Increases during exercise, peaks immediately post-exercise",
		CirculationHalfLife: "1-2 hours",
		SignalingPathways: []string{
			"JAK/STAT",
			"MAPK/ERK",
			"PI3K/Akt",
		},
		TargetTissues: []string{
			"Liver",
			"Adipose tissue",
			"Skeletal muscle",
			"Pancreatic β-cells",
		},
		PhysiologicalEffects: []string{
			"Increases glucose uptake",
			"Enhances fat oxidation",
			"Stimulates lipolysis",
			"Improves insulin sensitivity",
		},
		ClinicalSignificance: []string{
			"Insulin resistance",
			"Obesity",
			"Inflammation",
		},
		ExerciseTypes: []string{"Endurance", "HIIT", "Prolonged aerobic"},
	}

	// Add other key proteins: BDNF, Irisin, IL-15, Decorin, Myostatin, etc.
)

// CalculateProteinResponse estimates changes in protein levels with different exercise protocols
func CalculateProteinResponse(protein ExerciseProtein, exerciseType string,
	intensity float64, duration int, chronicWeeks int) float64 {
	// Implementation for predicting protein response based on exercise parameters
}

// Implementation for receptor binding, tissue effects, etc.
