package neural

// ExercisePrescription defines exercise parameters targeting neurological function
type ExercisePrescription struct {
	Name                 string
	Description          string
	PrimaryType          string // "Aerobic", "Resistance", "Mixed", "Skill", "Balance"
	IntensityPercent     int    // % of max HR, VO2max, or 1RM
	DurationMinutes      int
	FrequencyPerWeek     int
	TargetSystem         string   // "CNS", "PNS", "Both"
	TargetDomain         []string // "Cognitive", "Motor", "Sensory", "Autonomic"
	SpecificTarget       []string // Specific brain areas or nerve types
	TargetNeurokines     []string // Target neurochemical factors
	TargetReceptors      []string // Target neural receptors
	ExpectedBenefits     []string
	MechanismsOfAction   []string // How the prescription works
	ValidatedPopulations []string // Groups with evidence of efficacy
	ContraindicationsFor []string
	TimeToEffect         struct {
		Acute   string // Acute effects
		Chronic string // Long-term adaptations
	}
}

// NeuralResponse represents the predicted changes in neurological parameters
type NeuralResponse struct {
	NeurokineChanges       map[string]float64 // Changes in brain-derived factors
	ReceptorChanges        map[string]float64 // Changes in receptor sensitivity
	StructuralChanges      map[string]float64 // Structural brain/nerve changes
	FunctionalChanges      map[string]float64 // Changes in brain/nerve function
	CognitiveDomainChanges map[string]float64 // Changes in cognitive domains
}

