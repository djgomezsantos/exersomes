package immune

// ExercisePrescription defines parameters for immune-targeting exercise
type ExercisePrescription struct {
	Name                 string
	Description          string
	PrimaryType          string // "Aerobic", "HIIT", "Resistance", "Combined"
	IntensityPercent     int    // % of max HR, VO2max, or 1RM
	DurationMinutes      int
	FrequencyPerWeek     int
	IntervalStructure    string   // For HIIT protocols
	TargetImmuneFunction []string // Specific immune functions targeted
	TargetCytokines      []string // Cytokines modulated by this prescription
	BeneficialFor        []string // Conditions that may benefit
	ContraindicationsFor []string // Populations that should avoid this protocol
	TimeToEffect         struct {
		Acute   string // Immediate effects (0-24 hours)
		Chronic string // Long-term adaptations
	}
	RecoveryNeeds string // Recovery time needed between sessions
}

// ImmuneResponse represents the predicted changes in immune parameters
type ImmuneResponse struct {
	CytokineChanges         map[string]float64 // Changes in inflammatory markers
	CellCountChanges        map[string]float64 // Changes in immune cell populations
	CellFunctionChanges     map[string]float64 // Changes in immune cell function
	InflammatoryBalance     float64            // Pro-to-anti-inflammatory balance
	MucosalImmunity         float64            // Mucosal immunity strength
	InfectionRisk           float64            // Relative infection risk
	AntibodyResponse        float64            // Relative antibody response capacity
	TissueDamageRisk        float64            // Risk of exercise-induced tissue damage
	ImmunometabolismChanges map[string]float64 // Changes in immune cell metabolism
}

