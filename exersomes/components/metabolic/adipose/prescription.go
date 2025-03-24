package adipose

// ExercisePrescription defines exercise parameters targeting adipose tissue
type ExercisePrescription struct {
	Name                  string
	Description           string
	PrimaryType           string // "Aerobic", "Resistance", "HIIT", "Combined"
	SecondaryType         string // Optional secondary training type
	IntensityPercent      int    // % of max HR, VO2max, or 1RM
	DurationMinutes       int    // Per session
	FrequencyPerWeek      int
	TimeOfDay             string // "Morning", "Afternoon", "Evening", "Any"
	FeedingState          string // "Fasted", "Fed", "Either"
	TargetAdiposeFraction string // "White", "Brown", "Both"
	TargetAdipokines      []string
	TargetReceptors       []string
	ExpectedBenefits      []string
	TimeToEffect          struct {
		Acute   string // Acute effects timeframe
		Chronic string // Chronic adaptation timeframe
	}
}

// Calculate expected fat mass change in kg over time
func (p *ExercisePrescription) PredictFatLoss(weeks int, initialWeight float64, initialBodyFatPercent float64) float64 {
	// Simple model for demonstration, would be more sophisticated in practice

	// Base weekly fat loss rate factors by prescription type
	var weeklyRateFactor float64

	switch p.Name {
	case "Fat Loss HIIT":
		weeklyRateFactor = 0.004 // 0.4% of body weight per week
	case "Metabolic Health":
		weeklyRateFactor = 0.002 // 0.2% of body weight per week
	case "Brown Adipose Activation":
		weeklyRateFactor = 0.003 // 0.3% of body weight per week
	case "Combined Resistance-Cardio":
		weeklyRateFactor = 0.005 // 0.5% of body weight per week
	default:
		weeklyRateFactor = 0.002
	}

	// Adjust for frequency
	frequencyFactor := float64(p.FrequencyPerWeek) / 3.0 // Normalized to 3x/week
	weeklyRateFactor *= frequencyFactor

	// Calculate fat mass loss
	initialFatMass := initialWeight * initialBodyFatPercent / 100.0
	fatLoss := initialWeight * weeklyRateFactor * float64(weeks)

	// Cap fat loss to avoid unrealistic predictions
	maxSafeFatLoss := initialFatMass * 0.3 // Cap at 30% of initial fat mass
	if fatLoss > maxSafeFatLoss {
		fatLoss = maxSafeFatLoss
	}

	return fatLoss
}

