package liver

// LiverReceptor represents a receptor expressed in liver tissue
type LiverReceptor struct {
	Name              string
	Type              string   // Receptor type/family
	CellTypes         []string // Types of liver cells expressing this receptor
	Ligands           []string // Molecules that bind to this receptor
	SignalingPathways []string
	ExpressionLevel   string // "High", "Medium", "Low" in normal liver
	ExerciseEffect    struct {
		Regulation   string // "Up", "Down", "Biphasic", or "No change"
		Sensitivity  string // "Increased", "Decreased", "No change"
		TimeToChange string // "Acute", "Chronic"
	}
	MetabolicEffects   []string // Effects when activated
	DiseaseAssociation []string // Associated liver conditions
}

// Key liver receptors affected by exercise
var (
	InsulinReceptor = LiverReceptor{
		Name:              "Insulin Receptor",
		Type:              "Receptor tyrosine kinase",
		CellTypes:         []string{"Hepatocyte", "Stellate cells", "Kupffer cells"},
		Ligands:           []string{"Insulin"},
		SignalingPathways: []string{"PI3K/Akt", "MAPK/ERK", "mTOR"},
		ExpressionLevel:   "High",
		ExerciseEffect: struct {
			Regulation   string
			Sensitivity  string
			TimeToChange string
		}{
			Regulation:   "No change",
			Sensitivity:  "Increased",
			TimeToChange: "Chronic",
		},
		MetabolicEffects:   []string{"Glycogen synthesis", "Lipogenesis", "Protein synthesis", "Suppression of gluconeogenesis"},
		DiseaseAssociation: []string{"Insulin resistance", "NAFLD", "Type 2 diabetes"},
	}

	GlucagonReceptor = LiverReceptor{
		Name:              "Glucagon Receptor",
		Type:              "G-protein coupled receptor",
		CellTypes:         []string{"Hepatocyte"},
		Ligands:           []string{"Glucagon"},
		SignalingPathways: []string{"cAMP/PKA", "Ca2+ signaling"},
		ExpressionLevel:   "High",
		ExerciseEffect: struct {
			Regulation   string
			Sensitivity  string
			TimeToChange string
		}{
			Regulation:   "No change",
			Sensitivity:  "Increased",
			TimeToChange: "Acute",
		},
		MetabolicEffects:   []string{"Glycogenolysis", "Gluconeogenesis", "Ketogenesis"},
		DiseaseAssociation: []string{"Diabetes", "Hypoglycemia"},
	}

	AMPKR = LiverReceptor{
		Name:              "AMPK",
		Type:              "Metabolic sensor",
		CellTypes:         []string{"Hepatocyte", "Kupffer cells"},
		Ligands:           []string{"AMP", "ADP", "Metformin", "AICAR"},
		SignalingPathways: []string{"LKB1", "CaMKK2", "TAK1"},
		ExpressionLevel:   "Medium",
		ExerciseEffect: struct {
			Regulation   string
			Sensitivity  string
			TimeToChange string
		}{
			Regulation:   "Up",
			Sensitivity:  "Increased",
			TimeToChange: "Both",
		},
		MetabolicEffects:   []string{"Fatty acid oxidation", "Glucose uptake", "Mitochondrial biogenesis", "Inhibition of lipogenesis"},
		DiseaseAssociation: []string{"NAFLD", "Metabolic syndrome", "Hepatic steatosis"},
	}

	PPARalpha = LiverReceptor{
		Name:              "PPAR-α",
		Type:              "Nuclear receptor",
		CellTypes:         []string{"Hepatocyte"},
		Ligands:           []string{"Fatty acids", "Fibrates", "Eicosanoids"},
		SignalingPathways: []string{"RXR heterodimer", "PGC-1α"},
		ExpressionLevel:   "High",
		ExerciseEffect: struct {
			Regulation   string
			Sensitivity  string
			TimeToChange string
		}{
			Regulation:   "Up",
			Sensitivity:  "Increased",
			TimeToChange: "Chronic",
		},
		MetabolicEffects:   []string{"Fatty acid oxidation", "Ketogenesis", "Lipoprotein metabolism", "Anti-inflammatory"},
		DiseaseAssociation: []string{"NAFLD", "Dyslipidemia", "Atherosclerosis"},
	}

	GLP1R = LiverReceptor{
		Name:              "GLP-1 Receptor",
		Type:              "G-protein coupled receptor",
		CellTypes:         []string{"Hepatocyte", "Cholangiocytes"},
		Ligands:           []string{"GLP-1", "Exendin-4"},
		SignalingPathways: []string{"cAMP/PKA", "Epac"},
		ExpressionLevel:   "Low",
		ExerciseEffect: struct {
			Regulation   string
			Sensitivity  string
			TimeToChange string
		}{
			Regulation:   "Up",
			Sensitivity:  "Increased",
			TimeToChange: "Chronic",
		},
		MetabolicEffects:   []string{"Reduced lipogenesis", "Improved insulin sensitivity", "Reduced inflammation", "Antifibrotic"},
		DiseaseAssociation: []string{"NAFLD", "NASH", "Liver fibrosis", "Type 2 diabetes"},
	}

	FGF21Receptor = LiverReceptor{
		Name:              "FGFR1c/β-Klotho Complex",
		Type:              "Receptor tyrosine kinase/co-receptor",
		CellTypes:         []string{"Hepatocyte"},
		Ligands:           []string{"FGF21"},
		SignalingPathways: []string{"ERK1/2", "FRS2α", "PI3K/Akt"},
		ExpressionLevel:   "Medium",
		ExerciseEffect: struct {
			Regulation   string
			Sensitivity  string
			TimeToChange string
		}{
			Regulation:   "Up",
			Sensitivity:  "Increased",
			TimeToChange: "Chronic",
		},
		MetabolicEffects:   []string{"Fatty acid oxidation", "Ketogenesis", "Glucose regulation", "Energy expenditure"},
		DiseaseAssociation: []string{"NAFLD", "Obesity", "Type 2 diabetes"},
	}

	CytokineTNF = LiverReceptor{
		Name:              "TNF Receptor",
		Type:              "Death receptor",
		CellTypes:         []string{"Hepatocyte", "Kupffer cells", "Stellate cells"},
		Ligands:           []string{"TNF-α"},
		SignalingPathways: []string{"NF-κB", "JNK", "Caspase cascade"},
		ExpressionLevel:   "Medium",
		ExerciseEffect: struct {
			Regulation   string
			Sensitivity  string
			TimeToChange string
		}{
			Regulation:   "Down",
			Sensitivity:  "Decreased",
			TimeToChange: "Chronic",
		},
		MetabolicEffects:   []string{"Inflammation", "Insulin resistance", "Apoptosis", "Fibrogenesis"},
		DiseaseAssociation: []string{"NASH", "Hepatic inflammation", "Fibrosis", "Liver injury"},
	}
)