// Standard immune-targeting exercise prescriptions
var (
	AntiInflammatoryProtocol = ExercisePrescription{
		Name:              "Anti-inflammatory Protocol",
		Description:       "Moderate-intensity continuous training to reduce chronic inflammation",
		PrimaryType:       "Aerobic",
		IntensityPercent:  65,
		DurationMinutes:   45,
		FrequencyPerWeek:  4,
		IntervalStructure: "",
		TargetImmuneFunction: []string{
			"Reduce baseline inflammation",
			"Normalize cytokine balance",
			"Improve regulatory T cell function",
			"Shift macrophage phenotype to M2",
		},
		TargetCytokines: []string{
			"Decrease TNF-α",
			"Decrease IL-6 baseline",
			"Increase IL-10",
			"Increase IL-1RA",
		},
		BeneficialFor: []string{
			"Chronic low-grade inflammation",
			"Metabolic syndrome",
			"Rheumatic diseases (in stable phase)",
			"Cardiovascular disease prevention",
			"Obesity-related inflammation",
		},
		ContraindicationsFor: []string{
			"Acute infection",
			"Active autoimmune flare",
			"Post-surgical recovery (early phase)",
		},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "Temporary elevation then reduction within 24-48 hours",
			Chronic: "Significant reduction in baseline inflammation in 8-12 weeks",
		},
		RecoveryNeeds: "24 hours between sessions; avoid consecutive days if new to exercise",
	}

	ImmunoenhancementProtocol = ExercisePrescription{
		Name:              "Immune Enhancement Protocol",
		Description:       "Mixed-intensity training to boost immunity and surveillance",
		PrimaryType:       "Combined",
		IntensityPercent:  70,
		DurationMinutes:   40,
		FrequencyPerWeek:  3,
		IntervalStructure: "5-minute warm-up, 25 minutes moderate intensity, 5-minute HIIT, 5-minute cool-down",
		TargetImmuneFunction: []string{
			"Enhance natural killer cell activity",
			"Improve phagocytosis",
			"Increase immunosurveillance",
			"Boost mucosal immunity",
		},
		TargetCytokines: []string{
			"Optimize IL-6 response",
			"Increase IL-7",
			"Increase antimicrobial peptides",
		},
		BeneficialFor: []string{
			"Frequent respiratory infections",
			"Cancer prevention",
			"Aging immune system",
			"Recovery from illness",
			"Vaccine response",
		},
		ContraindicationsFor: []string{
			"Acute infection",
			"Severe immunodeficiency",
			"Uncontrolled autoimmunity",
		},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "Enhanced immune cell mobilization within hours",
			Chronic: "Improved immune surveillance and function in 4-8 weeks",
		},
		RecoveryNeeds: "48 hours between high-intensity components",
	}

	ImmunoregulationProtocol = ExercisePrescription{
		Name:              "Immunoregulation Protocol",
		Description:       "Regular moderate exercise to improve immune tolerance and self-regulation",
		PrimaryType:       "Aerobic",
		IntensityPercent:  60,
		DurationMinutes:   35,
		FrequencyPerWeek:  5,
		IntervalStructure: "",
		TargetImmuneFunction: []string{
			"Enhance regulatory T cell function",
			"Normalize self-tolerance",
			"Regulate dendritic cell phenotype",
			"Optimize immune signaling",
		},
		TargetCytokines: []string{
			"Increase IL-10",
			"Increase TGF-β",
			"Normalize IL-23/IL-17 axis",
		},
		BeneficialFor: []string{
			"Autoimmune conditions",
			"Allergic disorders",
			"Asthma",
			"Inflammatory bowel disease",
			"Systemic inflammation",
		},
		ContraindicationsFor: []string{
			"Acute disease flare",
			"Severe malnutrition",
			"Recent major surgery",
		},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "Temporary stress reduction within 1-2 hours",
			Chronic: "Improved immunoregulatory balance in 8-12 weeks",
		},
		RecoveryNeeds: "24 hours; consecutive days acceptable due to moderate intensity",
	}

	RecoveryImmunityProtocol = ExercisePrescription{
		Name:              "Recovery Immunity Protocol",
		Description:       "Low-intensity exercise to maintain immunity during heavy training periods",
		PrimaryType:       "Aerobic",
		IntensityPercent:  50,
		DurationMinutes:   30,
		FrequencyPerWeek:  2,
		IntervalStructure: "",
		TargetImmuneFunction: []string{
			"Prevent immunosuppression",
			"Maintain mucosal immunity",
			"Support lymphocyte function",
			"Avoid open window effect",
		},
		TargetCytokines: []string{
			"Limit stress hormone response",
			"Maintain sIgA levels",
			"Prevent excessive inflammatory signaling",
		},
		BeneficialFor: []string{
			"Athletes in intense training",
			"Recovery between competitions",
			"Prevention of upper respiratory tract infections",
			"Overtraining prevention",
		},
		ContraindicationsFor: []string{
			"None for target population",
		},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "Minimal stress on immune system",
			Chronic: "Maintained immune function during high training loads",
		},
		RecoveryNeeds: "Minimal; can be performed daily as active recovery",
	}
)

// GetPrescriptionByCondition returns appropriate immune-targeting exercise prescriptions
func GetPrescriptionByCondition(condition string) []ExercisePrescription {
	var prescriptions []ExercisePrescription

	switch condition {
	case "Chronic inflammation":
		prescriptions = append(prescriptions, AntiInflammatoryProtocol)
	case "Recurrent infections":
		prescriptions = append(prescriptions, ImmunoenhancementProtocol)
	case "Autoimmune disorder":
		prescriptions = append(prescriptions, ImmunoregulationProtocol)
	case "Athletic recovery":
		prescriptions = append(prescriptions, RecoveryImmunityProtocol)
	case "Obesity":
		prescriptions = append(prescriptions, AntiInflammatoryProtocol, ImmunoenhancementProtocol)
	case "Aging immune system":
		prescriptions = append(prescriptions, ImmunoenhancementProtocol, AntiInflammatoryProtocol)
	default:
		// Default general immune health prescription
		prescriptions = append(prescriptions, ImmunoenhancementProtocol)
	}

	return prescriptions
}

