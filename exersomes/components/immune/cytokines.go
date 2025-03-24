package immune

// Cytokine represents an immune signaling molecule affected by exercise
type Cytokine struct {
	Name               string
	Family             string   // Cytokine family
	MolecularWeight    float64  // In kDa
	SourceCells        []string // Cell types producing this factor
	TargetCells        []string
	AcuteRegulation    string   // "Up", "Down", "Biphasic", "No change"
	ChronicRegulation  string   // Long-term adaptation with regular exercise
	HalfLifeMinutes    float64  // Circulation half-life
	SystemicEffects    []string // Effects in circulation
	LocalEffects       []string // Effects at tissue level
	ExerciseThreshold  string   // Exercise intensity needed for significant changes
	RecoveryTimeframe  string   // Time until return to baseline after exercise
	PrimaryFunctions   []string // Main physiological functions
	ConcentrationRange struct {
		Baseline struct {
			Min  float64
			Max  float64
			Unit string
		}
		PostExerciseAcute struct {
			Min      float64
			Max      float64
			Unit     string
			TimeHour float64 // Hours post-exercise
		}
		TrainedBaseline struct {
			Min  float64
			Max  float64
			Unit string
		}
	}
}

// Key cytokines affected by exercise
var (
	IL6 = Cytokine{
		Name:              "Interleukin-6",
		Family:            "Interleukin",
		MolecularWeight:   21.0,
		SourceCells:       []string{"Myocytes", "Macrophages", "T cells", "Endothelial cells", "Fibroblasts"},
		TargetCells:       []string{"Hepatocytes", "Immune cells", "Adipocytes", "Brain cells", "Muscle cells"},
		AcuteRegulation:   "Up",
		ChronicRegulation: "Down", // Reduced baseline and exercise-induced response in trained individuals
		HalfLifeMinutes:   45.0,
		SystemicEffects:   []string{"Acute phase response", "Glucose metabolism", "Lipolysis", "HPA axis stimulation"},
		LocalEffects:      []string{"Satellite cell proliferation", "Muscle hypertrophy", "Fat oxidation", "Insulin sensitization"},
		ExerciseThreshold: "Moderate-to-high intensity, >30min duration",
		RecoveryTimeframe: "Returns to baseline within 6-24 hours",
		PrimaryFunctions:  []string{"Metabolic signaling", "Pro- and anti-inflammatory effects", "Muscle repair", "Exercise adaptation"},
		ConcentrationRange: struct {
			Baseline struct {
				Min  float64
				Max  float64
				Unit string
			}
			PostExerciseAcute struct {
				Min      float64
				Max      float64
				Unit     string
				TimeHour float64
			}
			TrainedBaseline struct {
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
				Min:  1.0,
				Max:  5.0,
				Unit: "pg/mL",
			},
			PostExerciseAcute: struct {
				Min      float64
				Max      float64
				Unit     string
				TimeHour float64
			}{
				Min:      10.0,
				Max:      120.0,
				Unit:     "pg/mL",
				TimeHour: 1.0,
			},
			TrainedBaseline: struct {
				Min  float64
				Max  float64
				Unit string
			}{
				Min:  0.8,
				Max:  3.0,
				Unit: "pg/mL",
			},
		},
	}

	TNF = Cytokine{
		Name:              "Tumor Necrosis Factor-α",
		Family:            "TNF family",
		MolecularWeight:   17.0,
		SourceCells:       []string{"Macrophages", "NK cells", "T cells", "Adipocytes"},
		TargetCells:       []string{"Widespread", "Immune cells", "Adipocytes", "Endothelial cells", "Muscle cells"},
		AcuteRegulation:   "Up",   // Modest increase with exercise
		ChronicRegulation: "Down", // Reduced baseline in trained individuals
		HalfLifeMinutes:   20.0,
		SystemicEffects:   []string{"Pro-inflammatory signaling", "Insulin resistance", "Endothelial activation"},
		LocalEffects:      []string{"Macrophage activation", "Cell death regulation", "Muscle catabolism"},
		ExerciseThreshold: "High intensity or prolonged exercise",
		RecoveryTimeframe: "Returns to baseline within 1-3 hours",
		PrimaryFunctions:  []string{"Inflammatory response", "Host defense", "Tissue remodeling", "Metabolic regulation"},
		ConcentrationRange: struct {
			Baseline struct {
				Min  float64
				Max  float64
				Unit string
			}
			PostExerciseAcute struct {
				Min      float64
				Max      float64
				Unit     string
				TimeHour float64
			}
			TrainedBaseline struct {
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
				Min:  1.0,
				Max:  10.0,
				Unit: "pg/mL",
			},
			PostExerciseAcute: struct {
				Min      float64
				Max      float64
				Unit     string
				TimeHour float64
			}{
				Min:      5.0,
				Max:      25.0,
				Unit:     "pg/mL",
				TimeHour: 0.5,
			},
			TrainedBaseline: struct {
				Min  float64
				Max  float64
				Unit string
			}{
				Min:  0.8,
				Max:  7.0,
				Unit: "pg/mL",
			},
		},
	}

	IL10 = Cytokine{
		Name:              "Interleukin-10",
		Family:            "Interleukin",
		MolecularWeight:   18.5,
		SourceCells:       []string{"Macrophages", "Regulatory T cells", "B cells", "Monocytes"},
		TargetCells:       []string{"Macrophages", "Dendritic cells", "T cells", "B cells", "NK cells"},
		AcuteRegulation:   "Up",
		ChronicRegulation: "Up", // Enhanced anti-inflammatory response with training
		HalfLifeMinutes:   120.0,
		SystemicEffects:   []string{"Anti-inflammatory signaling", "Immune tolerance", "Macrophage deactivation"},
		LocalEffects:      []string{"Tissue repair promotion", "Inflammation resolution", "T cell regulation"},
		ExerciseThreshold: "Moderate-to-high intensity, >45min duration",
		RecoveryTimeframe: "Peaks 6-24 hours post-exercise",
		PrimaryFunctions:  []string{"Anti-inflammatory response", "Immune regulation", "Tissue homeostasis", "Recovery promotion"},
		ConcentrationRange: struct {
			Baseline struct {
				Min  float64
				Max  float64
				Unit string
			}
			PostExerciseAcute struct {
				Min      float64
				Max      float64
				Unit     string
				TimeHour float64
			}
			TrainedBaseline struct {
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
				Min:  3.0,
				Max:  8.0,
				Unit: "pg/mL",
			},
			PostExerciseAcute: struct {
				Min      float64
				Max      float64
				Unit     string
				TimeHour float64
			}{
				Min:      8.0,
				Max:      30.0,
				Unit:     "pg/mL",
				TimeHour: 6.0,
			},
			TrainedBaseline: struct {
				Min  float64
				Max  float64
				Unit string
			}{
				Min:  4.0,
				Max:  10.0,
				Unit: "pg/mL",
			},
		},
	}

	IL1RA = Cytokine{
		Name:              "Interleukin-1 Receptor Antagonist",
		Family:            "Interleukin-1 family",
		MolecularWeight:   17.0,
		SourceCells:       []string{"Macrophages", "Monocytes", "Hepatocytes", "Neutrophils", "Myocytes"},
		TargetCells:       []string{"Cells expressing IL-1 receptor", "Immune cells", "Brain cells"},
		AcuteRegulation:   "Up",
		ChronicRegulation: "Up", // Enhanced anti-inflammatory response with training
		HalfLifeMinutes:   180.0,
		SystemicEffects:   []string{"IL-1 signaling inhibition", "Anti-inflammatory action", "Fever reduction"},
		LocalEffects:      []string{"Tissue repair facilitation", "Inflammatory response limitation"},
		ExerciseThreshold: "Moderate intensity, >30min duration",
		RecoveryTimeframe: "Elevated for 24+ hours post-exercise",
		PrimaryFunctions:  []string{"IL-1 antagonism", "Inflammatory regulation", "Exercise recovery", "Fever control"},
		ConcentrationRange: struct {
			Baseline struct {
				Min  float64
				Max  float64
				Unit string
			}
			PostExerciseAcute struct {
				Min      float64
				Max      float64
				Unit     string
				TimeHour float64
			}
			TrainedBaseline struct {
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
				Min:  200.0,
				Max:  500.0,
				Unit: "pg/mL",
			},
			PostExerciseAcute: struct {
				Min      float64
				Max      float64
				Unit     string
				TimeHour float64
			}{
				Min:      500.0,
				Max:      2000.0,
				Unit:     "pg/mL",
				TimeHour: 2.0,
			},
			TrainedBaseline: struct {
				Min  float64
				Max  float64
				Unit string
			}{
				Min:  250.0,
				Max:  600.0,
				Unit: "pg/mL",
			},
		},
	}

	TGFbeta = Cytokine{
		Name:              "Transforming Growth Factor-β",
		Family:            "TGF-β superfamily",
		MolecularWeight:   25.0,
		SourceCells:       []string{"Platelets", "Macrophages", "T cells", "Fibroblasts", "Endothelial cells"},
		TargetCells:       []string{"Fibroblasts", "Immune cells", "Epithelial cells", "Endothelial cells"},
		AcuteRegulation:   "Up",      // Modest increase with exercise
		ChronicRegulation: "Complex", // Tissue-dependent adaptation
		HalfLifeMinutes:   60.0,
		SystemicEffects:   []string{"Immune regulation", "Anti-inflammatory actions", "Tissue remodeling signals"},
		LocalEffects:      []string{"ECM production", "Fibrosis regulation", "Wound healing", "Epithelial-mesenchymal transition"},
		ExerciseThreshold: "Moderate intensity, longer durations favored",
		RecoveryTimeframe: "Sustained elevation for 24-48 hours post-exercise",
		PrimaryFunctions:  []string{"Tissue repair", "Immune tolerance", "Inflammation resolution", "Fibrosis regulation"},
		ConcentrationRange: struct {
			Baseline struct {
				Min  float64
				Max  float64
				Unit string
			}
			PostExerciseAcute struct {
				Min      float64
				Max      float64
				Unit     string
				TimeHour float64
			}
			TrainedBaseline struct {
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
				Max:  5.0,
				Unit: "ng/mL",
			},
			PostExerciseAcute: struct {
				Min      float64
				Max      float64
				Unit     string
				TimeHour float64
			}{
				Min:      3.0,
				Max:      10.0,
				Unit:     "ng/mL",
				TimeHour: 24.0,
			},
			TrainedBaseline: struct {
				Min  float64
				Max  float64
				Unit string
			}{
				Min:  2.0,
				Max:  6.0,
				Unit: "ng/mL",
			},
		},
	}
)

