package spleen

// SplenicFactor represents a signaling molecule related to splenic function affected by exercise
type SplenicFactor struct {
	Name               string
	Type               string // "Cytokine", "Chemokine", "Growth factor", etc.
	SourceCells        []string
	TargetCells        []string
	ExerciseRegulation string // "Up", "Down", "Biphasic"
	PrimaryEffects     []string
	SplenicFunction    string // How it affects splenic function
	SystemicEffect     string // Effect beyond the spleen
	ConcentrationRange struct {
		Baseline struct {
			Min  float64
			Max  float64
			Unit string
		}
		PostExercise struct {
			Min  float64
			Max  float64
			Unit string
		}
	}
}

// Key splenic factors affected by exercise
var (
	CCL19 = SplenicFactor{
		Name:               "CCL19 (C-C motif chemokine ligand 19)",
		Type:               "Chemokine",
		SourceCells:        []string{"T-zone fibroblastic reticular cells", "Dendritic cells"},
		TargetCells:        []string{"T cells", "Dendritic cells", "B cells"},
		ExerciseRegulation: "Up", // Increased with intense exercise
		PrimaryEffects:     []string{"T cell homing", "Dendritic cell migration", "White pulp organization"},
		SplenicFunction:    "Organizes T cell zones and facilitates T cell-DC interactions",
		SystemicEffect:     "Regulates redistribution of lymphocytes during and after exercise",
		ConcentrationRange: struct {
			Baseline struct {
				Min  float64
				Max  float64
				Unit string
			}
			PostExercise struct {
				Min  float64
				Max  float64
				Unit string
			}
		}{
			Baseline: struct {
				Min  float64
				Max  float64
				Unit string
			}{
				Min:  50.0,
				Max:  150.0,
				Unit: "pg/mL",
			},
			PostExercise: struct {
				Min  float64
				Max  float64
				Unit string
			}{
				Min:  70.0,
				Max:  200.0,
				Unit: "pg/mL",
			},
		},
	}

	TNFSplenic = SplenicFactor{
		Name:               "Tumor Necrosis Factor-α (Splenic)",
		Type:               "Cytokine",
		SourceCells:        []string{"Macrophages", "Dendritic cells", "T cells"},
		TargetCells:        []string{"Widespread immune cells", "Stromal cells"},
		ExerciseRegulation: "Biphasic", // Initial increase, then decrease with training
		PrimaryEffects:     []string{"Inflammation regulation", "Cellular activation", "Marginal zone organization"},
		SplenicFunction:    "Activates macrophages and facilitates germinal center formation",
		SystemicEffect:     "Contributes to post-exercise inflammatory response",
		ConcentrationRange: struct {
			Baseline struct {
				Min  float64
				Max  float64
				Unit string
			}
			PostExercise struct {
				Min  float64
				Max  float64
				Unit string
			}
		}{
			Baseline: struct {
				Min  float64
				Max  float64
				Unit string
			}{
				Min:  2.0,
				Max:  10.0,
				Unit: "pg/mg tissue",
			},
			PostExercise: struct {
				Min  float64
				Max  float64
				Unit string
			}{
				Min:  3.0,
				Max:  20.0,
				Unit: "pg/mg tissue",
			},
		},
	}

	BAFF = SplenicFactor{
		Name:               "B-cell Activating Factor (BAFF)",
		Type:               "Cytokine",
		SourceCells:        []string{"Dendritic cells", "Macrophages", "Neutrophils"},
		TargetCells:        []string{"B cells", "Plasma cells"},
		ExerciseRegulation: "Up", // Increased with regular exercise
		PrimaryEffects:     []string{"B cell survival", "B cell maturation", "Antibody production enhancement"},
		SplenicFunction:    "Maintains B cell follicles and promotes antibody responses",
		SystemicEffect:     "Enhances humoral immunity with regular exercise",
		ConcentrationRange: struct {
			Baseline struct {
				Min  float64
				Max  float64
				Unit string
			}
			PostExercise struct {
				Min  float64
				Max  float64
				Unit string
			}
		}{
			Baseline: struct {
				Min  float64
				Max  float64
				Unit string
			}{
				Min:  0.5,
				Max:  1.5,
				Unit: "ng/mL",
			},
			PostExercise: struct {
				Min  float64
				Max  float64
				Unit string
			}{
				Min:  0.7,
				Max:  2.0,
				Unit: "ng/mL",
			},
		},
	}
)