// PredictImmuneResponse estimates the effects of an exercise prescription
// on immune parameters based on the protocol and duration of adherence
func PredictImmuneResponse(prescription ExercisePrescription, timePoint string,
	chronicWeeks int, hasInflammation bool, isImmunocompromised bool) ImmuneResponse {

	// Initialize response
	response := ImmuneResponse{
		CytokineChanges:         make(map[string]float64),
		CellCountChanges:        make(map[string]float64),
		CellFunctionChanges:     make(map[string]float64),
		ImmunometabolismChanges: make(map[string]float64),
	}

	// Set baseline values
	response.InflammatoryBalance = 1.0 // 1.0 = balanced, >1.0 = pro-inflammatory
	response.MucosalImmunity = 1.0     // 1.0 = normal
	response.InfectionRisk = 1.0       // 1.0 = normal risk
	response.AntibodyResponse = 1.0    // 1.0 = normal response
	response.TissueDamageRisk = 1.0    // 1.0 = normal risk

	// Adjust baseline based on individual factors
	if hasInflammation {
		response.InflammatoryBalance = 1.4 // Higher pro-inflammatory state
		response.MucosalImmunity = 0.8     // Somewhat reduced
	}

	if isImmunocompromised {
		response.InfectionRisk = 1.5    // Higher baseline risk
		response.AntibodyResponse = 0.7 // Lower baseline response
	}

	// Calculate factors for determining response
	intensityFactor := float64(prescription.IntensityPercent) / 100.0
	durationFactor := 0.0
	if chronicWeeks > 0 {
		// Chronic adaptation factor (diminishing returns after 12 weeks)
		if chronicWeeks <= 12 {
			durationFactor = float64(chronicWeeks) / 12.0
		} else {
			durationFactor = 1.0 + (0.2 * float64(chronicWeeks-12) / 12.0)
			if durationFactor > 1.5 {
				durationFactor = 1.5 // Cap at 50% additional benefit
			}
		}
	}

	// Acute vs. chronic phase has different effects
	if timePoint == "Acute" {
		// Cytokine changes during/immediately after exercise

		// IL-6 shows large acute increase, especially with higher intensity
		response.CytokineChanges["IL-6"] = 3.0 + (10.0 * intensityFactor)

		// TNF-α shows small to moderate increase acutely
		response.CytokineChanges["TNF-α"] = 1.0 + (1.5 * intensityFactor)

		// IL-10 increases following IL-6
		response.CytokineChanges["IL-10"] = 1.0 + (2.0 * intensityFactor)

		// IL-1RA increases to counter inflammation
		response.CytokineChanges["IL-1RA"] = 1.0 + (3.0 * intensityFactor)

		// TGF-β shows modest acute changes
		response.CytokineChanges["TGF-β"] = 1.0 + (0.5 * intensityFactor)

		// Immune cell mobilization
		response.CellCountChanges["Neutrophils"] = 2.0 + (3.0 * intensityFactor)
		response.CellCountChanges["NK cells"] = 1.5 + (2.5 * intensityFactor)
		response.CellCountChanges["CD8+ T cells"] = 1.3 + (1.7 * intensityFactor)
		response.CellCountChanges["Monocytes"] = 1.2 + (1.3 * intensityFactor)

		// Cell function changes
		response.CellFunctionChanges["NK cell cytotoxicity"] = 1.3 + (0.7 * intensityFactor)
		response.CellFunctionChanges["Neutrophil ROS production"] = 1.2 + (0.8 * intensityFactor)
		response.CellFunctionChanges["T cell proliferation"] = 0.9 - (0.1 * intensityFactor) // Temporary decrease

		// Open window effect with very high intensity
		if intensityFactor > 0.8 {
			openWindowFactor := (intensityFactor - 0.8) * 5.0 // Scale factor for open window effect
			response.InfectionRisk = 1.0 + (0.5 * openWindowFactor)
			response.MucosalImmunity = 1.0 - (0.3 * openWindowFactor)
		}

		// Exercise metabolism changes
		response.ImmunometabolismChanges["Glycolysis"] = 1.5 + (1.5 * intensityFactor)
		response.ImmunometabolismChanges["Mitochondrial respiration"] = 1.2 + (0.8 * intensityFactor)
	} else {
		// Chronic adaptations after regular training

		// Calculate protocol-specific adaptation factors
		var inflammationReductionFactor, immunoenhancementFactor, regulationFactor float64

		switch prescription.Name {
		case "Anti-inflammatory Protocol":
			inflammationReductionFactor = 0.4 * durationFactor
			immunoenhancementFactor = 0.2 * durationFactor
			regulationFactor = 0.3 * durationFactor
		case "Immune Enhancement Protocol":
			inflammationReductionFactor = 0.2 * durationFactor
			immunoenhancementFactor = 0.5 * durationFactor
			regulationFactor = 0.2 * durationFactor
		case "Immunoregulation Protocol":
			inflammationReductionFactor = 0.3 * durationFactor
			immunoenhancementFactor = 0.1 * durationFactor
			regulationFactor = 0.5 * durationFactor
		case "Recovery Immunity Protocol":
			inflammationReductionFactor = 0.15 * durationFactor
			immunoenhancementFactor = 0.15 * durationFactor
			regulationFactor = 0.15 * durationFactor
		}

		// Apply chronic cytokine changes
		response.CytokineChanges["IL-6 baseline"] = 1.0 - (0.2 * inflammationReductionFactor)
		response.CytokineChanges["TNF-α baseline"] = 1.0 - (0.25 * inflammationReductionFactor)
		response.CytokineChanges["IL-10 baseline"] = 1.0 + (0.15 * regulationFactor)
		response.CytokineChanges["IL-1RA baseline"] = 1.0 + (0.2 * inflammationReductionFactor)
		response.CytokineChanges["TGF-β baseline"] = 1.0 + (0.1 * regulationFactor)

		// Cell count adaptations
		response.CellCountChanges["NK cells baseline"] = 1.0 + (0.2 * immunoenhancementFactor)
		response.CellCountChanges["T regulatory cells"] = 1.0 + (0.3 * regulationFactor)
		response.CellCountChanges["Senescent T cells"] = 1.0 - (0.15 * immunoenhancementFactor)

		// Cell function adaptations
		response.CellFunctionChanges["NK cell cytotoxicity"] = 1.0 + (0.3 * immunoenhancementFactor)
		response.CellFunctionChanges["Neutrophil chemotaxis"] = 1.0 + (0.25 * immunoenhancementFactor)
		response.CellFunctionChanges["Macrophage phagocytosis"] = 1.0 + (0.2 * immunoenhancementFactor)
		response.CellFunctionChanges["T cell differentiation balance"] = 1.0 + (0.3 * regulationFactor)

		// Overall immune parameters
		response.InflammatoryBalance = 1.0 - (0.2 * inflammationReductionFactor)
		if hasInflammation {
			response.InflammatoryBalance = 1.4 - (0.5 * inflammationReductionFactor)
			if response.InflammatoryBalance < 1.0 {
				response.InflammatoryBalance = 1.0 // Floor at healthy balance
			}
		}

		response.MucosalImmunity = 1.0 + (0.3 * immunoenhancementFactor)
		if isImmunocompromised {
			response.MucosalImmunity = 0.8 + (0.3 * immunoenhancementFactor)
		}

		response.InfectionRisk = 1.0 - (0.25 * immunoenhancementFactor)
		if isImmunocompromised {
			response.InfectionRisk = 1.5 - (0.5 * immunoenhancementFactor)
			if response.InfectionRisk < 1.0 {
				response.InfectionRisk = 1.0 // Floor at normal risk
			}
		}

		response.AntibodyResponse = 1.0 + (0.2 * immunoenhancementFactor)

		// Immune cell metabolism adaptations
		response.ImmunometabolismChanges["Mitochondrial biogenesis"] = 1.0 + (0.3 * durationFactor)
		response.ImmunometabolismChanges["Metabolic flexibility"] = 1.0 + (0.25 * durationFactor)
		response.ImmunometabolismChanges["ROS management"] = 1.0 + (0.2 * inflammationReductionFactor)
	}

	// Protocol-specific adjustments for tissue damage risk
	if prescription.IntensityPercent > 80 {
		response.TissueDamageRisk = 1.0 + (0.3 * (float64(prescription.IntensityPercent) - 80.0) / 20.0)
	} else {
		response.TissueDamageRisk = 0.9 // Slight protective effect of moderate exercise
	}

	return response
}

