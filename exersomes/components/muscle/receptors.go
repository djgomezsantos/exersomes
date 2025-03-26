// components/musculoskeletal/prescription.go
package muscle

// ExercisePrescription defines exercise parameters for musculoskeletal health
type ExercisePrescription struct {
	IntensityPercent1RM      int    // Percentage of 1-rep max
	SetsCount                int    // Number of sets
	RepsPerSet               int    // Reps per set
	RestBetweenSetsSec       int    // Rest between sets in seconds
	FrequencyPerWeek         int    // Sessions per week
	ExerciseType             string // e.g., "Resistance", "Plyometric"
	TargetMuscleGroups       []string
	ExpectedExerkineResponse map[string]float64 // Expected exerkine changes
}

// Predefined prescriptions
var (
	HypertrophyProtocol = ExercisePrescription{
		IntensityPercent1RM: 75,
		SetsCount:           4,
		RepsPerSet:          8,
		RestBetweenSetsSec:  90,
		FrequencyPerWeek:    3,
		ExerciseType:        "Resistance",
		TargetMuscleGroups:  []string{"Quadriceps", "Hamstrings", "Gluteus"},
		ExpectedExerkineResponse: map[string]float64{
			"IL-6":  4.5, // 4.5-fold increase
			"IL-15": 2.3, // 2.3-fold increase
		},
	}

	// Add more prescription protocols
)
