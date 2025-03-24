package heart

// CardiacReceptor represents a receptor expressed in cardiac tissue
type CardiacReceptor struct {
	Name              string
	Type              string   // Receptor family
	CellTypes         []string // Cardiac cells expressing this receptor
	Ligands           []string // Binding molecules
	SignalingPathways []string
	ExpressionLevel   string // "High", "Medium", "Low" in normal heart
	Function          string // Primary physiological role
	ExerciseResponse  struct {
		Expression   string // "Up", "Down", "No change"
		Sensitivity  string // "Increased", "Decreased", "No change"
		TimeToEffect string // "Acute", "Chronic", "Both"
	}
}

// Important cardiac receptors affected by exercise
var (
	BetaAdrenergic = CardiacReceptor{
		Name:              "β-Adrenergic Receptor",
		Type:              "G-protein coupled receptor",
		CellTypes:         []string{"Cardiomyocytes", "Nodal cells", "Coronary vessels"},
		Ligands:           []string{"Epinephrine", "Norepinephrine", "Isoproterenol"},
		SignalingPathways: []string{"cAMP/PKA", "CaMKII", "β-arrestin"},
		ExpressionLevel:   "High",
		Function:          "Chronotropy, inotropy, lusitropy, coronary dilation",
		ExerciseResponse: struct {
			Expression   string
			Sensitivity  string
			TimeToEffect string
		}{
			Expression:   "Down",      // Downregulated with chronic exercise
			Sensitivity:  "Increased", // But sensitivity improves
			TimeToEffect: "Chronic",
		},
	}

	AT1Receptor = CardiacReceptor{
		Name:              "Angiotensin II Receptor Type 1",
		Type:              "G-protein coupled receptor",
		CellTypes:         []string{"Cardiomyocytes", "Fibroblasts", "Vascular smooth muscle"},
		Ligands:           []string{"Angiotensin II"},
		SignalingPathways: []string{"Gq/PLC", "JAK/STAT", "MAPK", "ROS pathways"},
		ExpressionLevel:   "Medium",
		Function:          "Vasoconstriction, hypertrophy, fibrosis, ROS production",
		ExerciseResponse: struct {
			Expression   string
			Sensitivity  string
			TimeToEffect string
		}{
			Expression:   "Down",
			Sensitivity:  "Decreased",
			TimeToEffect: "Chronic",
		},
	}

	IGF1R = CardiacReceptor{
		Name:              "IGF-1 Receptor",
		Type:              "Receptor tyrosine kinase",
		CellTypes:         []string{"Cardiomyocytes", "Endothelial cells"},
		Ligands:           []string{"IGF-1", "Insulin (weak)"},
		SignalingPathways: []string{"PI3K/Akt", "MAPK/ERK", "JAK/STAT"},
		ExpressionLevel:   "Medium",
		Function:          "Physiological hypertrophy, anti-apoptosis, contractility",
		ExerciseResponse: struct {
			Expression   string
			Sensitivity  string
			TimeToEffect string
		}{
			Expression:   "Up",
			Sensitivity:  "Increased",
			TimeToEffect: "Chronic",
		},
	}

	NPRA = CardiacReceptor{
		Name:              "Natriuretic Peptide Receptor A",
		Type:              "Guanylyl cyclase receptor",
		CellTypes:         []string{"Cardiomyocytes", "Fibroblasts", "Vascular cells"},
		Ligands:           []string{"ANP", "BNP"},
		SignalingPathways: []string{"cGMP/PKG", "Phosphodiesterases"},
		ExpressionLevel:   "Medium",
		Function:          "Anti-hypertrophy, anti-fibrotic, lusitropy",
		ExerciseResponse: struct {
			Expression   string
			Sensitivity  string
			TimeToEffect string
		}{
			Expression:   "Up",
			Sensitivity:  "Increased",
			TimeToEffect: "Both",
		},
	}

	Adiponectin = CardiacReceptor{
		Name:              "Adiponectin Receptor 1",
		Type:              "Seven-transmembrane receptor",
		CellTypes:         []string{"Cardiomyocytes", "Endothelial cells"},
		Ligands:           []string{"Adiponectin"},
		SignalingPathways: []string{"AMPK", "PPARα", "Ceramidase"},
		ExpressionLevel:   "Medium",
		Function:          "Energy metabolism, anti-apoptosis, anti-inflammatory",
		ExerciseResponse: struct {
			Expression   string
			Sensitivity  string
			TimeToEffect string
		}{
			Expression:   "Up",
			Sensitivity:  "Increased",
			TimeToEffect: "Chronic",
		},
	}

	TLR4 = CardiacReceptor{
		Name:              "Toll-Like Receptor 4",
		Type:              "Pattern recognition receptor",
		CellTypes:         []string{"Cardiomyocytes", "Macrophages", "Fibroblasts"},
		Ligands:           []string{"LPS", "DAMPs", "Saturated fatty acids"},
		SignalingPathways: []string{"MyD88", "TRIF", "NF-κB", "IRF3"},
		ExpressionLevel:   "Low",
		Function:          "Innate immunity, inflammation, cell death signaling",
		ExerciseResponse: struct {
			Expression   string
			Sensitivity  string
			TimeToEffect string
		}{
			Expression:   "Down",
			Sensitivity:  "Decreased",
			TimeToEffect: "Chronic",
		},
	}
)

