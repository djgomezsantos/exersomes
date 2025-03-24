package thymus

// ThymicFactor represents a signaling molecule related to thymus function affected by exercise
type ThymicFactor struct {
	Name               string
	Type               string // "Hormone", "Cytokine", "Chemokine", etc.
	SourceCells        []string
	TargetCells        []string
	ExerciseRegulation string // "Up", "Down", "Biphasic"
	PrimaryEffects     []string
	ThymicFunction     string // How it affects thymus function
	AgingEffect        string // How it relates to thymic involution/aging
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

// Key thymic factors affected by exercise
var (
	Thymulin = ThymicFactor{
		Name:               "Thymulin",
		Type:               "Hormone",
		SourceCells:        []string{"Thymic epithelial cells"},
		TargetCells:        []string{"T cell precursors", "Mature T cells"},
		ExerciseRegulation: "Up", // Moderate increase with regular exercise
		PrimaryEffects:     []string{"T cell differentiation", "T cell function enhancement", "Anti-inflammatory"},
		ThymicFunction:     "Promotes thymocyte maturation and T cell function",
		AgingEffect:        "Declines with age; exercise may partially preserve levels",
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
				Min:  10.0,
				Max:  50.0,
				Unit: "fg/mL",
			},
			PostExercise: struct {
				Min  float64
				Max  float64
				Unit string
			}{
				Min:  15.0,
				Max:  70.0,
				Unit: "fg/mL",
			},
		},
	}

	IL7Thymic = ThymicFactor{
		Name:               "Interleukin-7 (Thymic)",
		Type:               "Cytokine",
		SourceCells:        []string{"Thymic epithelial cells", "Bone marrow stromal cells"},
		TargetCells:        []string{"Thymocytes", "Naive T cells", "B cell progenitors"},
		ExerciseRegulation: "Up", // Increased with regular moderate exercise
		PrimaryEffects:     []string{"T cell development", "Thymic cellularity maintenance", "Lymphocyte survival"},
		ThymicFunction:     "Essential for thymocyte development and survival",
		AgingEffect:        "Reduced with age; contributes to thymic involution",
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
				Max:  8.0,
				Unit: "pg/mL",
			},
			PostExercise: struct {
				Min  float64
				Max  float64
				Unit string
			}{
				Min:  3.0,
				Max:  12.0,
				Unit: "pg/mL",
			},
		},
	}

	KGF = ThymicFactor{
		Name:               "Keratinocyte Growth Factor (FGF-7)",
		Type:               "Growth factor",
		SourceCells:        []string{"Mesenchymal cells", "Fibroblasts"},
		TargetCells:        []string{"Thymic epithelial cells"},
		ExerciseRegulation: "Up", // Modest increase with regular exercise
		PrimaryEffects:     []string{"Thymic epithelial cell proliferation", "Thymic architecture maintenance"},
		ThymicFunction:     "Supports thymic epithelial cell health and function",
		AgingEffect:        "Reduced with age; exercise may slow decline",
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
				Min:  5.0,
				Max:  20.0,
				Unit: "pg/mL",
			},
			PostExercise: struct {
				Min  float64
				Max  float64
				Unit string
			}{
				Min:  8.0,
				Max:  25.0,
				Unit: "pg/mL",
			},
		},
	}
)

// CalculateThymicResponse estimates changes in thymic factors with exercise
func CalculateThymicResponse(factor ThymicFactor,
	exerciseIntensityPercent int, ageYears int, chronicTrainingWeeks int) float64 {

	// Base response factor (1.0 = no change)
	responseFactor := 1.0

	// Exercise parameters
	intensityFactor := float64(exerciseIntensityPercent) / 100.0

	// Optimal intensity for thymic health is moderate (around 60-70%)
	optimalityFactor := 1.0
	if intensityFactor > 0.8 {
		// High intensity can be stressful to thymus
		optimalityFactor = 1.0 - (0.5 * (intensityFactor - 0.8) / 0.2)
	} else if intensityFactor < 0.5 {
		// Low intensity has less effect
		optimalityFactor = 0.8 + (0.2 * intensityFactor / 0.5)
	} else {
		// Optimal range
		optimalityFactor = 1.0
	}

	// Age effect (thymic factors decline with age)
	ageFactor := 1.0
	if ageYears > 20 {
		// Simplified model of thymic involution
		ageFactor = 1.0 - (0.5 * float64(ageYears-20) / 60.0)
		if ageFactor < 0.3 {
			ageFactor = 0.3 // Floor at 30% of young adult levels
		}
	}

	// Training adaptation factor
	adaptationFactor := 1.0
	if chronicTrainingWeeks > 0 {
		adaptationWeeks := float64(chronicTrainingWeeks) / 12.0 // Normalized to 12 weeks
		if adaptationWeeks > 1.0 {
			adaptationWeeks = 1.0 + (0.2 * (adaptationWeeks - 1.0)) // Diminishing returns
		}
		// Regular exercise can partially counteract age-related decline
		adaptationFactor = 1.0 + (0.2 * adaptationWeeks)
	}

	// Calculate factor-specific response
	switch factor.Name {
	case "Thymulin":
		// Thymulin is quite responsive to exercise training
		responseFactor = ageFactor * (1.0 + (0.3 * optimalityFactor * adaptationFactor))

	case "Interleukin-7 (Thymic)":
		// IL-7 response depends on both acute and chronic exercise
		acuteEffect := 0.1 * intensityFactor * optimalityFactor
		chronicEffect := 0.0
		if chronicTrainingWeeks > 0 {
			chronicEffect = 0.2 * adaptationFactor
		}
		responseFactor = ageFactor * (1.0 + acuteEffect + chronicEffect)

	case "Keratinocyte Growth Factor (FGF-7)":
		// KGF responds more to chronic training than acute exercise
		acuteEffect := 0.05 * intensityFactor * optimalityFactor
		chronicEffect := 0.0
		if chronicTrainingWeeks > 0 {
			chronicEffect = 0.15 * adaptationFactor
		}
		responseFactor = ageFactor * (1.0 + acuteEffect + chronicEffect)
	}

	return responseFactor
}