// Standard neural-targeting exercise prescriptions
var (
	ExecutiveFunctionEnhancement = ExercisePrescription{
		Name:             "Executive Function Enhancement",
		Description:      "High-intensity interval aerobic protocol for improving prefrontal cortex function",
		PrimaryType:      "Aerobic",
		IntensityPercent: 80,
		DurationMinutes:  25,
		FrequencyPerWeek: 3,
		TargetSystem:     "CNS",
		TargetDomain:     []string{"Cognitive"},
		SpecificTarget:   []string{"Prefrontal Cortex", "DLPFC", "ACC"},
		TargetNeurokines: []string{"BDNF", "IGF-1", "Irisin", "Cathepsin B"},
		TargetReceptors:  []string{"TrkB", "IGF-1R", "NMDA Receptors"},
		ExpectedBenefits: []string{
			"Improved working memory",
			"Enhanced cognitive flexibility",
			"Better inhibitory control",
			"Increased attention span",
		},
		MechanismsOfAction: []string{
			"Increased BDNF signaling",
			"Enhanced synaptic plasticity",
			"Improved cerebral blood flow",
			"Increased prefrontal activation",
		},
		ValidatedPopulations: []string{
			"Healthy adults",
			"Older adults with mild cognitive decline",
			"ADHD",
		},
		ContraindicationsFor: []string{
			"Uncontrolled epilepsy",
			"Recent concussion",
		},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "Immediate improvements for 1-3 hours post-exercise",
			Chronic: "Structural changes within 6-8 weeks",
		},
	}

	HippocampalMemoryProtocol = ExercisePrescription{
		Name:             "Hippocampal Memory Enhancement",
		Description:      "Moderate aerobic protocol targeting hippocampal function and memory",
		PrimaryType:      "Aerobic",
		IntensityPercent: 70,
		DurationMinutes:  45,
		FrequencyPerWeek: 4,
		TargetSystem:     "CNS",
		TargetDomain:     []string{"Cognitive"},
		SpecificTarget:   []string{"Hippocampus", "Entorhinal Cortex"},
		TargetNeurokines: []string{"BDNF", "VEGF", "IGF-1"},
		TargetReceptors:  []string{"TrkB", "NMDA Receptors"},
		ExpectedBenefits: []string{
			"Enhanced episodic memory",
			"Improved spatial navigation",
			"Increased hippocampal volume",
			"Protection against age-related memory decline",
		},
		MechanismsOfAction: []string{
			"Adult hippocampal neurogenesis",
			"Increased dendritic complexity",
			"Enhanced long-term potentiation",
			"Improved cerebral blood flow",
		},
		ValidatedPopulations: []string{
			"Healthy adults",
			"Older adults",
			"Early Alzheimer's disease",
			"Mild cognitive impairment",
		},
		ContraindicationsFor: []string{
			"Advanced dementia",
		},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "Modest improvements within hours",
			Chronic: "Significant improvements in 8-12 weeks",
		},
	}

	PeripheralNerveRegeneration = ExercisePrescription{
		Name:             "Peripheral Nerve Regeneration Protocol",
		Description:      "Combined exercise approach to enhance peripheral nerve regeneration and function",
		PrimaryType:      "Mixed",
		IntensityPercent: 60,
		DurationMinutes:  30,
		FrequencyPerWeek: 3,
		TargetSystem:     "PNS",
		TargetDomain:     []string{"Sensory", "Motor"},
		SpecificTarget:   []string{"Dorsal Root Ganglia", "Peripheral Nerves", "Neuromuscular Junction"},
		TargetNeurokines: []string{"BDNF", "NGF", "NT-3", "IGF-1"},
		TargetReceptors:  []string{"TrkA", "TrkB", "p75NTR"},
		ExpectedBenefits: []string{
			"Enhanced nerve regeneration after injury",
			"Improved nerve conduction velocity",
			"Reduced neuropathic pain",
			"Enhanced sensorimotor function",
		},
		MechanismsOfAction: []string{
			"Increased neurotrophin signaling",
			"Schwann cell activation",
			"Reduced neuroinflammation",
			"Enhanced axonal transport",
		},
		ValidatedPopulations: []string{
			"Diabetic neuropathy",
			"Chemotherapy-induced neuropathy",
			"Nerve compression syndromes",
			"Post-surgical nerve recovery",
		},
		ContraindicationsFor: []string{
			"Acute nerve injury (first 72 hours)",
			"Unstable neurological condition",
		},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "Minor symptom relief within days",
			Chronic: "Significant improvements in 12-16 weeks",
		},
	}

	MotorLearningEnhancement = ExercisePrescription{
		Name:             "Motor Learning and Skill Acquisition",
		Description:      "Specialized protocol combining skill training with aerobic exercise for motor learning",
		PrimaryType:      "Skill",
		IntensityPercent: 65,
		DurationMinutes:  40,
		FrequencyPerWeek: 4,
		TargetSystem:     "Both",
		TargetDomain:     []string{"Motor", "Cognitive"},
		SpecificTarget:   []string{"Motor Cortex", "Cerebellum", "Basal Ganglia", "Neuromuscular Junction"},
		TargetNeurokines: []string{"BDNF", "GDNF", "IGF-1"},
		TargetReceptors:  []string{"TrkB", "AMPA Receptors", "Dopamine Receptors"},
		ExpectedBenefits: []string{
			"Accelerated motor skill acquisition",
			"Enhanced motor memory consolidation",
			"Improved movement precision",
			"Better coordination and balance",
		},
		MechanismsOfAction: []string{
			"Optimized priming of motor circuits",
			"Enhanced neuroplasticity in motor areas",
			"Strengthened synaptic connections",
			"Improved cerebellar function",
		},
		ValidatedPopulations: []string{
			"Athletes",
			"Rehabilitation patients",
			"Stroke recovery",
			"Parkinson's disease",
		},
		ContraindicationsFor: []string{
			"Severe movement disorders",
			"Uncontrolled seizures",
		},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "Enhanced skill learning within single session",
			Chronic: "Consolidated motor skills in 4-6 weeks",
		},
	}

	MoodRegulationProtocol = ExercisePrescription{
		Name:             "Mood Regulation and Stress Resilience",
		Description:      "Exercise protocol targeting emotional regulation and stress response systems",
		PrimaryType:      "Aerobic",
		IntensityPercent: 65,
		DurationMinutes:  35,
		FrequencyPerWeek: 5,
		TargetSystem:     "CNS",
		TargetDomain:     []string{"Cognitive", "Autonomic"},
		SpecificTarget:   []string{"Prefrontal Cortex", "Hippocampus", "Amygdala", "HPA Axis"},
		TargetNeurokines: []string{"BDNF", "β-endorphin", "Serotonin", "GABA"},
		TargetReceptors:  []string{"TrkB", "Serotonin Receptors", "Opioid Receptors", "GABA Receptors"},
		ExpectedBenefits: []string{
			"Reduced anxiety and depressive symptoms",
			"Enhanced stress resilience",
			"Improved emotional regulation",
			"Better sleep quality",
		},
		MechanismsOfAction: []string{
			"Normalized HPA axis function",
			"Enhanced serotonergic signaling",
			"Increased GABA release",
			"Endorphin-mediated analgesia and euphoria",
		},
		ValidatedPopulations: []string{
			"Mild to moderate depression",
			"Generalized anxiety",
			"Stress-related disorders",
			"PTSD",
		},
		ContraindicationsFor: []string{
			"Recent manic episode",
			"Exercise addiction risk",
		},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "Mood elevation within 30-60 minutes",
			Chronic: "Sustained mood improvements in 3-4 weeks",
		},
	}
)

