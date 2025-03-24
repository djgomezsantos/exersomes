// components/musculoskeletal/exerkines.go
package musculoskeletal

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