// GetAcutelyUpregulatedCytokines returns cytokines that increase acutely with exercise
func GetAcutelyUpregulatedCytokines() []Cytokine {
	return []Cytokine{IL6, TNF, IL10, IL1RA, TGFbeta}
}

// GetChronicallyDownregulatedCytokines returns inflammatory cytokines reduced with training
func GetChronicallyDownregulatedCytokines() []Cytokine {
	return []Cytokine{IL6, TNF} // These show reduced baseline and/or response with regular training
}

// GetAntiInflammatoryCytokines returns anti-inflammatory cytokines affected by exercise
func GetAntiInflammatoryCytokines() []Cytokine {
	return []Cytokine{IL10, IL1RA, TGFbeta}
}

// CalculateAcuteResponse estimates cytokine changes immediately after exercise
func CalculateAcuteResponse(cytokine Cytokine, exerciseIntensityPercent int,
	durationMinutes int, trainingStatus string) float64 {

	// Base response factor (1.0 = no change)
	responseFactor := 1.0

	// Convert parameters to normalized factors
	intensityFactor := float64(exerciseIntensityPercent) / 100.0
	durationFactor := float64(durationMinutes) / 60.0 // Normalized to 1 hour

	// Training status factor (trained individuals often have blunted acute inflammatory responses)
	trainingFactor := 1.0
	if trainingStatus == "Trained" || trainingStatus == "Athlete" {
		trainingFactor = 0.7 // 30% reduction in inflammatory response for trained individuals
	}

	// Calculate response based on specific cytokine
	switch cytokine.Name {
	case "Interleukin-6":
		// IL-6 increases dramatically with exercise, especially with longer durations
		if intensityFactor > 0.7 {
			responseFactor = 1.0 + (5.0 * intensityFactor * durationFactor * trainingFactor)
		} else {
			responseFactor = 1.0 + (3.0 * intensityFactor * durationFactor * trainingFactor)
		}

	case "Tumor Necrosis Factor-α":
		// TNF-α increases modestly with high-intensity exercise
		if intensityFactor > 0.8 || durationFactor > 1.5 {
			responseFactor = 1.0 + (1.5 * intensityFactor * trainingFactor)
		} else {
			responseFactor = 1.0 + (0.5 * intensityFactor * trainingFactor)
		}

	case "Interleukin-10":
		// IL-10 increases following the inflammatory response
		responseFactor = 1.0 + (2.0 * intensityFactor * durationFactor * trainingFactor)

	case "Interleukin-1 Receptor Antagonist":
		// IL-1RA increases significantly with exercise
		responseFactor = 1.0 + (3.0 * intensityFactor * durationFactor * trainingFactor)

	case "Transforming Growth Factor-β":
		// TGF-β shows modest increases with exercise
		responseFactor = 1.0 + (0.5 * intensityFactor * durationFactor)
	}

	return responseFactor
}