// GetCytokineTiming returns the temporal pattern of cytokine changes during and after exercise
func GetCytokineTiming(cytokine string) map[string]float64 {
	timings := make(map[string]float64)

	switch cytokine {
	case "IL-6":
		// IL-6 increases rapidly during exercise, peaks after, and normalizes in 24h
		timings["during_exercise"] = 3.0
		timings["immediate_post"] = 8.0
		timings["2hr_post"] = 5.0
		timings["6hr_post"] = 2.0
		timings["24hr_post"] = 1.2
	case "TNF-α":
		// TNF-α shows more modest increases
		timings["during_exercise"] = 1.5
		timings["immediate_post"] = 2.0
		timings["2hr_post"] = 1.8
		timings["6hr_post"] = 1.3
		timings["24hr_post"] = 1.0
	case "IL-10":
		// IL-10 increases later, following the pro-inflammatory response
		timings["during_exercise"] = 1.0
		timings["immediate_post"] = 2.0
		timings["2hr_post"] = 3.0
		timings["6hr_post"] = 2.0
		timings["24hr_post"] = 1.2
	case "IL-1RA":
		// IL-1RA increases to counter IL-1β
		timings["during_exercise"] = 1.5
		timings["immediate_post"] = 3.0
		timings["2hr_post"] = 4.0
		timings["6hr_post"] = 2.5
		timings["24hr_post"] = 1.3
	case "TGF-β":
		// TGF-β shows slower changes
		timings["during_exercise"] = 1.1
		timings["immediate_post"] = 1.2
		timings["2hr_post"] = 1.5
		timings["6hr_post"] = 1.5
		timings["24hr_post"] = 1.3
	}

	return timings
}

