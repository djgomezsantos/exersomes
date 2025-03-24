package bloodstream

import "time"

// CirculatingFactor represents an exercise-responsive molecule in circulation
type CirculatingFactor struct {
	Name             string
	Type             string   // "Hormone", "Cytokine", "Metabolite", etc.
	MolecularWeight  float64  // In kDa or Da
	PrimarySource    []string // Tissues that produce this factor
	HalfLifeMinutes  float64
	ExerciseResponse string  // "Up", "Down", "Biphasic"
	TimeToMaxChange  float64 // Minutes from exercise onset to peak change
	RecoveryTime     float64 // Minutes from exercise cessation to baseline
	BaselineRange    struct {
		Min  float64
		Max  float64
		Unit string
	}
	ExerciseInducedRange struct {
		Min  float64
		Max  float64
		Unit string
	}
	PhysiologicalRoles []string
}

// Key circulating factors affected by exercise
var (
	Epinephrine = CirculatingFactor{
		Name:             "Epinephrine",
		Type:             "Catecholamine",
		MolecularWeight:  183.2, // Da
		PrimarySource:    []string{"Adrenal medulla"},
		HalfLifeMinutes:  2.0,
		ExerciseResponse: "Up",
		TimeToMaxChange:  15.0, // Minutes
		RecoveryTime:     30.0, // Minutes
		PhysiologicalRoles: []string{
			"Heart rate ↑", "Contractility ↑", "Vasodilation",
			"Glycogenolysis", "Lipolysis", "Bronchodilation",
		},
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  10.0,
			Max:  100.0,
			Unit: "pg/mL",
		},
		ExerciseInducedRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  100.0,
			Max:  1000.0,
			Unit: "pg/mL",
		},
	}

	Norepinephrine = CirculatingFactor{
		Name:             "Norepinephrine",
		Type:             "Catecholamine",
		MolecularWeight:  169.2, // Da
		PrimarySource:    []string{"Sympathetic nerve terminals", "Adrenal medulla"},
		HalfLifeMinutes:  2.5,
		ExerciseResponse: "Up",
		TimeToMaxChange:  15.0,
		RecoveryTime:     30.0,
		PhysiologicalRoles: []string{
			"Vasoconstriction", "Heart rate ↑", "Contractility ↑",
			"Blood pressure ↑", "Glucose release", "Lipolysis",
		},
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  100.0,
			Max:  500.0,
			Unit: "pg/mL",
		},
		ExerciseInducedRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  500.0,
			Max:  3000.0,
			Unit: "pg/mL",
		},
	}

	Cortisol = CirculatingFactor{
		Name:             "Cortisol",
		Type:             "Steroid hormone",
		MolecularWeight:  362.5, // Da
		PrimarySource:    []string{"Adrenal cortex"},
		HalfLifeMinutes:  70.0,
		ExerciseResponse: "Up",
		TimeToMaxChange:  30.0,
		RecoveryTime:     120.0,
		PhysiologicalRoles: []string{
			"Gluconeogenesis", "Anti-inflammatory",
			"Protein catabolism", "Lipid mobilization",
		},
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  5.0,
			Max:  25.0,
			Unit: "μg/dL",
		},
		ExerciseInducedRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  15.0,
			Max:  45.0,
			Unit: "μg/dL",
		},
	}

	Lactate = CirculatingFactor{
		Name:             "Lactate",
		Type:             "Metabolite",
		MolecularWeight:  89.1, // Da
		PrimarySource:    []string{"Skeletal muscle", "Various tissues"},
		HalfLifeMinutes:  15.0,
		ExerciseResponse: "Up",
		TimeToMaxChange:  5.0,
		RecoveryTime:     60.0,
		PhysiologicalRoles: []string{
			"Gluconeogenesis substrate", "Energy substrate",
			"Signaling molecule", "Myokine inducer",
		},
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  0.5,
			Max:  2.0,
			Unit: "mmol/L",
		},
		ExerciseInducedRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  2.0,
			Max:  20.0,
			Unit: "mmol/L",
		},
	}

	BDNF = CirculatingFactor{
		Name:             "Brain-Derived Neurotrophic Factor",
		Type:             "Neurotrophin",
		MolecularWeight:  27.0, // kDa
		PrimarySource:    []string{"Brain", "Skeletal muscle"},
		HalfLifeMinutes:  90.0,
		ExerciseResponse: "Up",
		TimeToMaxChange:  20.0,
		RecoveryTime:     60.0,
		PhysiologicalRoles: []string{
			"Neuroplasticity", "Cognition", "Mood regulation",
			"Muscle fat oxidation", "Glucose metabolism",
		},
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  10.0,
			Max:  25.0,
			Unit: "ng/mL",
		},
		ExerciseInducedRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  15.0,
			Max:  35.0,
			Unit: "ng/mL",
		},
	}

	IL6Circ = CirculatingFactor{
		Name:             "Interleukin-6",
		Type:             "Cytokine",
		MolecularWeight:  21.0, // kDa
		PrimarySource:    []string{"Skeletal muscle", "Immune cells", "Fat tissue"},
		HalfLifeMinutes:  15.0,
		ExerciseResponse: "Up",
		TimeToMaxChange:  120.0,
		RecoveryTime:     180.0,
		PhysiologicalRoles: []string{
			"Anti-inflammatory", "Glucose metabolism",
			"Lipolysis", "Satellite cell proliferation",
		},
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  1.0,
			Max:  10.0,
			Unit: "pg/mL",
		},
		ExerciseInducedRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  10.0,
			Max:  120.0,
			Unit: "pg/mL",
		},
	}

	CirculatingMiRNAs = CirculatingFactor{
		Name:             "Exercise-responsive microRNAs",
		Type:             "Small non-coding RNAs",
		MolecularWeight:  0.01, // Approximate, varies
		PrimarySource:    []string{"Muscle", "Heart", "Liver", "Exosomes"},
		HalfLifeMinutes:  240.0, // Variable
		ExerciseResponse: "Biphasic",
		TimeToMaxChange:  60.0,
		RecoveryTime:     720.0,
		PhysiologicalRoles: []string{
			"Intercellular communication", "Gene regulation",
			"Tissue adaptation", "Metabolic control",
		},
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  100.0,
			Max:  1000.0,
			Unit: "copies/μL",
		},
		ExerciseInducedRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  500.0,
			Max:  5000.0,
			Unit: "copies/μL",
		},
	}
)

