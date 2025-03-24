package adipose

// Adipokine represents a signaling molecule secreted from adipose tissue
type Adipokine struct {
	Name                 string
	AdiposeFraction      string  // "White", "Brown", "Beige", or "All"
	MolecularWeight      float64 // In kDa
	TargetOrgans         []string
	ExerciseRegulation   string // "Up", "Down", "Biphasic"
	ExerciseResponseTime string // "Acute", "Chronic", "Both"
	PrimaryEffects       []string
	MetabolicAction      string // "Insulin sensitizing", "Pro-inflammatory", etc.
	BaselineRange        struct {
		Min  float64
		Max  float64
		Unit string
	}
}

// Key adipokines affected by exercise
var (
	Adiponectin = Adipokine{
		Name:                 "Adiponectin",
		AdiposeFraction:      "All",
		MolecularWeight:      30.0,
		TargetOrgans:         []string{"Muscle", "Liver", "Brain", "Cardiovascular"},
		ExerciseRegulation:   "Up",
		ExerciseResponseTime: "Chronic",
		PrimaryEffects:       []string{"Insulin sensitizing", "Anti-inflammatory", "Anti-atherogenic"},
		MetabolicAction:      "Insulin sensitizing",
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  5.0,
			Max:  30.0,
			Unit: "μg/mL",
		},
	}

	Leptin = Adipokine{
		Name:                 "Leptin",
		AdiposeFraction:      "White",
		MolecularWeight:      16.0,
		TargetOrgans:         []string{"Brain", "Muscle", "Liver", "Pancreas"},
		ExerciseRegulation:   "Down",
		ExerciseResponseTime: "Both",
		PrimaryEffects:       []string{"Appetite regulation", "Energy expenditure", "Immune modulation"},
		MetabolicAction:      "Energy homeostasis",
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  1.0,
			Max:  50.0,
			Unit: "ng/mL",
		},
	}

	IL6Adipose = Adipokine{
		Name:                 "IL-6",
		AdiposeFraction:      "White",
		MolecularWeight:      21.0,
		TargetOrgans:         []string{"Liver", "Muscle", "Immune cells"},
		ExerciseRegulation:   "Biphasic", // Acute increase, chronic decrease
		ExerciseResponseTime: "Both",
		PrimaryEffects:       []string{"Lipolysis", "Insulin signaling modulation", "Inflammation"},
		MetabolicAction:      "Context-dependent",
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  1.0,
			Max:  10.0,
			Unit: "pg/mL",
		},
	}

	TNFalpha = Adipokine{
		Name:                 "TNF-α",
		AdiposeFraction:      "White",
		MolecularWeight:      17.0,
		TargetOrgans:         []string{"Adipose", "Muscle", "Liver"},
		ExerciseRegulation:   "Down",
		ExerciseResponseTime: "Chronic",
		PrimaryEffects:       []string{"Insulin resistance", "Lipolysis", "Inflammation"},
		MetabolicAction:      "Pro-inflammatory",
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  0.5,
			Max:  2.0,
			Unit: "pg/mL",
		},
	}

	Irisin = Adipokine{
		Name:                 "Irisin",
		AdiposeFraction:      "Beige", // Primarily muscle-derived but affects adipose
		MolecularWeight:      12.0,
		TargetOrgans:         []string{"Adipose", "Bone", "Brain"},
		ExerciseRegulation:   "Up",
		ExerciseResponseTime: "Acute",
		PrimaryEffects:       []string{"Browning of white adipose", "Thermogenesis", "Bone formation"},
		MetabolicAction:      "Energy expenditure",
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  3.0,
			Max:  10.0,
			Unit: "ng/mL",
		},
	}

	FABP4 = Adipokine{
		Name:                 "FABP4",
		AdiposeFraction:      "White",
		MolecularWeight:      15.0,
		TargetOrgans:         []string{"Liver", "Macrophages", "Endothelium"},
		ExerciseRegulation:   "Down",
		ExerciseResponseTime: "Chronic",
		PrimaryEffects:       []string{"Fatty acid transport", "Insulin signaling", "Inflammation"},
		MetabolicAction:      "Insulin resistance",
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  10.0,
			Max:  50.0,
			Unit: "ng/mL",
		},
	}
)

// GetExerciseUpregulatedAdipokines returns adipokines that increase with exercise
func GetExerciseUpregulatedAdipokines() []Adipokine {
	return []Adipokine{Adiponectin, Irisin}
}

// GetExerciseDownregulatedAdipokines returns adipokines that decrease with exercise
func GetExerciseDownregulatedAdipokines() []Adipokine {
	return []Adipokine{Leptin, TNFalpha, FABP4}
}

// PredictAdipokineLevels calculates estimated adipokine changes based on exercise and body composition
func PredictAdipokineLevels(exerciseType string, intensityPercent int,
	durationWeeks int, bodyFatPercent float64) map[string]float64 {

	// Basic implementation - would be more sophisticated in practice
	levels := make(map[string]float64)

	// Set baseline levels based on body fat percent
	levels["Adiponectin"] = 15.0 - (bodyFatPercent * 0.3) // Inverse relationship with body fat
	levels["Leptin"] = 5.0 + (bodyFatPercent * 0.9)       // Direct relationship with body fat
	levels["TNF-α"] = 0.8 + (bodyFatPercent * 0.05)       // Direct relationship with body fat
	levels["Irisin"] = 5.0                                // Less affected by initial body fat
	levels["FABP4"] = 20.0 + (bodyFatPercent * 0.6)       // Direct relationship with body fat

	// Modify based on exercise
	if exerciseType == "Aerobic" || exerciseType == "HIIT" {
		// Effect increases with intensity, duration
		intensityFactor := float64(intensityPercent) / 100.0
		durationFactor := float64(durationWeeks) / 12.0 // Normalized to 12 weeks
		combinedFactor := intensityFactor * durationFactor

		// Upregulated adipokines
		levels["Adiponectin"] *= (1.0 + 0.3*combinedFactor)
		levels["Irisin"] *= (1.0 + 0.5*combinedFactor)

		// Downregulated adipokines
		levels["Leptin"] *= (1.0 - 0.2*combinedFactor)
		levels["TNF-α"] *= (1.0 - 0.3*combinedFactor)
		levels["FABP4"] *= (1.0 - 0.2*combinedFactor)
	}

	if exerciseType == "Resistance" {
		// Resistance has different effects
		durationFactor := float64(durationWeeks) / 12.0

		levels["Adiponectin"] *= (1.0 + 0.2*durationFactor)
		levels["Irisin"] *= (1.0 + 0.3*durationFactor)
		levels["Leptin"] *= (1.0 - 0.15*durationFactor)
		levels["TNF-α"] *= (1.0 - 0.25*durationFactor)
		levels["FABP4"] *= (1.0 - 0.15*durationFactor)
	}

	return levels
}