// TissueEffects returns the primary effects of exercise on specific immune tissues
func TissueEffects(tissue string, exerciseType string, chronicWeeks int) map[string]string {
	effects := make(map[string]string)

	switch tissue {
	case "Thymus":
		effects["Acute"] = "Temporary stress response with potential cortisol-induced apoptosis of immature thymocytes"
		if chronicWeeks > 8 {
			effects["Chronic"] = "Potential slowing of thymic involution, preserved T cell output, and maintained T cell repertoire diversity"
		} else {
			effects["Chronic"] = "Beginning adaptations toward preserving thymic function"
		}
		effects["Key Factors"] = "IL-7, IGF-1, GH, Thymulin, lower cortisol sensitivity"

	case "Bone Marrow":
		effects["Acute"] = "Mobilization of leukocytes, increased hematopoiesis"
		effects["Chronic"] = "Enhanced hematopoietic stem cell maintenance, improved myeloid/lymphoid balance"
		effects["Key Factors"] = "G-CSF, IL-7, reduced oxidative stress, enhanced blood flow"

	case "Lymph Nodes":
		effects["Acute"] = "Increased lymphatic flow, cell trafficking, and antigen presentation"
		effects["Chronic"] = "Enhanced germinal center function, improved B and T cell interactions, better antibody affinity maturation"
		effects["Key Factors"] = "Increased lymph flow, DC migration, follicular T cell function"

	case "Spleen":
		effects["Acute"] = "Splenic contraction, leukocyte mobilization, especially NK and CD8+ T cells"
		if exerciseType == "HIIT" || exerciseType == "Resistance" {
			effects["Acute"] += ", pronounced catecholamine-driven cell release"
		}
		effects["Chronic"] = "Improved splenic reserve, enhanced filtering capacity, better pathogen clearance"
		effects["Key Factors"] = "Catecholamines, improved blood flow, enhanced macrophage function"

	case "MALT (Mucosal-Associated Lymphoid Tissue)":
		effects["Acute"] = "Temporary stress on mucosal barriers, fluctuations in sIgA"
		if exerciseType == "HIIT" && chronicWeeks < 4 {
			effects["Acute"] += ", potential temporary decrease in mucosal protection"
		}
		effects["Chronic"] = "Enhanced sIgA production, improved epithelial barrier function, balanced mucosal immunity"
		effects["Key Factors"] = "sIgA, antimicrobial peptides, regulatory T cells, microbiome stability"
	}

	return effects
}

