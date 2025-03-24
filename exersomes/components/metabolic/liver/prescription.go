package liver

// ExercisePrescription defines exercise parameters targeting liver health
type ExercisePrescription struct {
	Name              string
	Description       string
	PrimaryType       string // "Aerobic", "Resistance", "HIIT", "Combined"
	SecondaryType     string // Optional secondary component
	IntensityPercent  int    // % of max HR, VO2max, or 1RM
	DurationMinutes   int    // Per session
	FrequencyPerWeek  int
	TargetHepatokines []string // Hepatokines targeted
	TargetReceptors   []string // Liver receptors targeted
	ExpectedBenefits  []string // Expected liver health improvements
	DietaryContext    string   // "Fasted", "Carb-restricted", "Regular"
	IndicationsFor    []string // Specific liver conditions
	Contraindications []string // When not to use this protocol
	TimeToEffect      struct {
		Acute   string // Acute effects
		Chronic string // Chronic effects
	}
}

// Standard liver-targeting exercise prescriptions
var (
	NAFLDReduction = ExercisePrescription{
		Name:              "NAFLD Reduction Protocol",
		Description:       "Combined exercise intervention to reduce hepatic fat and improve liver function",
		PrimaryType:       "Aerobic",
		SecondaryType:     "Resistance",
		IntensityPercent:  70,
		DurationMinutes:   45,
		FrequencyPerWeek:  4,
		TargetHepatokines: []string{"FGF21", "Fetuin-A", "Follistatin"},
		TargetReceptors:   []string{"AMPK", "PPAR-α", "Insulin Receptor"},
		ExpectedBenefits: []string{
			"Reduced liver fat content",
			"Improved insulin sensitivity",
			"Enhanced fatty acid oxidation",
			"Decreased hepatic inflammation",
		},
		DietaryContext:    "Carb-restricted",
		IndicationsFor:    []string{"NAFLD", "Metabolic-associated fatty liver disease", "Hepatic steatosis"},
		Contraindications: []string{"Decompensated cirrhosis", "Severe cardiovascular disease"},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "Minor changes in enzymes within 24h",
			Chronic: "10% liver fat reduction in 8-12 weeks",
		},
	}

	HepatitisCProtocol = ExercisePrescription{
		Name:              "Hepatitis C Adjunct Protocol",
		Description:       "Moderate activity to support standard treatment and reduce progression",
		PrimaryType:       "Aerobic",
		SecondaryType:     "",
		IntensityPercent:  60,
		DurationMinutes:   30,
		FrequencyPerWeek:  5,
		TargetHepatokines: []string{"FGF21", "Follistatin"},
		TargetReceptors:   []string{"AMPK", "TNF Receptor"},
		ExpectedBenefits: []string{
			"Reduced hepatic inflammation",
			"Improved response to antiviral therapy",
			"Decreased fibrosis progression",
			"Enhanced quality of life",
		},
		DietaryContext:    "Regular",
		IndicationsFor:    []string{"Chronic Hepatitis C", "Post-treatment maintenance"},
		Contraindications: []string{"Acute hepatitis flare", "Severe thrombocytopenia", "Portal hypertension"},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "Improved well-being within days",
			Chronic: "Reduced inflammation markers in 12-16 weeks",
		},
	}

	LiverGlucoseMetabolism = ExercisePrescription{
		Name:              "Hepatic Glucose Metabolism Optimizing Protocol",
		Description:       "High-intensity interval training to maximize insulin sensitivity and glucose handling",
		PrimaryType:       "HIIT",
		SecondaryType:     "Aerobic",
		IntensityPercent:  85,
		DurationMinutes:   25,
		FrequencyPerWeek:  3,
		TargetHepatokines: []string{"FGF21", "Selenoprotein P"},
		TargetReceptors:   []string{"Insulin Receptor", "GLP-1 Receptor", "AMPK"},
		ExpectedBenefits: []string{
			"Enhanced hepatic insulin sensitivity",
			"Reduced glucose production",
			"Improved postprandial glucose handling",
			"Decreased fasting glucose",
		},
		DietaryContext:    "Carb timing around exercise",
		IndicationsFor:    []string{"Insulin resistance", "Prediabetes", "Type 2 diabetes"},
		Contraindications: []string{"Uncontrolled diabetes", "Advanced cardiovascular disease"},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "Improved insulin sensitivity for 24-48 hours",
			Chronic: "Significant metabolic improvements in 6-8 weeks",
		},
	}

	LiverFibrosis = ExercisePrescription{
		Name:              "Anti-Fibrotic Liver Protocol",
		Description:       "Gentle progressive exercise approach to prevent fibrosis progression",
		PrimaryType:       "Aerobic",
		SecondaryType:     "Flexibility",
		IntensityPercent:  50,
		DurationMinutes:   30,
		FrequencyPerWeek:  5,
		TargetHepatokines: []string{"FGF21"},
		TargetReceptors:   []string{"PPAR-α", "GLP-1 Receptor"},
		ExpectedBenefits: []string{
			"Reduced stellate cell activation",
			"Decreased profibrotic signaling",
			"Improved liver blood flow",
			"Enhanced quality of life",
		},
		DietaryContext:    "Regular, alcohol-free",
		IndicationsFor:    []string{"Early fibrosis", "NASH", "Alcoholic liver disease recovery"},
		Contraindications: []string{"Esophageal varices", "Hepatic encephalopathy", "Severe portal hypertension"},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "No immediate effects",
			Chronic: "Potential stabilization of fibrosis in 16-24 weeks",
		},
	}
)