// PredictThymicSize estimates change in thymic size with exercise training
func PredictThymicSize(ageYears int, chronicTrainingWeeks int,
	intensityPercent int, consistencyPercent int) float64 {

	// Base size factor (1.0 = normal for age)
	sizeFactor := 1.0

	// Calculate baseline age effect on thymus size (thymic involution)
	if ageYears <= 20 {
		sizeFactor = 1.0 // Young thymus
	} else if ageYears <= 40 {
		sizeFactor = 1.0 - (0.3 * float64(ageYears-20) / 20.0) // Gradual decline
	} else if ageYears <= 60 {
		sizeFactor = 0.7 - (0.2 * float64(ageYears-40) / 20.0) // Continued decline
	} else {
		sizeFactor = 0.5 - (0.2 * float64(ageYears-60) / 20.0) // Further decline
		if sizeFactor < 0.2 {
			sizeFactor = 0.2 // Floor at 20% of young adult size
		}
	}

	// Exercise effect (can partially counteract age-related decline)
	if chronicTrainingWeeks > 0 {
		// Calculate training effect
		trainingWeeks := float64(chronicTrainingWeeks)
		intensityEffect := float64(intensityPercent) / 70.0 // Normalized to moderate intensity
		consistencyEffect := float64(consistencyPercent) / 100.0

		// Optimal intensity is moderate; too high can be counterproductive
		if intensityEffect > 1.2 {
			intensityEffect = 1.2 - (0.5 * (intensityEffect - 1.2))
		}

		// Calculate exercise benefit (diminishing returns over time)
		if trainingWeeks <= 12 {
			exerciseBenefit := 0.05 * (trainingWeeks / 12.0) * intensityEffect * consistencyEffect
			sizeFactor *= (1.0 + exerciseBenefit)
		} else {
			exerciseBenefit := 0.05 + (0.05 * (trainingWeeks - 12.0) / 48.0)
			if exerciseBenefit > 0.15 {
				exerciseBenefit = 0.15 // Cap at 15% improvement
			}
			sizeFactor *= (1.0 + exerciseBenefit*intensityEffect*consistencyEffect)
		}
	}

	return sizeFactor
}

// PredictTCellDiversity estimates T cell receptor diversity as affected by exercise
func PredictTCellDiversity(ageYears int, exerciseYears int,
	daysPerWeek int) float64 {

	// Base diversity factor (1.0 = young adult diversity)
	diversityFactor := 1.0

	// Age-related decline in TCR diversity
	if ageYears <= 30 {
		diversityFactor = 1.0 // Young adult diversity
	} else if ageYears <= 60 {
		diversityFactor = 1.0 - (0.3 * float64(ageYears-30) / 30.0) // Gradual decline
	} else {
		diversityFactor = 0.7 - (0.3 * float64(ageYears-60) / 30.0) // Accelerated decline
		if diversityFactor < 0.3 {
			diversityFactor = 0.3 // Floor at 30% of young adult diversity
		}
	}

	// Exercise effect (can help maintain diversity)
	if exerciseYears > 0 {
		// Calculate long-term exercise effect
		frequencyFactor := float64(daysPerWeek) / 3.0 // Normalized to 3 days/week
		if frequencyFactor > 1.5 {
			frequencyFactor = 1.5 // Cap frequency benefit
		}

		durationEffect := float64(exerciseYears)
		if durationEffect > 10 {
			durationEffect = 10.0 + (0.2 * (durationEffect - 10.0)) // Diminishing returns
		}

		// Calculate diversity preservation (up to 25% improvement)
		preservationFactor := 0.25 * (durationEffect / 10.0) * frequencyFactor

		// Apply exercise benefit
		diversityFactor *= (1.0 + preservationFactor)

		// Cap at 90% of young adult diversity for older adults
		if ageYears > 60 && diversityFactor > 0.9 {
			diversityFactor = 0.9
		}
	}

	return diversityFactor
}
