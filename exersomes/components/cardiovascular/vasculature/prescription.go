package bloodstream

// ExercisePrescription defines parameters for bloodstream-targeted exercise
type ExercisePrescription struct {
	Name               string
	Description        string
	PrimaryType        string // "Aerobic", "HIIT", "Resistance", "Combined"
	IntensityPercent   int    // % of max HR, VO2max, or 1RM
	DurationMinutes    int
	FrequencyPerWeek   int
	TargetCircFactors  []string // Target circulating factors
	TargetCells        []string // Target immune/blood cells
	ExpectedBenefits   []string
	HemodynamicEffects []string // Effects on blood flow, pressure, etc.
	IndicationsFor     []string
	Contraindications  []string
	TimeToEffect       struct {
		Acute   string // Immediate effects
		Chronic string // Long-term adaptations
	}
}

// CirculatoryResponse represents the predicted changes in bloodstream parameters
type CirculatoryResponse struct {
	FactorChanges       map[string]float64 // Changes in circulating factors
	CellChanges         map[string]float64 // Changes in immune cell counts/function
	FlowMetrics         map[string]float64 // Blood flow parameters
	CoagulationMetrics  map[string]float64 // Clotting/bleeding metrics
	InflammationMetrics map[string]float64 // Inflammatory markers
}

// Standard bloodstream-targeting exercise prescriptions
var (
	EndothelialHealth = ExercisePrescription{
		Name:              "Endothelial Health Protocol",
		Description:       "Moderate intensity continuous training to improve endothelial function",
		PrimaryType:       "Aerobic",
		IntensityPercent:  65,
		DurationMinutes:   45,
		FrequencyPerWeek:  4,
		TargetCircFactors: []string{"Nitric Oxide", "VEGF", "ET-1", "Prostacyclin"},
		TargetCells:       []string{"Endothelial Cells", "EPCs"},
		ExpectedBenefits: []string{
			"Improved flow-mediated dilation",
			"Enhanced nitric oxide bioavailability",
			"Reduced endothelial inflammation",
			"Increased capillarization",
		},
		HemodynamicEffects: []string{
			"Reduced peripheral resistance",
			"Improved microvascular perfusion",
			"Enhanced vasodilatory capacity",
		},
		IndicationsFor: []string{
			"Hypertension",
			"Early atherosclerosis",
			"Endothelial dysfunction",
			"Microvascular disease",
		},
		Contraindications: []string{
			"Severe aortic stenosis",
			"Hypertrophic cardiomyopathy with outflow obstruction",
			"Acute vascular injury",
		},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "Flow improvements within 1-3 hours",
			Chronic: "Structural adaptation in 4-8 weeks",
		},
	}

	InflammationReduction = ExercisePrescription{
		Name:              "Vascular Inflammation Reduction",
		Description:       "Low-to-moderate intensity exercise to reduce systemic inflammation",
		PrimaryType:       "Aerobic",
		IntensityPercent:  55,
		DurationMinutes:   40,
		FrequencyPerWeek:  5,
		TargetCircFactors: []string{"IL-6", "IL-10", "CRP", "TNF-α", "IL-1Ra"},
		TargetCells:       []string{"Monocytes", "T-cells", "Neutrophils"},
		ExpectedBenefits: []string{
			"Reduced inflammatory cytokines",
			"Shift toward anti-inflammatory phenotype",
			"Decreased vascular inflammation markers",
			"Improved metabolic profile",
		},
		HemodynamicEffects: []string{
			"Reduced arterial stiffness",
			"Improved endothelial function",
		},
		IndicationsFor: []string{
			"Chronic low-grade inflammation",
			"Metabolic syndrome",
			"Atherosclerosis",
			"Autoimmune conditions",
		},
		Contraindications: []string{
			"Acute inflammatory flare",
			"Febrile illness",
			"Recent surgery",
		},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "Minor changes within 1-2 hours",
			Chronic: "Significant reduction in 6-12 weeks",
		},
	}

	AnticoagulationProtocol = ExercisePrescription{
		Name:              "Coagulation Profile Improvement",
		Description:       "Moderate-intensity exercise designed to optimize coagulation balance",
		PrimaryType:       "Combined",
		IntensityPercent:  60,
		DurationMinutes:   35,
		FrequencyPerWeek:  4,
		TargetCircFactors: []string{"Fibrinogen", "D-dimer", "PAI-1", "tPA"},
		TargetCells:       []string{"Platelets", "Endothelial Cells"},
		ExpectedBenefits: []string{
			"Increased fibrinolytic activity",
			"Reduced platelet aggregation",
			"Balanced coagulation profile",
			"Decreased thrombotic risk",
		},
		HemodynamicEffects: []string{
			"Improved blood fluidity",
			"Reduced abnormal clotting tendency",
		},
		IndicationsFor: []string{
			"Hypercoagulable states",
			"Sedentary lifestyle",
			"Post-thrombotic syndrome",
			"Metabolic syndrome",
		},
		Contraindications: []string{
			"Acute bleeding",
			"Severe thrombocytopenia",
			"Recent pulmonary embolism",
			"Unstable cardiovascular disease",
		},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "Changes in coagulation profile within 2-4 hours",
			Chronic: "Stable improvements in 4-8 weeks",
		},
	}

	CirculatingProgenitorStimulation = ExercisePrescription{
		Name:              "Progenitor Cell Mobilization",
		Description:       "Interval-based protocol to maximize mobilization of circulating progenitor cells",
		PrimaryType:       "HIIT",
		IntensityPercent:  85,
		DurationMinutes:   30,
		FrequencyPerWeek:  3,
		TargetCircFactors: []string{"VEGF", "G-CSF", "SDF-1", "NO"},
		TargetCells:       []string{"EPCs", "CD34+ Cells", "HSCs"},
		ExpectedBenefits: []string{
			"Increased circulating progenitor cells",
			"Enhanced vascular repair capacity",
			"Improved angiogenic potential",
			"Regeneration of damaged endothelium",
		},
		HemodynamicEffects: []string{
			"Increased peripheral blood flow",
			"Enhanced tissue perfusion during recovery",
		},
		IndicationsFor: []string{
			"Vascular repair needs",
			"Post-infarction recovery",
			"Peripheral arterial disease",
			"Diabetic vascular disease",
		},
		Contraindications: []string{
			"Bone marrow suppression",
			"Recent stem cell transplantation",
			"Hematological malignancies",
			"Severe cardiopulmonary disease",
		},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "Peak mobilization at 10-30 minutes post-exercise",
			Chronic: "Sustained increases after 4 weeks",
		},
	}
)

