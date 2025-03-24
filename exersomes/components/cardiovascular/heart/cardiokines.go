package heart

// Cardiokine represents a signaling molecule secreted from cardiac tissue
type Cardiokine struct {
	Name               string
	MolecularWeight    float64  // In kDa
	SourceCells        []string // Cell types producing this factor
	TargetOrgans       []strings
	ExerciseRegulation string // "Up", "Down", "Biphasic"
	TemporalPattern    string // "Acute", "Chronic", "Both"
	PrimaryEffects     []string
	CardiacEffect      string // Primary effect on cardiac function
	BaselineRange      struct {
		Min  float64
		Max  float64
		Unit string
	}
}

// Key cardiokines affected by exercise
var (
	NatriureticPeptides = Cardiokine{
		Name:               "Natriuretic Peptides (ANP/BNP)",
		MolecularWeight:    3.5, // ANP ~3.5, BNP ~3.9
		SourceCells:        []string{"Cardiomyocytes", "Atrial cells"},
		TargetOrgans:       []string{"Kidney", "Vasculature", "Adipose", "Brain"},
		ExerciseRegulation: "Up",
		TemporalPattern:    "Both",
		PrimaryEffects:     []string{"Natriuresis", "Vasodilation", "Lipolysis", "Anti-fibrotic"},
		CardiacEffect:      "Reduced cardiac load, anti-hypertrophic",
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  5.0,
			Max:  100.0,
			Unit: "pg/mL", // For BNP
		},
	}

	FGF23 = Cardiokine{
		Name:               "Fibroblast Growth Factor 23",
		MolecularWeight:    32.0,
		SourceCells:        []string{"Cardiomyocytes", "Fibroblasts"},
		TargetOrgans:       []string{"Kidney", "Parathyroid", "Heart"},
		ExerciseRegulation: "Up", // Acute increase
		TemporalPattern:    "Acute",
		PrimaryEffects:     []string{"Phosphate regulation", "Vitamin D metabolism", "Calcium homeostasis"},
		CardiacEffect:      "Hypertrophy (chronic elevation)", // But exercise-induced elevation is transient
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  40.0,
			Max:  100.0,
			Unit: "pg/mL",
		},
	}

	GDF15 = Cardiokine{
		Name:               "Growth Differentiation Factor 15",
		MolecularWeight:    35.0,
		SourceCells:        []string{"Cardiomyocytes", "Macrophages"},
		TargetOrgans:       []string{"Heart", "Brain", "Adipose", "Muscle"},
		ExerciseRegulation: "Biphasic", // Acute increase, chronic decrease in baseline
		TemporalPattern:    "Both",
		PrimaryEffects:     []string{"Anti-inflammatory", "Anti-hypertrophic", "Metabolic regulation"},
		CardiacEffect:      "Cardioprotective, anti-remodeling",
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  200.0,
			Max:  1200.0,
			Unit: "pg/mL",
		},
	}

	Follistatin3 = Cardiokine{
		Name:               "Follistatin-like 3",
		MolecularWeight:    27.0,
		SourceCells:        []string{"Cardiomyocytes"},
		TargetOrgans:       []string{"Heart", "Vasculature"},
		ExerciseRegulation: "Down", // With chronic training
		TemporalPattern:    "Chronic",
		PrimaryEffects:     []string{"TGF-β pathway modulation", "Metabolic regulation"},
		CardiacEffect:      "Anti-hypertrophic",
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  1.0,
			Max:  5.0,
			Unit: "ng/mL",
		},
	}

	Adropin = Cardiokine{
		Name:               "Adropin",
		MolecularWeight:    4.5,
		SourceCells:        []string{"Cardiomyocytes", "Endothelial cells"},
		TargetOrgans:       []string{"Endothelium", "Liver", "Brain"},
		ExerciseRegulation: "Up",
		TemporalPattern:    "Both",
		PrimaryEffects:     []string{"Endothelial function", "Energy homeostasis", "Insulin sensitivity"},
		CardiacEffect:      "Improved endothelial function, anti-atherogenic",
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  1.0,
			Max:  5.0,
			Unit: "ng/mL",
		},
	}

	CTRP9 = Cardiokine{
		Name:               "C1q/TNF-related protein 9",
		MolecularWeight:    40.0,
		SourceCells:        []string{"Cardiomyocytes", "Epicardial adipocytes"},
		TargetOrgans:       []string{"Heart", "Vasculature"},
		ExerciseRegulation: "Up",
		TemporalPattern:    "Chronic",
		PrimaryEffects:     []string{"Glucose metabolism", "Fatty acid oxidation", "Anti-inflammatory"},
		CardiacEffect:      "Cardioprotective, anti-apoptotic",
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  5.0,
			Max:  40.0,
			Unit: "ng/mL",
		},
	}
)

// GetExerciseUpregulatedCardiokines returns cardiokines that increase with exercise
func GetExerciseUpregulatedCardiokines() []Cardiokine {
	return []Cardiokine{NatriureticPeptides, FGF23, Adropin, CTRP9}
}

