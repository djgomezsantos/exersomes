package bone

// ExercisePrescription defines parameters for bone-targeted exercise
type ExercisePrescription struct {
	Name              string
	Description       string
	PrimaryType       string // "Resistance", "Impact", "Combined", "Plyometric"
	IntensityPercent  int    // % of max or 1RM for resistance
	LoadMagnitude     string // "Low", "Moderate", "High" for impact exercises
	DurationMinutes   int
	FrequencyPerWeek  int
	TargetOsteokines  []string // Target bone-derived factors
	TargetCells       []string // Target bone cells
	ExpectedBenefits  []string
	MechanicalEffects []string // Effects on bone stress, strain, etc.
	IndicationsFor    []string
	Contraindications []string
	TimeToEffect      struct {
		Acute   string // Immediate effects
		Chronic string // Long-term adaptations
	}
}

// BoneResponse represents the predicted changes in bone parameters
type BoneResponse struct {
	OsteokineChanges    map[string]float64 // Changes in bone-derived factors
	CellActivityChanges map[string]float64 // Changes in bone cell activity
	StructuralMetrics   map[string]float64 // Bone density, architecture parameters
	FormationMetrics    map[string]float64 // Bone formation markers
	ResorptionMetrics   map[string]float64 // Bone resorption markers
}

