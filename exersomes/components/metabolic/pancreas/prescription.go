package pancreas

// ExercisePrescription defines exercise parameters targeting pancreatic health
type ExercisePrescription struct {
	Name                 string
	PrimaryType          string   // "Aerobic", "Resistance", "HIIT", "Combined"
	IntensityPercent     int      // % of max HR, VO2max, or 1RM
	DurationMinutes      int      // Session duration
	FrequencyPerWeek     int      // Sessions per week
	TargetExerkines      []string // Exerkines targeted by this prescription
	ExpectedBenefits     []string
	IndicationsFor       []string // Specific conditions this targets
	ContraindicationsFor []string // When not to use this prescription
	TimeCourse           string   // Acute vs chronic effects
}

// Calculate time to expected clinical benefit in weeks
func (p *ExercisePrescription) TimeToEffect() int {
	switch p.Name {
	case "Acute Glucose Control":
		return 0 // Immediate effect
	case "Beta Cell Support":
		return 4 // 4 weeks
	case "Insulin Sensitivity":
		return 2 // 2 weeks
	default:
		return 6 // Default timeframe
	}
}

// Standard pancreatic health exercise prescriptions
var (
	AcuteGlucoseControl = ExercisePrescription{
		Name:             "Acute Glucose Control",
		PrimaryType:      "Aerobic",
		IntensityPercent: 60,
		DurationMinutes:  20,
		FrequencyPerWeek: 1,
		TargetExerkines:  []string{"Insulin", "Glucagon"},
		ExpectedBenefits: []string{
			"Reduced postprandial hyperglycemia",
			"Enhanced insulin sensitivity for 24-48 hours",
		},
		IndicationsFor: []string{
			"Type 2 diabetes",
			"Prediabetes",
			"Postprandial hyperglycemia",
		},
		ContraindicationsFor: []string{
			"Uncontrolled diabetes",
			"Recent diabetic ketoacidosis",
		},
		TimeCourse: "Acute (hours)",
	}

	BetaCellSupport = ExercisePrescription{
		Name:             "Beta Cell Support",
		PrimaryType:      "Combined",
		IntensityPercent: 70,
		DurationMinutes:  45,
		FrequencyPerWeek: 3,
		TargetExerkines:  []string{"Insulin", "Amylin"},
		ExpectedBenefits: []string{
			"Improved beta cell function",
			"Enhanced beta cell mass preservation",
			"Improved insulin secretion dynamics",
		},
		IndicationsFor: []string{
			"Early-stage type 2 diabetes",
			"Prediabetes",
		},
		ContraindicationsFor: []string{
			"Advanced beta cell failure",
		},
		TimeCourse: "Chronic (weeks to months)",
	}

	InsulinSensitivity = ExercisePrescription{
		Name:             "Insulin Sensitivity",
		PrimaryType:      "HIIT",
		IntensityPercent: 85,
		DurationMinutes:  20,
		FrequencyPerWeek: 3,
		TargetExerkines:  []string{"Insulin", "Glucagon"},
		ExpectedBenefits: []string{
			"Increased muscle GLUT4 expression",
			"Enhanced skeletal muscle insulin sensitivity",
			"Improved glycemic control",
		},
		IndicationsFor: []string{
			"Insulin resistance",
			"Type 2 diabetes",
			"Metabolic syndrome",
		},
		ContraindicationsFor: []string{
			"Uncontrolled hypertension",
			"Advanced cardiovascular disease",
		},
		TimeCourse: "Subacute (days to weeks)",
	}
)

// GetPrescriptionByCondition returns appropriate exercise prescriptions for a condition
func GetPrescriptionByCondition(condition string) []ExercisePrescription {
	var prescriptions []ExercisePrescription

	switch condition {
	case "Type 2 diabetes":
		prescriptions = append(prescriptions, AcuteGlucoseControl, InsulinSensitivity)
	case "Prediabetes":
		prescriptions = append(prescriptions, BetaCellSupport, InsulinSensitivity)
	case "Insulin resistance":
		prescriptions = append(prescriptions, InsulinSensitivity)
	default:
		prescriptions = append(prescriptions, AcuteGlucoseControl)
	}

	return prescriptions
}