// CalculateExerciseInflammationCurve models how inflammatory responses change with exercise intensity
func CalculateExerciseInflammationCurve() map[string]map[string]float64 {
	// Mapping exercise intensity to inflammatory responses
	curve := make(map[string]map[string]float64)

	// Initialize intensity levels
	intensities := []string{"Sedentary", "Low", "Moderate", "Vigorous", "Very High", "Extreme"}

	for _, intensity := range intensities {
		curve[intensity] = make(map[string]float64)
	}

	// Acute responses
	curve["Sedentary"]["IL-6_acute"] = 1.0
	curve["Low"]["IL-6_acute"] = 2.0
	curve["Moderate"]["IL-6_acute"] = 5.0
	curve["Vigorous"]["IL-6_acute"] = 10.0
	curve["Very High"]["IL-6_acute"] = 18.0
	curve["Extreme"]["IL-6_acute"] = 25.0

	curve["Sedentary"]["TNF_acute"] = 1.0
	curve["Low"]["TNF_acute"] = 1.2
	curve["Moderate"]["TNF_acute"] = 1.5
	curve["Vigorous"]["TNF_acute"] = 2.0
	curve["Very High"]["TNF_acute"] = 3.0
	curve["Extreme"]["TNF_acute"] = 4.0

	curve["Sedentary"]["IL-10_acute"] = 1.0
	curve["Low"]["IL-10_acute"] = 1.2
	curve["Moderate"]["IL-10_acute"] = 1.8
	curve["Vigorous"]["IL-10_acute"] = 3.0
	curve["Very High"]["IL-10_acute"] = 4.5
	curve["Extreme"]["IL-10_acute"] = 5.0

	// Chronic baseline adaptations
	curve["Sedentary"]["IL-6_chronic"] = 1.0
	curve["Low"]["IL-6_chronic"] = 0.95
	curve["Moderate"]["IL-6_chronic"] = 0.8
	curve["Vigorous"]["IL-6_chronic"] = 0.7
	curve["Very High"]["IL-6_chronic"] = 0.7
	curve["Extreme"]["IL-6_chronic"] = 0.9 // U-shaped curve with potential negative effects at extreme levels

	curve["Sedentary"]["TNF_chronic"] = 1.0
	curve["Low"]["TNF_chronic"] = 0.9
	curve["Moderate"]["TNF_chronic"] = 0.8
	curve["Vigorous"]["TNF_chronic"] = 0.7
	curve["Very High"]["TNF_chronic"] = 0.75
	curve["Extreme"]["TNF_chronic"] = 0.9 // U-shaped curve

	curve["Sedentary"]["IL-10_chronic"] = 1.0
	curve["Low"]["IL-10_chronic"] = 1.1
	curve["Moderate"]["IL-10_chronic"] = 1.2
	curve["Vigorous"]["IL-10_chronic"] = 1.3
	curve["Very High"]["IL-10_chronic"] = 1.25
	curve["Extreme"]["IL-10_chronic"] = 1.1 // Diminishing returns

	// Composite measure: Anti-to-pro inflammatory ratio
	curve["Sedentary"]["Anti_pro_ratio"] = 1.0
	curve["Low"]["Anti_pro_ratio"] = 1.2
	curve["Moderate"]["Anti_pro_ratio"] = 1.5
	curve["Vigorous"]["Anti_pro_ratio"] = 1.8
	curve["Very High"]["Anti_pro_ratio"] = 1.7
	curve["Extreme"]["Anti_pro_ratio"] = 1.2 // U-shaped curve

	return curve
}