// GetPrescriptionByCondition returns appropriate exercise prescriptions for a condition
func GetPrescriptionByCondition(condition string) []ExercisePrescription {
	var prescriptions []ExercisePrescription

	switch condition {
	case "Atherosclerosis":
		prescriptions = append(prescriptions, EndothelialHealth, InflammationReduction)
	case "Hypertension":
		prescriptions = append(prescriptions, EndothelialHealth)
	case "Thrombosis Risk":
		prescriptions = append(prescriptions, AnticoagulationProtocol)
	case "Vascular Repair":
		prescriptions = append(prescriptions, CirculatingProgenitorStimulation)
	case "Metabolic Syndrome":
		prescriptions = append(prescriptions, InflammationReduction, EndothelialHealth)
	default:
		prescriptions = append(prescriptions, EndothelialHealth)
	}

	return prescriptions
}

// PredictCirculatoryResponse estimates the effects of an exercise prescription
// on bloodstream parameters based on the protocol, duration of adherence, and patient factors
func PredictCirculatoryResponse(prescription ExercisePrescription, weeks int,
	baselineInflammation bool, endothelialDysfunction bool) CirculatoryResponse {

	// Initialize response
	response := CirculatoryResponse{
		FactorChanges:       make(map[string]float64),
		CellChanges:         make(map[string]float64),
		FlowMetrics:         make(map[string]float64),
		CoagulationMetrics:  make(map[string]float64),
		InflammationMetrics: make(map[string]float64),
	}

	// Calculate intensity factor
	intensityFactor := float64(prescription.IntensityPercent) / 100.0

	// Calculate duration factor (diminishing returns after 8 weeks)
	durationFactor := 0.0
	if weeks <= 8 {
		durationFactor = float64(weeks) / 8.0
	} else {
		durationFactor = 1.0 + (0.1 * float64(weeks-8) / 8.0)
	}

	// Base effect factors
	var endothelialEffect, inflammationEffect, coagulationEffect, progenitorEffect float64

	// Calculate base effects based on prescription type
	switch prescription.Name {
	case "Endothelial Health Protocol":
		endothelialEffect = 0.4 * intensityFactor * durationFactor
		inflammationEffect = 0.2 * intensityFactor * durationFactor
		coagulationEffect = 0.1 * intensityFactor * durationFactor
		progenitorEffect = 0.1 * intensityFactor * durationFactor

	case "Vascular Inflammation Reduction":
		endothelialEffect = 0.2 * intensityFactor * durationFactor
		inflammationEffect = 0.5 * intensityFactor * durationFactor
		coagulationEffect = 0.2 * intensityFactor * durationFactor
		progenitorEffect = 0.1 * intensityFactor * durationFactor

	case "Coagulation Profile Improvement":
		endothelialEffect = 0.2 * intensityFactor * durationFactor
		inflammationEffect = 0.2 * intensityFactor * durationFactor
		coagulationEffect = 0.5 * intensityFactor * durationFactor
		progenitorEffect = 0.1 * intensityFactor * durationFactor

	case "Progenitor Cell Mobilization":
		endothelialEffect = 0.3 * intensityFactor * durationFactor
		inflammationEffect = 0.1 * intensityFactor * durationFactor
		coagulationEffect = 0.1 * intensityFactor * durationFactor
		progenitorEffect = 0.6 * intensityFactor * durationFactor
	}

	// Adjust for baseline conditions (higher effect if pathological at baseline)
	if endothelialDysfunction {
		endothelialEffect *= 1.5
	}

	if baselineInflammation {
		inflammationEffect *= 1.5
	}

	// Calculate specific factor changes
	// Endothelial factors
	response.FactorChanges["Nitric Oxide"] = 15.0 + (40.0 * endothelialEffect)
	response.FactorChanges["VEGF"] = 10.0 + (30.0 * endothelialEffect) + (40.0 * progenitorEffect)
	response.FactorChanges["ET-1"] = -5.0 - (20.0 * endothelialEffect) // Negative is good (reduction)
	response.FactorChanges["Prostacyclin"] = 10.0 + (20.0 * endothelialEffect)

	// Inflammatory factors
	response.FactorChanges["IL-6 (acute)"] = 50.0 + (100.0 * intensityFactor)      // Acute increase is normal
	response.FactorChanges["IL-6 (chronic)"] = -10.0 - (20.0 * inflammationEffect) // Chronic decrease is beneficial
	response.FactorChanges["CRP"] = -5.0 - (25.0 * inflammationEffect)
	response.FactorChanges["TNF-α"] = -5.0 - (20.0 * inflammationEffect)
	response.FactorChanges["IL-10"] = 10.0 + (30.0 * inflammationEffect)  // Anti-inflammatory
	response.FactorChanges["IL-1Ra"] = 15.0 + (35.0 * inflammationEffect) // Anti-inflammatory

	// Coagulation factors
	response.FactorChanges["Fibrinogen"] = -5.0 - (15.0 * coagulationEffect)
	response.FactorChanges["D-dimer (acute)"] = 10.0 + (20.0 * intensityFactor)     // Acute increase is normal
	response.FactorChanges["D-dimer (chronic)"] = -5.0 - (10.0 * coagulationEffect) // Chronic decrease is beneficial
	response.FactorChanges["PAI-1"] = -10.0 - (30.0 * coagulationEffect)
	response.FactorChanges["tPA"] = 15.0 + (25.0 * coagulationEffect)

	// Progenitor/stem cell factors
	response.FactorChanges["SDF-1"] = 20.0 + (40.0 * progenitorEffect)
	response.FactorChanges["G-CSF"] = 15.0 + (35.0 * progenitorEffect)

	// Cell changes
	response.CellChanges["EPCs"] = 50.0 + (150.0 * progenitorEffect)
	response.CellChanges["CD34+ Cells"] = 30.0 + (120.0 * progenitorEffect)
	response.CellChanges["Monocytes (anti-inflammatory)"] = 10.0 + (30.0 * inflammationEffect)
	response.CellChanges["Regulatory T-cells"] = 5.0 + (20.0 * inflammationEffect)
	response.CellChanges["Neutrophil ROS production"] = -10.0 - (30.0 * inflammationEffect)
	response.CellChanges["Platelet aggregation"] = -5.0 - (25.0 * coagulationEffect)

	// Flow metrics
	response.FlowMetrics["Flow-mediated dilation"] = 5.0 + (25.0 * endothelialEffect)
	response.FlowMetrics["Arterial compliance"] = 3.0 + (17.0 * endothelialEffect)
	response.FlowMetrics["Pulse wave velocity"] = -2.0 - (13.0 * endothelialEffect) // Decrease is good
	response.FlowMetrics["Reactive hyperemia index"] = 5.0 + (20.0 * endothelialEffect)

	// Coagulation metrics
	response.CoagulationMetrics["Clotting time"] = 3.0 + (10.0 * coagulationEffect)
	response.CoagulationMetrics["Platelet aggregation threshold"] = 5.0 + (15.0 * coagulationEffect)
	response.CoagulationMetrics["Fibrinolytic capacity"] = 10.0 + (30.0 * coagulationEffect)
	response.CoagulationMetrics["Clot lysis time"] = -5.0 - (15.0 * coagulationEffect) // Decrease is good

	// Inflammation metrics
	response.InflammationMetrics["Systemic inflammation score"] = -5.0 - (25.0 * inflammationEffect)
	response.InflammationMetrics["Endothelial activation markers"] = -5.0 - (20.0 * endothelialEffect) // Decrease is good
	response.InflammationMetrics["Oxidative stress markers"] = -5.0 - (15.0 * (endothelialEffect + inflammationEffect) / 2)

	return response
}