// GetExerciseResponsiveReceptors returns cardiac receptors that respond to exercise
func GetExerciseResponsiveReceptors() []CardiacReceptor {
	return []CardiacReceptor{
		BetaAdrenergic,
		AT1Receptor,
		IGF1R,
		NPRA,
		Adiponectin,
		TLR4,
	}
}

// CalculateReceptorResponse estimates changes in expression and sensitivity
func CalculateReceptorResponse(receptor CardiacReceptor,
	exerciseType string, intensityPercent int, weeks int) (float64, float64) {

	// Default: no change
	expressionChange := 1.0
	sensitivityChange := 1.0

	// Training duration effect
	durationFactor := float64(weeks) / 12.0 // Normalized to 12 weeks

	// Apply receptor-specific changes
	switch receptor.Name {
	case "β-Adrenergic Receptor":
		// Expression decreases to reduce chronic sympathetic activation
		expressionChange = 1.0 - (0.2 * durationFactor)
		if expressionChange < 0.7 { // Floor at 30% reduction
			expressionChange = 0.7
		}

		// But sensitivity increases for more efficient signaling
		sensitivityChange = 1.0 + (0.15 * durationFactor)

		// HIIT has stronger effects on beta-receptor remodeling
		if exerciseType == "HIIT" {
			expressionChange *= 0.9
			sensitivityChange *= 1.1
		}

	case "Angiotensin II Receptor Type 1":
		// Downregulated with endurance training
		if exerciseType == "Aerobic" || exerciseType == "HIIT" {
			expressionChange = 1.0 - (0.25 * durationFactor)
			sensitivityChange = 1.0 - (0.15 * durationFactor)
		} else {
			expressionChange = 1.0 - (0.1 * durationFactor)
			sensitivityChange = 1.0 - (0.05 * durationFactor)
		}

	case "IGF-1 Receptor":
		// Upregulated with all exercise types, especially resistance
		expressionChange = 1.0 + (0.2 * durationFactor)
		sensitivityChange = 1.0 + (0.15 * durationFactor)

		if exerciseType == "Resistance" {
			expressionChange *= 1.2
			sensitivityChange *= 1.1
		}

	case "Natriuretic Peptide Receptor A":
		// Acute and chronic upregulation
		expressionChange = 1.0 + (0.3 * durationFactor)
		sensitivityChange = 1.0 + (0.2 * durationFactor)

	case "Adiponectin Receptor 1":
		// Progressive increase with training
		expressionChange = 1.0 + (0.25 * durationFactor)
		sensitivityChange = 1.0 + (0.2 * durationFactor)

	case "Toll-Like Receptor 4":
		// Downregulated with regular exercise
		expressionChange = 1.0 - (0.3 * durationFactor)
		if expressionChange < 0.6 { // Floor at 40% reduction
			expressionChange = 0.6
		}
		sensitivityChange = 1.0 - (0.2 * durationFactor)
	}

	// Ensure we don't have negative values
	if expressionChange < 0 {
		expressionChange = 0
	}
	if sensitivityChange < 0 {
		sensitivityChange = 0
	}

	return expressionChange, sensitivityChange
}

// GetReceptorsByCardiacCondition returns receptors relevant to specific heart conditions
func GetReceptorsByCardiacCondition(condition string) []CardiacReceptor {
	var receptors []CardiacReceptor

	switch condition {
	case "Heart Failure":
		receptors = []CardiacReceptor{BetaAdrenergic, AT1Receptor, NPRA, Adiponectin}
	case "Hypertension":
		receptors = []CardiacReceptor{AT1Receptor, NPRA}
	case "Coronary Artery Disease":
		receptors = []CardiacReceptor{Adiponectin, TLR4}
	case "Physiological Hypertrophy":
		receptors = []CardiacReceptor{IGF1R, NPRA}
	default:
		receptors = GetExerciseResponsiveReceptors()
	}

	return receptors
}