// Standard bone-targeting exercise prescriptions
var (
	BoneMineralDensity = ExercisePrescription{
		Name:             "Bone Mineral Density Protocol",
		Description:      "Progressive resistance training to enhance bone mineral density",
		PrimaryType:      "Resistance",
		IntensityPercent: 75,
		LoadMagnitude:    "Moderate",
		DurationMinutes:  40,
		FrequencyPerWeek: 3,
		TargetOsteokines: []string{"Osteocalcin", "Sclerostin", "Osteopontin", "RANKL/OPG"},
		TargetCells:      []string{"Osteoblasts", "Osteocytes"},
		ExpectedBenefits: []string{
			"Increased bone mineral density",
			"Enhanced bone microarchitecture",
			"Improved bone strength",
			"Reduced fracture risk",
		},
		MechanicalEffects: []string{
			"Increased bone strain",
			"Enhanced mechanotransduction",
			"Improved bone cross-sectional area",
		},
		IndicationsFor: []string{
			"Osteopenia",
			"Osteoporosis",
			"Age-related bone loss",
			"Disuse osteoporosis",
		},
		Contraindications: []string{
			"Recent fracture",
			"Severe osteoporosis with vertebral fractures",
			"Acute bone metastases",
		},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "Osteokine changes within 0.5-2 hours",
			Chronic: "Structural changes in 3-6 months",
		},
	}

	OsteocyteActivation = ExercisePrescription{
		Name:             "Osteocyte Network Stimulation",
		Description:      "Impact exercise protocol to activate osteocyte networks",
		PrimaryType:      "Impact",
		IntensityPercent: 65,
		LoadMagnitude:    "Variable",
		DurationMinutes:  30,
		FrequencyPerWeek: 4,
		TargetOsteokines: []string{"Sclerostin", "FGF23", "PGE2", "DKK1", "DMP1"},
		TargetCells:      []string{"Osteocytes", "Bone Lining Cells"},
		ExpectedBenefits: []string{
			"Enhanced mechanosensing",
			"Improved osteocyte viability",
			"Optimized bone remodeling",
			"Enhanced bone material properties",
		},
		MechanicalEffects: []string{
			"Fluid flow in lacuno-canalicular network",
			"Dynamic strain patterns",
			"Enhanced bone interstitial fluid movement",
		},
		IndicationsFor: []string{
			"Age-related osteocyte dysfunction",
			"Impaired mechanotransduction",
			"Disuse-related bone loss",
			"Early osteoporosis",
		},
		Contraindications: []string{
			"Severe arthritis",
			"Acute joint inflammation",
			"Recent lower extremity fracture",
		},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "Signaling changes within 0.5-1 hour",
			Chronic: "Network improvements in 6-12 weeks",
		},
	}

	BoneRemodeling = ExercisePrescription{
		Name:             "Bone Remodeling Optimization",
		Description:      "Combined loading protocol to optimize bone turnover",
		PrimaryType:      "Combined",
		IntensityPercent: 70,
		LoadMagnitude:    "Moderate",
		DurationMinutes:  45,
		FrequencyPerWeek: 3,
		TargetOsteokines: []string{"RANKL", "OPG", "Osteocalcin", "TGF-β"},
		TargetCells:      []string{"Osteoblasts", "Osteoclasts", "Osteocytes"},
		ExpectedBenefits: []string{
			"Balanced bone remodeling",
			"Enhanced bone quality",
			"Improved mineralization",
			"Optimized collagen matrix",
		},
		MechanicalEffects: []string{
			"Targeted microdamage repair",
			"Enhanced bone material properties",
			"Improved microarchitecture",
		},
		IndicationsFor: []string{
			"Dysregulated bone turnover",
			"Metabolic bone diseases",
			"Secondary osteoporosis",
			"Recovery from immobilization",
		},
		Contraindications: []string{
			"Paget's disease (active phase)",
			"High-turnover bone disorders",
			"Recent bisphosphonate treatment",
		},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "Turnover marker changes within 24-48 hours",
			Chronic: "Remodeling optimization in 3-4 months",
		},
	}

	BoneAnabolism = ExercisePrescription{
		Name:             "Osteogenic Loading Protocol",
		Description:      "High-intensity, brief loading to maximize bone formation",
		PrimaryType:      "Plyometric",
		IntensityPercent: 85,
		LoadMagnitude:    "High",
		DurationMinutes:  20,
		FrequencyPerWeek: 2,
		TargetOsteokines: []string{"IGF-1", "BMP-2", "Wnt ligands", "Osteocalcin"},
		TargetCells:      []string{"Osteoprogenitors", "Osteoblasts", "MSCs"},
		ExpectedBenefits: []string{
			"Stimulated bone formation",
			"Recruited osteoprogenitors",
			"Enhanced periosteal expansion",
			"Improved bone geometry",
		},
		MechanicalEffects: []string{
			"High strain rates",
			"Peak compressive forces",
			"Enhanced mechanotransduction",
		},
		IndicationsFor: []string{
			"Stable osteopenia",
			"Athletic bone strengthening",
			"Post-fracture recovery phase",
			"Spaceflight-induced bone loss",
		},
		Contraindications: []string{
			"Unstable fractures",
			"Severe osteoporosis",
			"Joint instability",
			"Balance disorders",
		},
		TimeToEffect: struct {
			Acute   string
			Chronic string
		}{
			Acute:   "Anabolic signaling within 1-6 hours",
			Chronic: "Measurable formation in 6-8 weeks",
		},
	}
)

// GetPrescriptionByCondition returns appropriate exercise prescriptions for a bone condition
func GetPrescriptionByCondition(condition string) []ExercisePrescription {
	var prescriptions []ExercisePrescription

	switch condition {
	case "Osteopenia":
		prescriptions = append(prescriptions, BoneMineralDensity, OsteocyteActivation)
	case "Osteoporosis":
		prescriptions = append(prescriptions, BoneMineralDensity)
	case "Disuse Osteoporosis":
		prescriptions = append(prescriptions, OsteocyteActivation, BoneRemodeling)
	case "Athletic Performance":
		prescriptions = append(prescriptions, BoneAnabolism)
	case "Metabolic Bone Disease":
		prescriptions = append(prescriptions, BoneRemodeling)
	default:
		prescriptions = append(prescriptions, BoneMineralDensity)
	}

	return prescriptions
}