// GetPrescriptionByTarget returns appropriate exercise prescriptions for specific neural targets
func GetPrescriptionByTarget(target string) []ExercisePrescription {
	var prescriptions []ExercisePrescription

	switch target {
	case "Executive Function":
		prescriptions = append(prescriptions, ExecutiveFunctionEnhancement)
	case "Memory":
		prescriptions = append(prescriptions, HippocampalMemoryProtocol)
	case "Peripheral Nerves":
		prescriptions = append(prescriptions, PeripheralNerveRegeneration)
	case "Motor Skills":
		prescriptions = append(prescriptions, MotorLearningEnhancement)
	case "Mood":
		prescriptions = append(prescriptions, MoodRegulationProtocol)
	case "Cognitive Function":
		prescriptions = append(prescriptions, ExecutiveFunctionEnhancement, HippocampalMemoryProtocol)
	default:
		// Return a general recommendation if target not specified
		prescriptions = append(prescriptions, ExecutiveFunctionEnhancement)
	}

	return prescriptions
}

// PredictNeuralResponse estimates the effects of an exercise prescription
// on neural parameters based on the protocol and individual factors
func PredictNeuralResponse(prescription ExercisePrescription, weeks int,
	age int, hasNeurologicalCondition bool) NeuralResponse {

	// Initialize response
	response := NeuralResponse{
		NeurokineChanges:       make(map[string]float64),
		ReceptorChanges:        make(map[string]float64),
		StructuralChanges:      make(map[string]float64),
		FunctionalChanges:      make(map[string]float64),
		CognitiveDomainChanges: make(map[string]float64),
	}

	// Calculate baseline factors
	intensityFactor := float64(prescription.IntensityPercent) / 100.0
	frequencyFactor := float64(prescription.FrequencyPerWeek) / 3.0 // Normalized to 3x/week

	// Calculate duration factor (diminishing returns after 12 weeks)
	durationFactor := 0.0
	if weeks <= 12 {
		durationFactor = float64(weeks) / 12.0
	} else {
		durationFactor = 1.0 + (0.1 * float64(weeks-12) / 12.0)
	}

	// Age adjustment factor (neural plasticity decreases with age)
	ageFactor := 1.0
	if age > 30 {
		ageFactor = 1.0 - (float64(age-30) / 100.0)
		if ageFactor < 0.5 { // Floor at 50% of young adult response
			ageFactor = 0.5
		}
	}

	// Condition adjustment (greater potential improvement if baseline is impaired)
	conditionFactor := 1.0
	if hasNeurologicalCondition {
		conditionFactor = 1.3 // 30% more improvement potential
	}

	// Calculate combined effect factor
	effectFactor := intensityFactor * frequencyFactor * durationFactor * ageFactor * conditionFactor

	// Calculate neurokine changes based on prescription type
	switch prescription.Name {
	case "Executive Function Enhancement":
		response.NeurokineChanges["BDNF"] = 30.0 + (70.0 * effectFactor)
		response.NeurokineChanges["IGF-1"] = 15.0 + (35.0 * effectFactor)
		response.NeurokineChanges["Irisin"] = 20.0 + (50.0 * effectFactor)

		response.ReceptorChanges["TrkB sensitivity"] = 15.0 + (25.0 * effectFactor)
		response.ReceptorChanges["NMDA receptor function"] = 10.0 + (20.0 * effectFactor)

		response.StructuralChanges["Prefrontal gray matter volume"] = 1.0 + (3.0 * effectFactor)
		response.StructuralChanges["Prefrontal connectivity"] = 5.0 + (15.0 * effectFactor)

		response.FunctionalChanges["Prefrontal activation"] = 10.0 + (25.0 * effectFactor)
		response.FunctionalChanges["Default mode network efficiency"] = 5.0 + (20.0 * effectFactor)

		response.CognitiveDomainChanges["Working memory"] = 10.0 + (25.0 * effectFactor)
		response.CognitiveDomainChanges["Cognitive control"] = 15.0 + (25.0 * effectFactor)
		response.CognitiveDomainChanges["Attention"] = 10.0 + (30.0 * effectFactor)

	case "Hippocampal Memory Protocol":
		response.NeurokineChanges["BDNF"] = 25.0 + (55.0 * effectFactor)
		response.NeurokineChanges["VEGF"] = 15.0 + (35.0 * effectFactor)
		response.NeurokineChanges["IGF-1"] = 10.0 + (30.0 * effectFactor)

		response.ReceptorChanges["TrkB sensitivity"] = 15.0 + (25.0 * effectFactor)
		response.ReceptorChanges["NMDA receptor function"] = 15.0 + (25.0 * effectFactor)

		response.StructuralChanges["Hippocampal volume"] = 2.0 + (4.0 * effectFactor)
		response.StructuralChanges["Dendritic complexity"] = 5.0 + (15.0 * effectFactor)
		response.StructuralChanges["Neurogenesis"] = 30.0 + (70.0 * effectFactor)

		response.FunctionalChanges["Hippocampal activation"] = 15.0 + (25.0 * effectFactor)
		response.FunctionalChanges["Hippocampal-PFC connectivity"] = 10.0 + (20.0 * effectFactor)

		response.CognitiveDomainChanges["Episodic memory"] = 15.0 + (25.0 * effectFactor)
		response.CognitiveDomainChanges["Spatial memory"] = 10.0 + (30.0 * effectFactor)

	case "Peripheral Nerve Regeneration Protocol":
		response.NeurokineChanges["BDNF"] = 20.0 + (40.0 * effectFactor)
		response.NeurokineChanges["NGF"] = 25.0 + (45.0 * effectFactor)
		response.NeurokineChanges["NT-3"] = 15.0 + (35.0 * effectFactor)

		response.ReceptorChanges["TrkA sensitivity"] = 15.0 + (25.0 * effectFactor)
		response.ReceptorChanges["TrkB sensitivity"] = 10.0 + (20.0 * effectFactor)
		response.ReceptorChanges["p75NTR function"] = 15.0 + (25.0 * effectFactor)

		response.StructuralChanges["Axon diameter"] = 5.0 + (15.0 * effectFactor)
		response.StructuralChanges["Myelin thickness"] = 5.0 + (10.0 * effectFactor)
		response.StructuralChanges["DRG neuron survival"] = 10.0 + (20.0 * effectFactor)

		response.FunctionalChanges["Nerve conduction velocity"] = 5.0 + (15.0 * effectFactor)
		response.FunctionalChanges["Sensory threshold"] = 10.0 + (20.0 * effectFactor)
		response.FunctionalChanges["Motor unit recruitment"] = 10.0 + (20.0 * effectFactor)

	case "Motor Learning and Skill Acquisition":
		response.NeurokineChanges["BDNF"] = 25.0 + (45.0 * effectFactor)
		response.NeurokineChanges["GDNF"] = 15.0 + (35.0 * effectFactor)

		response.ReceptorChanges["AMPA receptor trafficking"] = 20.0 + (40.0 * effectFactor)
		response.ReceptorChanges["Dopamine receptor sensitivity"] = 15.0 + (25.0 * effectFactor)

		response.StructuralChanges["Motor cortex map expansion"] = 10.0 + (30.0 * effectFactor)
		response.StructuralChanges["Cerebellar synaptic density"] = 5.0 + (15.0 * effectFactor)

		response.FunctionalChanges["Motor cortex activation efficiency"] = 15.0 + (35.0 * effectFactor)
		response.FunctionalChanges["Cerebellar recruitment"] = 10.0 + (30.0 * effectFactor)
		response.FunctionalChanges["Motor sequence chunking"] = 15.0 + (45.0 * effectFactor)

		response.CognitiveDomainChanges["Procedural memory"] = 20.0 + (40.0 * effectFactor)
		response.CognitiveDomainChanges["Sequence learning"] = 15.0 + (45.0 * effectFactor)

	case "Mood Regulation and Stress Resilience":
		response.NeurokineChanges["BDNF"] = 20.0 + (40.0 * effectFactor)
		response.NeurokineChanges["β-endorphin"] = 30.0 + (70.0 * effectFactor)
		response.NeurokineChanges["Serotonin"] = 15.0 + (35.0 * effectFactor)
		response.NeurokineChanges["GABA"] = 10.0 + (30.0 * effectFactor)

		response.ReceptorChanges["Serotonin receptor sensitivity"] = 10.0 + (30.0 * effectFactor)
		response.ReceptorChanges["Opioid receptor binding"] = 15.0 + (25.0 * effectFactor)
		response.ReceptorChanges["Glucocorticoid receptor sensitivity"] = 15.0 + (25.0 * effectFactor)

		response.StructuralChanges["PFC gray matter volume"] = 1.0 + (3.0 * effectFactor)
		response.StructuralChanges["Hippocampal volume"] = 1.0 + (3.0 * effectFactor)
		response.StructuralChanges["Amygdala reactivity"] = -5.0 - (15.0 * effectFactor) // Decrease is beneficial

		response.FunctionalChanges["HPA axis reactivity"] = -10.0 - (30.0 * effectFactor) // Decrease is beneficial
		response.FunctionalChanges["PFC-amygdala coupling"] = 15.0 + (25.0 * effectFactor)

		response.CognitiveDomainChanges["Emotional regulation"] = 15.0 + (35.0 * effectFactor)
		response.CognitiveDomainChanges["Stress resilience"] = 20.0 + (40.0 * effectFactor)
		response.CognitiveDomainChanges["Anxiety symptoms"] = -15.0 - (35.0 * effectFactor) // Decrease is beneficial
	}

	return response
}

