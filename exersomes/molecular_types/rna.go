package molecular_types

// ExerciseRNA represents an RNA affected by exercise (excluding miRNAs)
type ExerciseRNA struct {
	Name               string
	Type               string   // "lncRNA", "circRNA", "mRNA", "tRNA", etc.
	Source             string   // Primary tissue source
	Length             int      // Nucleotide length
	ExerciseRegulation string   // "Up", "Down", "Complex"
	TransportMechanism []string // How it's transported in circulation
	TargetTissues      []string // Tissues that take up this RNA
	Function           []string // Functional roles
	AssociatedPathways []string // Pathways it affects
	RelatedDiseases    []string // Disease contexts
	ExerciseTiming     string   // When changes occur during/after exercise
}

// Key exercise-responsive RNAs
var (
	IL6_RNA = ExerciseRNA{
		Name:               "IL6",
		Type:               "pcRNA",
		Source:             "Multiple tissues",
		Length:             8708,
		ExerciseRegulation: "Up with endurance exercise",
		TransportMechanism: []string{"Extracellular vesicles", "Protein complexes"},
		TargetTissues:      []string{"Endothelial cells", "Muscle", "Adipose"},
		Function: []string{
			"Modulates immune and inflammatory responses",
			"Regulates metabolic adaptation to exercise",
			"Promotes myogenesis and muscle hypertrophy",
		},
		AssociatedPathways: []string{
			"JAK-STAT signaling pathway",
			"Cytokine-cytokine receptor interaction",
			"PI3K-Akt signaling pathway",
		},
		RelatedDiseases: []string{
			"Infectious disease",
			"Autoimmune disease",
			"Inflammatory bowel disease",
		},
		ExerciseTiming: "Gradual increase, peaks several hours post-exercise",
	}

	// Add other key RNAs: circular RNAs, other lncRNAs, etc.
)

// PredictRNAResponse estimates RNA changes with exercise
func PredictRNAResponse(rna ExerciseRNA, exerciseType string,
	intensity float64, duration int) float64 {
	// Implementation for predicting RNA levels post-exercise
	return 0
}