// PredictBoneResponse estimates the effects of an exercise prescription
// on bone parameters based on the protocol, duration of adherence, and patient factors
func PredictBoneResponse(prescription ExercisePrescription, weeks int,
	lowBaselineDensity bool, highResorption bool) BoneResponse {

	// Initialize response
	response := BoneResponse{
		OsteokineChanges:    make(map[string]float64),
		CellActivityChanges: make(map[string]float64),
		StructuralMetrics:   make(map[string]float64),
		FormationMetrics:    make(map[string]float64),
		ResorptionMetrics:   make(map[string]float64),
	}

	// Calculate intensity factor
	intensityFactor := float64(prescription.IntensityPercent) / 100.0

	// Load magnitude factor
	loadFactor := 0.0
	switch prescription.LoadMagnitude {
	case "Low":
		loadFactor = 0.6
	case "Moderate":
		loadFactor = 0.8
	case "High":
		loadFactor = 1.0
	case "Variable":
		loadFactor = 0.85
	}

	// Calculate duration factor (diminishing returns after 16 weeks for bone)
	durationFactor := 0.0
	if weeks <= 16 {
		durationFactor = float64(weeks) / 16.0
	} else {
		durationFactor = 1.0 + (0.1 * float64(weeks-16) / 16.0)
	}

	// Base effect factors
	var formationEffect, resorptionEffect, mineralizationEffect, architectureEffect float64

	// Calculate base effects based on prescription type
	switch prescription.Name {
	case "Bone Mineral Density Protocol":
		formationEffect = 0.4 * intensityFactor * loadFactor * durationFactor
		resorptionEffect = -0.2 * intensityFactor * loadFactor * durationFactor
		mineralizationEffect = 0.4 * intensityFactor * loadFactor * durationFactor
		architectureEffect = 0.3 * intensityFactor * loadFactor * durationFactor

	case "Osteocyte Network Stimulation":
		formationEffect = 0.3 * intensityFactor * loadFactor * durationFactor
		resorptionEffect = -0.1 * intensityFactor * loadFactor * durationFactor
		mineralizationEffect = 0.2 * intensityFactor * loadFactor * durationFactor
		architectureEffect = 0.4 * intensityFactor * loadFactor * durationFactor

	case "Bone Remodeling Optimization":
		formationEffect = 0.3 * intensityFactor * loadFactor * durationFactor
		resorptionEffect = -0.3 * intensityFactor * loadFactor * durationFactor
		mineralizationEffect = 0.3 * intensityFactor * loadFactor * durationFactor
		architectureEffect = 0.3 * intensityFactor * loadFactor * durationFactor

	case "Osteogenic Loading Protocol":
		formationEffect = 0.5 * intensityFactor * loadFactor * durationFactor
		resorptionEffect = -0.1 * intensityFactor * loadFactor * durationFactor
		mineralizationEffect = 0.3 * intensityFactor * loadFactor * durationFactor
		architectureEffect = 0.4 * intensityFactor * loadFactor * durationFactor
	}

	// Adjust for baseline conditions
	if lowBaselineDensity {
		formationEffect *= 1.3
		mineralizationEffect *= 1.2
	}

	if highResorption {
		resorptionEffect *= 1.5
	}

	// Calculate specific factor changes
	// Osteokine changes
	response.OsteokineChanges["Osteocalcin"] = 15.0 + (40.0 * formationEffect)
	response.OsteokineChanges["Sclerostin"] = -10.0 - (30.0 * formationEffect) // Negative is good (reduction)
	response.OsteokineChanges["Osteopontin"] = 10.0 + (25.0 * architectureEffect)
	response.OsteokineChanges["RANKL"] = -5.0 - (20.0 * resorptionEffect)
	response.OsteokineChanges["OPG"] = 10.0 + (30.0 * formationEffect)
	response.OsteokineChanges["IGF-1 (local)"] = 15.0 + (35.0 * formationEffect)
	response.OsteokineChanges["TGF-β"] = 5.0 + (20.0 * architectureEffect)
	response.OsteokineChanges["FGF23"] = -5.0 - (15.0 * mineralizationEffect)

	// Cell activity changes
	response.CellActivityChanges["Osteoblast activity"] = 20.0 + (60.0 * formationEffect)
	response.CellActivityChanges["Osteoclast activity"] = -5.0 - (25.0 * resorptionEffect)
	response.CellActivityChanges["Osteocyte viability"] = 10.0 + (30.0 * architectureEffect)
	response.CellActivityChanges["MSC recruitment"] = 15.0 + (45.0 * formationEffect)
	response.CellActivityChanges["Osteocyte mechanosensitivity"] = 10.0 + (40.0 * architectureEffect)

	// Structural metrics
	response.StructuralMetrics["Bone mineral density"] = 0.5 + (3.0 * mineralizationEffect)
	response.StructuralMetrics["Trabecular number"] = 0.5 + (2.5 * architectureEffect)
	response.StructuralMetrics["Trabecular thickness"] = 0.3 + (2.0 * architectureEffect)
	response.StructuralMetrics["Cortical thickness"] = 0.2 + (1.5 * mineralizationEffect)
	response.StructuralMetrics["Bone volume fraction"] = 0.5 + (2.5 * architectureEffect)

	// Formation metrics
	response.FormationMetrics["P1NP"] = 10.0 + (40.0 * formationEffect)
	response.FormationMetrics["Bone-specific ALP"] = 5.0 + (30.0 * formationEffect)
	response.FormationMetrics["Osteoid surface"] = 5.0 + (25.0 * formationEffect)
	response.FormationMetrics["Mineralizing surface"] = 5.0 + (20.0 * mineralizationEffect)

	// Resorption metrics
	response.ResorptionMetrics["CTX"] = -5.0 - (20.0 * resorptionEffect) // Negative is reduction
	response.ResorptionMetrics["NTX"] = -5.0 - (20.0 * resorptionEffect) // Negative is reduction
	response.ResorptionMetrics["TRAP5b"] = -3.0 - (15.0 * resorptionEffect)
	response.ResorptionMetrics["Erosion depth"] = -2.0 - (10.0 * resorptionEffect)

	return response
}

