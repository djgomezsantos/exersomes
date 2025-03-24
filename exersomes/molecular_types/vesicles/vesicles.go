package vesicles

// ExerciseVesicle represents extracellular vesicles affected by exercise
type ExerciseVesicle struct {
	Type               string   // "Exosome", "Microvesicle", "Apoptotic body", etc.
	SizeRange          string   // Size range in nm
	PrimarySource      string   // Main tissue source
	SecondarySource    []string // Other tissue sources
	ExerciseRegulation string   // How exercise affects their release
	TemporalDynamics   string   // Time course of appearance in circulation
	MarkerProteins     []string // Characteristic surface markers
	CargoTypes         []string // Types of cargo ("miRNA", "protein", "lipids", etc.)
	KeyCargo           []string // Specific important cargo molecules
	TargetTissues      []string // Tissues that take up these vesicles
	FunctionalEffects  []string // Physiological effects
	BestInducers       []string // Exercise types that best induce these vesicles
}

// Key exercise-responsive vesicles
var (
	MuscleDerivedExosomes = ExerciseVesicle{
		Type:               "Exosome",
		SizeRange:          "30-150 nm",
		PrimarySource:      "Skeletal muscle",
		SecondarySource:    []string{},
		ExerciseRegulation: "Up-regulated acutely with both aerobic and resistance exercise",
		TemporalDynamics:   "Peak 30min-2hr post-exercise, return to baseline by 24h",
		MarkerProteins: []string{
			"CD9",
			"CD63",
			"CD81",
			"HSP70",
			"Muscle-specific markers (e.g., MSTN, MyoD)",
		},
		CargoTypes: []string{
			"miRNAs",
			"mRNAs",
			"Proteins",
			"Metabolites",
		},
		KeyCargo: []string{
			"miR-1",
			"miR-133a",
			"miR-206",
			"miR-486",
			"PGC-1α mRNA",
			"Myostatin protein",
		},
		TargetTissues: []string{
			"Liver",
			"Adipose tissue",
			"Endothelial cells",
			"Brain",
		},
		FunctionalEffects: []string{
			"Improved glucose metabolism in recipient tissues",
			"Enhanced angiogenesis",
			"Mitochondrial adaptation",
			"Anti-inflammatory effects",
		},
		BestInducers: []string{
			"HIIT",
			"Prolonged endurance exercise",
			"Resistance exercise",
		},
	}

	// Add other vesicle types
)

// PredictVesicleResponse estimates vesicle changes with exercise
func PredictVesicleResponse(vesicle ExerciseVesicle, exerciseType string,
	intensity float64, duration int) float64 {
	// Implementation for predicting vesicle release based on exercise parameters
}

// Functions for cargo loading, vesicle biodistribution, etc.