// GetExerciseDownregulatedCardiokines returns cardiokines that decrease with exercise
func GetExerciseDownregulatedCardiokines() []Cardiokine {
	return []Cardiokine{Follistatin3}
}

// CalculateExerciseResponse estimates cardiokine changes from exercise
func CalculateExerciseResponse(cardiokine Cardiokine, exerciseType string,
	intensityPercent int, durationMinutes int, trainingWeeks int) float64 {

	// Base level factor (centered at 1.0 = no change)
	var changeFactor float64 = 1.0

	// Acute response factors
	acuteIntensityEffect := float64(intensityPercent) / 100.0
	acuteDurationEffect := float64(durationMinutes) / 60.0

	// Chronic adaptation factors
	chronicEffect := float64(trainingWeeks) / 12.0 // Normalized to 12 weeks

	// Different cardiokines have distinct response patterns to exercise
	switch cardiokine.Name {
	case "Natriuretic Peptides (ANP/BNP)":
		// Strongly intensity-dependent, increases with both acute and chronic exercise
		acuteEffect := 1.0 + (0.8 * acuteIntensityEffect * acuteDurationEffect)
		chronicEffect := 1.0 + (0.2 * float64(trainingWeeks) / 12.0) // Resting levels also increase

		if exerciseType == "HIIT" || exerciseType == "Aerobic" {
			// Greater effect with cardio exercise
			changeFactor = acuteEffect * chronicEffect * 1.2
		} else {
			changeFactor = acuteEffect * chronicEffect
		}

	case "Fibroblast Growth Factor 23":
		// Primarily acute response to intense exercise
		if intensityPercent > 70 {
			changeFactor = 1.0 + (0.4 * acuteIntensityEffect * acuteDurationEffect)
		} else {
			changeFactor = 1.0 + (0.1 * acuteIntensityEffect * acuteDurationEffect)
		}

	case "Growth Differentiation Factor 15":
		// Acute increase but chronic decrease in baseline
		acuteEffect := 1.0 + (0.5 * acuteIntensityEffect * acuteDurationEffect)
		chronicEffect := 1.0 - (0.2 * float64(trainingWeeks) / 12.0)

		// Decide whether to return acute or chronic effect based on context
		if trainingWeeks < 2 {
			changeFactor = acuteEffect // Acute response dominates early
		} else {
			changeFactor = chronicEffect // Chronic adaptation takes over
		}

	case "Follistatin-like 3":
		// Gradually decreases with training
		if trainingWeeks > 4 {
			changeFactor = 1.0 - (0.15 * float64(trainingWeeks) / 12.0)
			if changeFactor < 0.7 { // Floor at 30% reduction
				changeFactor = 0.7
			}
		}

	case "Adropin":
		// Increases with both acute and chronic exercise
		if exerciseType == "Aerobic" {
			acuteEffect := 1.0 + (0.3 * acuteDurationEffect)
			chronicEffect := 1.0 + (0.3 * float64(trainingWeeks) / 12.0)
			changeFactor = acuteEffect * chronicEffect
		} else {
			acuteEffect := 1.0 + (0.1 * acuteDurationEffect)
			chronicEffect := 1.0 + (0.2 * float64(trainingWeeks) / 12.0)
			changeFactor = acuteEffect * chronicEffect
		}

	case "C1q/TNF-related protein 9":
		// Primarily increases with chronic training
		changeFactor = 1.0 + (0.3 * float64(trainingWeeks) / 12.0)
	}

	return changeFactor
}

// PredictCardiokineResponse estimates cardiokine levels after an exercise protocol
func PredictCardiokineResponse(exerciseType string, intensityPercent int,
	durationMinutes int, trainingWeeks int, hasHeartDisease bool) map[string]float64 {

	// Get all cardiokines
	cardiokines := append(GetExerciseUpregulatedCardiokines(), GetExerciseDownregulatedCardiokines()...)

	// Calculate responses
	responses := make(map[string]float64)

	for _, cardiokine := range cardiokines {
		// Calculate baseline level (middle of reference range)
		baseline := (cardiokine.BaselineRange.Min + cardiokine.BaselineRange.Max) / 2.0

		// Adjust baseline for heart disease if applicable
		if hasHeartDisease {
			switch cardiokine.Name {
			case "Natriuretic Peptides (ANP/BNP)":
				baseline *= 2.5 // Significantly elevated in heart disease
			case "Growth Differentiation Factor 15":
				baseline *= 1.8 // Elevated in heart failure
			case "Follistatin-like 3":
				baseline *= 1.4 // Moderately elevated
			case "C1q/TNF-related protein 9":
				baseline *= 0.7 // Reduced in heart disease
			}
		}

		// Calculate response factor
		responseFactor := CalculateExerciseResponse(cardiokine, exerciseType,
			intensityPercent, durationMinutes, trainingWeeks)

		// Calculate predicted level
		responses[cardiokine.Name] = baseline * responseFactor
	}

	return responses
}