// GetExerciseDuration calculates recommended exercise duration based on condition and patient factors
func GetExerciseDuration(prescription ExercisePrescription, ageYears int, conditionSeverity string) int {
	// Base duration from prescription
	duration := prescription.DurationMinutes

	// Adjust for age
	if ageYears > 60 {
		duration = int(float64(duration) * 0.9) // 10% shorter for older adults
	}
	if ageYears > 75 {
		duration = int(float64(duration) * 0.85) // Further reduction for elderly
	}

	// Adjust for condition severity
	switch conditionSeverity {
	case "Mild":
		// No change from base
	case "Moderate":
		duration = int(float64(duration) * 0.9)
	case "Severe":
		duration = int(float64(duration) * 0.7) // 30% shorter for severe conditions
		// Minimum acceptable duration
		if duration < 20 {
			duration = 20
		}
	}

	return duration
}

// CalculateTimeToEffect estimates weeks needed to achieve a specific benefit
func CalculateTimeToEffect(prescription ExercisePrescription, benefitTarget string, benefitMagnitude float64) float64 {
	// Default times based on typical adaptation rates
	baselineWeeks := 0.0

	switch benefitTarget {
	case "Endothelial function":
		baselineWeeks = 4.0
	case "Inflammation reduction":
		baselineWeeks = 6.0
	case "Coagulation profile":
		baselineWeeks = 5.0
	case "Progenitor cell count":
		baselineWeeks = 3.0
	default:
		baselineWeeks = 8.0
	}

	// Adjust for prescription intensity and frequency
	intensityFactor := float64(prescription.IntensityPercent) / 70.0 // Normalized to moderate intensity
	frequencyFactor := float64(prescription.FrequencyPerWeek) / 3.0  // Normalized to 3x/week

	// Adjust for magnitude of desired benefit
	magnitudeFactor := benefitMagnitude / 25.0 // Normalized to 25% improvement

	// Calculate adjusted timeframe
	adjustedWeeks := baselineWeeks * magnitudeFactor / (intensityFactor * frequencyFactor)

	// Ensure minimum timeframe
	if adjustedWeeks < 1.0 {
		adjustedWeeks = 1.0
	}

	return adjustedWeeks
}