// Standard adipose-targeting exercise prescriptions
var (
	FatLossHIIT = ExercisePrescription{
		Name:                  "Fat Loss HIIT",
		Description:           "High intensity interval training protocol optimized for fat loss",
		PrimaryType:           "HIIT",
		SecondaryType:         "Aerobic",
		IntensityPercent:      85,
		DurationMinutes:       20,
		FrequencyPerWeek:      3,
		TimeOfDay:             "Morning",
		FeedingState:          "Fasted",
		TargetAdiposeFraction: "White",
		TargetAdipokines:      []string{"Leptin", "Adiponectin", "IL-6"},
		TargetReceptors:       []string{"β3-Adrenergic Receptor"},
		ExpectedBenefits:      []string{"Increased fat oxidation", "Improved insulin sensitivity", "Reduced visceral fat"},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "24-48 hours",
			Chronic: "4-8 weeks",
		},
	}

	MetabolicHealth = ExercisePrescription{
		Name:                  "Metabolic Health",
		Description:           "Moderate intensity exercise protocol for improving adipose metabolic profile",
		PrimaryType:           "Aerobic",
		SecondaryType:         "",
		IntensityPercent:      65,
		DurationMinutes:       45,
		FrequencyPerWeek:      4,
		TimeOfDay:             "Any",
		FeedingState:          "Either",
		TargetAdiposeFraction: "White",
		TargetAdipokines:      []string{"Adiponectin", "TNF-α", "FABP4"},
		TargetReceptors:       []string{"Insulin Receptor", "PPAR-γ"},
		ExpectedBenefits:      []string{"Reduced inflammation", "Improved insulin sensitivity", "Healthier adipokine profile"},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "12-24 hours",
			Chronic: "6-12 weeks",
		},
	}

	BrownAdiposeActivation = ExercisePrescription{
		Name:                  "Brown Adipose Activation",
		Description:           "Cold-exposure combined with high-intensity exercise to increase BAT activity",
		PrimaryType:           "HIIT",
		SecondaryType:         "Aerobic",
		IntensityPercent:      80,
		DurationMinutes:       30,
		FrequencyPerWeek:      3,
		TimeOfDay:             "Morning",
		FeedingState:          "Either",
		TargetAdiposeFraction: "Brown",
		TargetAdipokines:      []string{"Irisin", "FGF21"},
		TargetReceptors:       []string{"β3-Adrenergic Receptor"},
		ExpectedBenefits:      []string{"Increased energy expenditure", "WAT browning", "Improved metabolic flexibility"},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "0-12 hours",
			Chronic: "3-6 weeks",
		},
	}

	CombinedResistanceCardio = ExercisePrescription{
		Name:                  "Combined Resistance-Cardio",
		Description:           "Integration of resistance and aerobic training for optimal body composition",
		PrimaryType:           "Combined",
		SecondaryType:         "",
		IntensityPercent:      75,
		DurationMinutes:       60,
		FrequencyPerWeek:      4,
		TimeOfDay:             "Afternoon",
		FeedingState:          "Fed",
		TargetAdiposeFraction: "Both",
		TargetAdipokines:      []string{"Adiponectin", "Leptin", "Irisin", "IL-6"},
		TargetReceptors:       []string{"Insulin Receptor", "β3-Adrenergic Receptor", "PPAR-γ"},
		ExpectedBenefits:      []string{"Fat loss", "Muscle preservation", "Metabolic flexibility", "Improved body composition"},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "24-48 hours",
			Chronic: "8-12 weeks",
		},
	}
)

// GetPrescriptionByTarget returns appropriate exercise prescriptions for specific targets
func GetPrescriptionByTarget(target string) []ExercisePrescription {
	var prescriptions []ExercisePrescription

	switch target {
	case "Fat loss":
		prescriptions = append(prescriptions, FatLossHIIT, CombinedResistanceCardio)
	case "Metabolic syndrome":
		prescriptions = append(prescriptions, MetabolicHealth, CombinedResistanceCardio)
	case "Inflammation":
		prescriptions = append(prescriptions, MetabolicHealth)
	case "Thermogenesis":
		prescriptions = append(prescriptions, BrownAdiposeActivation)
	case "Body composition":
		prescriptions = append(prescriptions, CombinedResistanceCardio)
	default:
		prescriptions = append(prescriptions, MetabolicHealth)
	}

	return prescriptions
}

// PredictAdiposeResponse estimates effects on adipose tissue over time
func PredictAdiposeResponse(prescription ExercisePrescription, weeks int, bodyFatPercent float64) map[string]float64 {
	// Convert prescription to parameters for adipokine prediction
	response := make(map[string]float64)

	// Get adipokine changes
	adipokines := PredictAdipokineLevels(prescription.PrimaryType, prescription.IntensityPercent, weeks, bodyFatPercent)
	for k, v := range adipokines {
		response[k] = v
	}

	// Add receptor sensitivity changes
	receptors := GetExerciseResponsiveReceptors()
	for _, receptor := range receptors {
		sensitivity := CalculateReceptorSensitization(receptor, prescription.PrimaryType,
			prescription.IntensityPercent, weeks)
		response[receptor.Name+" sensitivity"] = sensitivity
	}

	// Predict fat loss
	// Assuming 80kg person for this simple example
	fatLoss := prescription.PredictFatLoss(weeks, 80.0, bodyFatPercent)
	response["Fat mass change (kg)"] = -fatLoss

	return response
}