// CalculateOptimalPrescription customizes a neural exercise prescription based on
// individual factors, including age, fitness level, and therapeutic goals
func CalculateOptimalPrescription(basePrescription ExercisePrescription,
	ageYears int, fitnesLevel string, neurologicalCondition string) ExercisePrescription {

	// Start with the base prescription
	customized := basePrescription

	// Adjust duration based on age and fitness
	if ageYears > 60 || fitnesLevel == "Low" {
		// Reduce duration for older or less fit individuals
		customized.DurationMinutes = int(float64(customized.DurationMinutes) * 0.8)
		if customized.DurationMinutes < 20 { // Ensure minimum effective duration
			customized.DurationMinutes = 20
		}
	} else if fitnesLevel == "High" {
		// Increase duration for fitter individuals
		customized.DurationMinutes = int(float64(customized.DurationMinutes) * 1.2)
	}

	// Adjust intensity based on age and fitness
	if ageYears > 60 {
		customized.IntensityPercent -= 10
	} else if ageYears < 30 && fitnesLevel == "High" {
		customized.IntensityPercent += 5
	}

	if fitnesLevel == "Low" {
		customized.IntensityPercent -= 10
	} else if fitnesLevel == "High" {
		customized.IntensityPercent += 5
	}

	// Cap intensity within reasonable bounds
	if customized.IntensityPercent < 50 {
		customized.IntensityPercent = 50
	} else if customized.IntensityPercent > 90 {
		customized.IntensityPercent = 90
	}

	// Adjust frequency based on condition
	if neurologicalCondition != "" {
		// Increase frequency for therapeutic applications
		customized.FrequencyPerWeek += 1
		if customized.FrequencyPerWeek > 5 { // Cap at 5 days per week
			customized.FrequencyPerWeek = 5
		}
	}

	return customized
}

