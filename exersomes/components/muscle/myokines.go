// components/muscle/myokines.go
package muscle

import "../molecular_types"

// MuscleExerkine represents a muscle-derived signaling molecule
type MuscleExerkine struct {
	Name               string
	TargetOrgans       []string
	ResponseToExercise string
	SignalingPathway   string
	PeakTimeMinutes    int
}

// Example muscle exerkines
var (
	IL6 = MuscleExerkine{
		Name:               "Interleukin-6",
		TargetOrgans:       []string{"Liver", "Adipose", "Immune cells"},
		ResponseToExercise: "Increases with glycogen depletion",
		SignalingPathway:   "JAK-STAT",
		PeakTimeMinutes:    30,
	}

	// Add more muscle exerkines
)

var (
	IL6 = molecular_types.Exerkine{
		Name:           "Interleukin-6",
		Category:       "Protein",
		TissueSources:  []string{"Skeletal muscle", "Immune cells"},
		BiologicalFunc: "Immune response, inflammation, metabolism",
		Sequence:       "", // Would be full AA sequence
		PDBID:          "1ALU",
		Receptors:      []string{"IL6R", "IL6ST"},
	}

	Irisin = molecular_types.Exerkine{
		Name:           "Irisin (FNDC5)",
		Category:       "Protein",
		TissueSources:  []string{"Skeletal muscle"},
		BiologicalFunc: "Browning of white adipose tissue",
		Sequence:       "", // Full AA sequence
		PDBID:          "4LSD",
		Receptors:      []string{"Unknown"},
	}
)

// Additional muscle-specific functions
func GetMuscleExerkines() []molecular_types.Exerkine {
	return []molecular_types.Exerkine{IL6, Irisin /*, ... */}
}