// GetExerciseResponsiveReceptors returns liver receptors that respond to exercise
func GetExerciseResponsiveReceptors() []LiverReceptor {
	return []LiverReceptor{
		InsulinReceptor,
		AMPKR,
		PPARalpha,
		GLP1R,
		FGF21Receptor,
		CytokineTNF,
	}
}

// CalculateReceptorResponse estimates receptor sensitivity change
// based on exercise type, intensity, and duration
func CalculateReceptorResponse(receptor LiverReceptor,
	exerciseType string, intensityPercent int, weeks int) ReceptorResponse {

	// Initialize response object
	response := ReceptorResponse{
		Name:              receptor.Name,
		ExpressionChange:  1.0, // No change by default
		SensitivityChange: 1.0, // No change by default
	}

	// Base change factors
	intensityFactor := float64(intensityPercent) / 100.0
	durationFactor := float64(weeks) / 12.0 // Normalized to 12 weeks

	// Apply receptor-specific changes
	switch receptor.Name {
	case "Insulin Receptor":
		// Insulin sensitivity improves with all exercise types
		response.SensitivityChange = 1.0 + (0.2 * durationFactor)
		if exerciseType == "HIIT" {
			response.SensitivityChange += 0.1 * intensityFactor
		}

	case "AMPK":
		// AMPK is strongly activated by acute exercise and expression increases with training
		response.ExpressionChange = 1.0 + (0.15 * durationFactor)
		response.SensitivityChange = 1.0 + (0.3 * intensityFactor)

	case "PPAR-α":
		// PPAR-α increases with endurance training
		if exerciseType == "Aerobic" {
			response.ExpressionChange = 1.0 + (0.2 * durationFactor)
			response.SensitivityChange = 1.0 + (0.15 * durationFactor)
		} else {
			response.ExpressionChange = 1.0 + (0.1 * durationFactor)
		}

	case "GLP-1 Receptor":
		// GLP-1R increases with regular exercise
		response.ExpressionChange = 1.0 + (0.1 * durationFactor)

	case "FGFR1c/β-Klotho Complex":
		// FGF21 receptor complex increases with training
		response.ExpressionChange = 1.0 + (0.15 * durationFactor)
		response.SensitivityChange = 1.0 + (0.1 * durationFactor)

	case "TNF Receptor":
		// TNF receptor decreases with regular exercise
		response.ExpressionChange = 1.0 - (0.2 * durationFactor)
		if response.ExpressionChange < 0.6 {
			response.ExpressionChange = 0.6 // Floor at 40% reduction
		}
	}

	return response
}

// ReceptorResponse represents changes in receptor expression and sensitivity
type ReceptorResponse struct {
	Name              string
	ExpressionChange  float64 // Factor of change (1.0 = no change)
	SensitivityChange float64 // Factor of change (1.0 = no change)
}

// GetReceptorsByLiverCondition returns receptors relevant to specific liver conditions
func GetReceptorsByLiverCondition(condition string) []LiverReceptor {
	var receptors []LiverReceptor

	switch condition {
	case "NAFLD":
		receptors = []LiverReceptor{InsulinReceptor, AMPKR, PPARalpha, GLP1R, FGF21Receptor}
	case "Insulin Resistance":
		receptors = []LiverReceptor{InsulinReceptor, AMPKR, GLP1R}
	case "Inflammation":
		receptors = []LiverReceptor{CytokineTNF, GLP1R, PPARalpha}
	default:
		receptors = GetExerciseResponsiveReceptors()
	}

	return receptors
}