// GetExerciseDuration calculates recommended exercise duration based on bone condition and patient factors
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
		duration = int(float64(duration) * 0.75) // 25% shorter for severe conditions
		// Minimum acceptable duration
		if duration < 15 {
			duration = 15
		}
	}

	return duration
}

// CalculateTimeToEffect estimates weeks needed to achieve a specific bone benefit
func CalculateTimeToEffect(prescription ExercisePrescription, benefitTarget string, benefitMagnitude float64) float64 {
	// Default times based on typical adaptation rates (bone is slower than cardiovascular)
	baselineWeeks := 0.0

	switch benefitTarget {
	case "Bone formation markers":
		baselineWeeks = 6.0
	case "Bone mineral density":
		baselineWeeks = 16.0
	case "Trabecular architecture":
		baselineWeeks = 12.0
	case "Cortical thickness":
		baselineWeeks = 20.0
	default:
		baselineWeeks = 12.0
	}

	// Adjust for prescription intensity and frequency
	intensityFactor := float64(prescription.IntensityPercent) / 70.0 // Normalized to moderate intensity
	frequencyFactor := float64(prescription.FrequencyPerWeek) / 3.0  // Normalized to 3x/week

	// Adjust for magnitude of desired benefit
	magnitudeFactor := benefitMagnitude / 20.0 // Normalized to 20% improvement

	// Calculate adjusted timeframe
	adjustedWeeks := baselineWeeks * magnitudeFactor / (intensityFactor * frequencyFactor)

	// Ensure minimum timeframe for bone (bone adapts more slowly than cardiovascular)
	if adjustedWeeks < 4.0 {
		adjustedWeeks = 4.0
	}

	return adjustedWeeks
}