// CalculateContractileResponse estimates changes in spleen contraction with exercise
func CalculateContractileResponse(exerciseIntensityPercent int,
	durationMinutes int, timePoint string) float64 {

	// Baseline spleen volume (normalized to 1.0)
	volumeFactor := 1.0

	// Exercise parameters
	intensityFactor := float64(exerciseIntensityPercent) / 100.0
	durationFactor := float64(durationMinutes) / 60.0 // Normalized to 1 hour

	// Calculate splenic contraction based on timepoint
	switch timePoint {
	case "During exercise":
		// Spleen contracts during exercise, more with higher intensity
		contractileFactor := 0.3 * intensityFactor
		if durationMinutes > 30 {
			// Additional contraction with longer duration
			contractileFactor += 0.1
		}
		volumeFactor = 1.0 - contractileFactor

	case "Immediate post":
		// Still contracted immediately after exercise
		contractileFactor := 0.2 * intensityFactor
		volumeFactor = 1.0 - contractileFactor

	case "30min post":
		// Begins returning to normal
		contractileFactor := 0.1 * intensityFactor
		volumeFactor = 1.0 - contractileFactor

	case "60min post":
		// Nearly returned to baseline
		contractileFactor := 0.05 * intensityFactor
		volumeFactor = 1.0 - contractileFactor

	default: // "Baseline" or any other state
		volumeFactor = 1.0
	}

	return volumeFactor
}

// CalculateLeukocyteRedistribution estimates how exercise-induced splenic
// contraction affects circulating immune cells
func CalculateLeukocyteRedistribution(spleenVolumeFactor float64) map[string]float64 {

	// Results map for different cell populations in circulation
	redistribution := make(map[string]float64)

	// Calculate redistribution based on spleen contraction
	// (1.0 = baseline levels, >1.0 = increased in circulation)

	// Degree of contraction (0 = no contraction, 0.3 = 30% contraction)
	contractionDegree := 1.0 - spleenVolumeFactor

	// Different cell types mobilized from spleen at different rates
	redistribution["NK cells"] = 1.0 + (4.0 * contractionDegree)    // Strong mobilization
	redistribution["Monocytes"] = 1.0 + (2.0 * contractionDegree)   // Moderate mobilization
	redistribution["B cells"] = 1.0 + (1.0 * contractionDegree)     // Mild mobilization
	redistribution["T cells"] = 1.0 + (1.5 * contractionDegree)     // Moderate mobilization
	redistribution["Neutrophils"] = 1.0 + (1.0 * contractionDegree) // Mild mobilization

	return redistribution
}

// PredictSplenicFactorResponse estimates changes in splenic factors with exercise
func PredictSplenicFactorResponse(factor SplenicFactor,
	exerciseIntensityPercent int, chronicTrainingWeeks int) float64 {

	// Base response factor (1.0 = no change)
	responseFactor := 1.0

	// Exercise parameters
	intensityFactor := float64(exerciseIntensityPercent) / 100.0

	// Training adaptation factor
	adaptationFactor := 0.0
	if chronicTrainingWeeks > 0 {
		adaptationFactor = float64(chronicTrainingWeeks) / 12.0 // Normalized to 12 weeks
		if adaptationFactor > 1.0 {
			adaptationFactor = 1.0 + (0.2 * (adaptationFactor - 1.0)) // Diminishing returns
		}
	}

	// Calculate factor-specific response
	switch factor.Name {
	case "CCL19 (C-C motif chemokine ligand 19)":
		// CCL19 increases with exercise to help with lymphocyte organization
		responseFactor = 1.0 + (0.3 * intensityFactor)
		if chronicTrainingWeeks > 0 {
			responseFactor *= (1.0 + (0.2 * adaptationFactor))
		}

	case "Tumor Necrosis Factor-α (Splenic)":
		// TNF shows biphasic response - acute increase, but lower with chronic training
		acuteResponse := 1.0 + (0.5 * intensityFactor)
		chronicAdaptation := 1.0
		if chronicTrainingWeeks > 0 {
			chronicAdaptation = 1.0 - (0.3 * adaptationFactor)
			if chronicAdaptation < 0.7 {
				chronicAdaptation = 0.7 // Limit the reduction
			}
		}
		responseFactor = acuteResponse * chronicAdaptation

	case "B-cell Activating Factor (BAFF)":
		// BAFF increases more with chronic training than acute exercise
		acuteResponse := 1.0 + (0.2 * intensityFactor)
		chronicAdaptation := 1.0
		if chronicTrainingWeeks > 0 {
			chronicAdaptation = 1.0 + (0.3 * adaptationFactor)
		}
		responseFactor = acuteResponse * chronicAdaptation
	}

	return responseFactor
}