// CalculateChronicAdaptation estimates long-term changes in cytokine response
// with regular exercise training over specified weeks
func CalculateChronicAdaptation(cytokine Cytokine, weeksDuration int,
	weeklyFrequency int, intensityPercent int) float64 {

	// Base adaptation factor (1.0 = no change from untrained)
	adaptationFactor := 1.0

	// Training volume factors
	durationFactor := float64(weeksDuration) / 12.0 // Normalized to 12 weeks
	if durationFactor > 1.5 {
		durationFactor = 1.5 // Cap the duration effect
	}

	frequencyFactor := float64(weeklyFrequency) / 3.0   // Normalized to 3x weekly
	intensityFactor := float64(intensityPercent) / 70.0 // Normalized to 70% intensity

	// Combined training stimulus
	trainingStimulus := durationFactor * frequencyFactor * intensityFactor

	// Calculate adaptation based on specific cytokine
	switch cytokine.Name {
	case "Interleukin-6":
		// IL-6: Reduced baseline and blunted exercise response
		adaptationFactor = 1.0 - (0.2 * trainingStimulus)
		if adaptationFactor < 0.6 {
			adaptationFactor = 0.6 // Cap the reduction
		}

	case "Tumor Necrosis Factor-α":
		// TNF-α: Reduced baseline with training
		adaptationFactor = 1.0 - (0.25 * trainingStimulus)
		if adaptationFactor < 0.7 {
			adaptationFactor = 0.7 // Cap the reduction
		}

	case "Interleukin-10":
		// IL-10: Enhanced anti-inflammatory response
		adaptationFactor = 1.0 + (0.3 * trainingStimulus)

	case "Interleukin-1 Receptor Antagonist":
		// IL-1RA: Enhanced anti-inflammatory protection
		adaptationFactor = 1.0 + (0.25 * trainingStimulus)

	case "Transforming Growth Factor-β":
		// TGF-β: Complex adaptation, modest increase
		adaptationFactor = 1.0 + (0.1 * trainingStimulus)
	}

	return adaptationFactor
}