// GetPrescriptionByCondition returns appropriate exercise prescriptions for liver conditions
func GetPrescriptionByCondition(condition string) []ExercisePrescription {
	var prescriptions []ExercisePrescription

	switch condition {
	case "NAFLD":
		prescriptions = []ExercisePrescription{NAFLDReduction, LiverGlucoseMetabolism}
	case "Hepatitis C":
		prescriptions = []ExercisePrescription{HepatitisCProtocol}
	case "Fibrosis":
		prescriptions = []ExercisePrescription{LiverFibrosis}
	case "Insulin Resistance":
		prescriptions = []ExercisePrescription{LiverGlucoseMetabolism, NAFLDReduction}
	default:
		prescriptions = []ExercisePrescription{NAFLDReduction}
	}

	return prescriptions
}

// PredictLiverResponse estimates the effects of an exercise prescription on liver health
func PredictLiverResponse(prescription ExercisePrescription, weeks int, hasNAFLD bool,
	hasInsulinResistance bool) map[string]interface{} {

	response := make(map[string]interface{})

	// Predict hepatokine changes
	hepatokines := PredictHepatokineLevels(prescription.PrimaryType,
		prescription.IntensityPercent, prescription.DurationMinutes, weeks, hasNAFLD)

	for k, v := range hepatokines {
		response[k+"_level"] = v
	}

	// Predict receptor changes
	receptors := GetExerciseResponsiveReceptors()
	for _, receptor := range receptors {
		receptorResponse := CalculateReceptorResponse(receptor,
			prescription.PrimaryType, prescription.IntensityPercent, weeks)

		response[receptor.Name+"_expression"] = receptorResponse.ExpressionChange
		response[receptor.Name+"_sensitivity"] = receptorResponse.SensitivityChange
	}

	// Predict clinical outcomes
	liverFatChange := 0.0
	insulinSensitivityChange := 0.0
	inflammationChange := 0.0

	// Liver fat reduction calculation
	switch prescription.Name {
	case "NAFLD Reduction Protocol":
		liverFatChange = -5.0 - (float64(weeks) * 0.6) // About 5% baseline plus 0.6% per week
	case "Hepatic Glucose Metabolism Optimizing Protocol":
		liverFatChange = -3.0 - (float64(weeks) * 0.4) // About 3% baseline plus 0.4% per week
	default:
		liverFatChange = -2.0 - (float64(weeks) * 0.3) // About 2% baseline plus 0.3% per week
	}

	// Cap the maximum change
	if liverFatChange < -30.0 {
		liverFatChange = -30.0
	}

	// Adjust for starting conditions
	if hasNAFLD {
		liverFatChange *= 1.5 // Greater reduction potential if starting with NAFLD
	}

	// Insulin sensitivity improvement
	insulinSensitivityChange = 10.0 + (float64(weeks) * 1.2)
	if hasInsulinResistance {
		insulinSensitivityChange *= 1.3 // Greater improvement if insulin resistant at baseline
	}

	// Inflammation reduction
	inflammationChange = -5.0 - (float64(weeks) * 0.8)

	// Store clinical outcomes
	response["liver_fat_percent_change"] = liverFatChange
	response["insulin_sensitivity_percent_improvement"] = insulinSensitivityChange
	response["inflammation_percent_reduction"] = inflammationChange

	// Predict liver enzyme changes
	response["ALT_percent_change"] = -5.0 - (float64(weeks) * 0.7)
	response["AST_percent_change"] = -4.0 - (float64(weeks) * 0.6)

	return response
}