// TimeToImprovement calculates expected weeks until meaningful improvement
// for a specific neural function based on the prescription and individual factors
func TimeToImprovement(prescription ExercisePrescription, targetFunction string,
	ageYears int, hasNeurologicalCondition bool) int {

	// Baseline weeks by neural function
	baselineWeeks := 0

	switch targetFunction {
	case "Executive function":
		baselineWeeks = 6
	case "Memory":
		baselineWeeks = 8
	case "Mood":
		baselineWeeks = 4
	case "Peripheral nerve function":
		baselineWeeks = 12
	case "Motor learning":
		baselineWeeks = 5
	default:
		baselineWeeks = 8
	}

	// Adjust for prescription intensity and frequency
	intensityFactor := float64(prescription.IntensityPercent) / 70.0 // Normalized to 70%
	frequencyFactor := float64(prescription.FrequencyPerWeek) / 3.0  // Normalized to 3x/week

	// Age adjustment (neural adaptations slower with age)
	ageFactor := 1.0
	if ageYears > 40 {
		ageFactor = 1.0 + (float64(ageYears-40) / 100.0)
	}

	// Calculate adjusted timeframe
	adjustedWeeks := float64(baselineWeeks) * ageFactor / (intensityFactor * frequencyFactor)

	// If neurological condition is present, may need longer for improvement
	if hasNeurologicalCondition {
		adjustedWeeks *= 1.2
	}

	return int(adjustedWeeks)
}