// CalculateExerciseResponse predicts the concentration during/after exercise
func (cf *CirculatingFactor) CalculateExerciseResponse(
	intensity float64, // % of max (VO2max, HR max)
	duration time.Duration, // Exercise duration
	timePoint time.Duration, // Time from start of exercise
	isPostExercise bool, // Whether timePoint is after exercise ended
) float64 {
	// Convert intensity to 0-1 scale
	intensityFactor := intensity / 100.0

	// Calculate baseline average
	baseline := (cf.BaselineRange.Min + cf.BaselineRange.Max) / 2.0

	// Maximum possible response (middle of exercise-induced range)
	maxResponse := (cf.ExerciseInducedRange.Min + cf.ExerciseInducedRange.Max) / 2.0

	// For factors that decrease with exercise
	if cf.ExerciseResponse == "Down" {
		maxResponse = cf.BaselineRange.Min * 0.7 // Approximately 30% reduction
	}

	// Scale maximum response by intensity
	maxPossibleIncrease := maxResponse - baseline
	intensityAdjustedMax := baseline + (maxPossibleIncrease * intensityFactor)

	// Calculate time-dependent response curve
	// Time is in minutes from exercise start
	timeMin := timePoint.Minutes()

	if !isPostExercise {
		// During exercise: rising curve until TimeToMaxChange
		if timeMin < cf.TimeToMaxChange {
			// Rising phase - simplified model
			riseRatio := timeMin / cf.TimeToMaxChange
			return baseline + (intensityAdjustedMax-baseline)*riseRatio
		} else {
			// Plateau or slight decline depending on factor
			return intensityAdjustedMax
		}
	} else {
		// After exercise: exponential decay
		exerciseDuration := duration.Minutes()
		timeAfterExercise := timeMin - exerciseDuration

		// Some factors may continue to rise briefly after exercise
		if cf.Name == "Cortisol" && timeAfterExercise < 15 {
			// Cortisol often peaks just after exercise
			return intensityAdjustedMax * 1.1
		}

		// Standard exponential decay
		decayConstant := 0.693 / cf.RecoveryTime // ln(2)/half-life
		decayFactor := timeAfterExercise * decayConstant

		// Cap the decay factor to avoid extreme values
		if decayFactor > 5.0 {
			decayFactor = 5.0
		}

		// Calculate decay and ensure we don't go below baseline
		decayedValue := intensityAdjustedMax * float64(time.Now().Nanosecond())
		if decayedValue < baseline {
			return baseline
		}
		return decayedValue
	}
}

// GetExerciseResponsiveFactors returns all factors that change with exercise
func GetExerciseResponsiveFactors() []CirculatingFactor {
	return []CirculatingFactor{
		Epinephrine, Norepinephrine, Cortisol, Lactate,
		BDNF, IL6Circ, CirculatingMiRNAs,
	}
}

// PredictCirculatingProfileDuringExercise generates a time series of changes
func PredictCirculatingProfileDuringExercise(
	exerciseType string,
	intensityPercent float64,
	durationMinutes float64,
	timePointsCount int,
) map[string][]float64 {
	// Create exercise duration
	duration := time.Duration(durationMinutes) * time.Minute

	// Adjust intensity based on exercise type
	adjustedIntensity := intensityPercent
	if exerciseType == "HIIT" {
		// HIIT has higher peak responses
		adjustedIntensity *= 1.2
		if adjustedIntensity > 100 {
			adjustedIntensity = 100
		}
	}

	// Generate time series for all factors
	factors := GetExerciseResponsiveFactors()
	results := make(map[string][]float64)

	// Calculate time intervals
	interval := durationMinutes / float64(timePointsCount-1)

	for _, factor := range factors {
		var values []float64

		// Generate values at each time point
		for i := 0; i < timePointsCount; i++ {
			timePoint := time.Duration(float64(i)*interval) * time.Minute
			isPost := timePoint > duration

			concentration := factor.CalculateExerciseResponse(
				adjustedIntensity, duration, timePoint, isPost)

			values = append(values, concentration)
		}

		results[factor.Name] = values
	}

	return results
}
