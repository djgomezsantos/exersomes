package adipose

// AdiposeReceptor represents a receptor expressed in adipose tissue
type AdiposeReceptor struct {
	Name               string
	Type               string   // Receptor type/family
	AdiposeFraction    string   // "White", "Brown", "Beige", or "All"
	CellType           string   // "Adipocyte", "Preadipocyte", "SVF", "Macrophage", etc.
	Ligands            []string // Molecules that bind to this receptor
	SignalingPathways  []string
	ExpressionLevel    string // "High", "Medium", "Low"
	ExerciseRegulation string // How exercise affects receptor expression/sensitivity
	MetabolicEffect    string // Primary metabolic effect when activated
}

// Key adipose receptors affected by exercise
var (
	Beta3AdrenergicReceptor = AdiposeReceptor{
		Name:               "β3-Adrenergic Receptor",
		Type:               "G-protein coupled receptor",
		AdiposeFraction:    "All",
		CellType:           "Adipocyte",
		Ligands:            []string{"Norepinephrine", "Epinephrine"},
		SignalingPathways:  []string{"cAMP/PKA", "p38 MAPK"},
		ExpressionLevel:    "High",
		ExerciseRegulation: "Up",
		MetabolicEffect:    "Lipolysis, thermogenesis",
	}

	InsulinReceptor = AdiposeReceptor{
		Name:               "Insulin Receptor",
		Type:               "Receptor tyrosine kinase",
		AdiposeFraction:    "All",
		CellType:           "Adipocyte",
		Ligands:            []string{"Insulin"},
		SignalingPathways:  []string{"PI3K/Akt", "MAPK/ERK"},
		ExpressionLevel:    "High",
		ExerciseRegulation: "Up", // Increased sensitivity with exercise
		MetabolicEffect:    "Glucose uptake, lipogenesis, anti-lipolysis",
	}

	PPARGamma = AdiposeReceptor{
		Name:               "PPAR-γ",
		Type:               "Nuclear receptor",
		AdiposeFraction:    "White",
		CellType:           "Adipocyte, Preadipocyte",
		Ligands:            []string{"Fatty acids", "Prostaglandins", "TZDs"},
		SignalingPathways:  []string{"RXR heterodimer", "PGC-1α"},
		ExpressionLevel:    "High",
		ExerciseRegulation: "Up", // Chronically upregulated with training
		MetabolicEffect:    "Adipogenesis, insulin sensitivity, lipid storage",
	}

	LeptinReceptor = AdiposeReceptor{
		Name:               "Leptin Receptor",
		Type:               "Cytokine receptor",
		AdiposeFraction:    "All",
		CellType:           "Adipocyte, Macrophage",
		Ligands:            []string{"Leptin"},
		SignalingPathways:  []string{"JAK/STAT", "PI3K", "MAPK"},
		ExpressionLevel:    "Medium",
		ExerciseRegulation: "Up", // Exercise can improve leptin sensitivity
		MetabolicEffect:    "Lipolysis, energy expenditure",
	}

	AdipoR1 = AdiposeReceptor{
		Name:               "Adiponectin Receptor 1",
		Type:               "Seven-transmembrane receptor",
		AdiposeFraction:    "All",
		CellType:           "Adipocyte",
		Ligands:            []string{"Adiponectin"},
		SignalingPathways:  []string{"AMPK", "p38 MAPK", "PPARα"},
		ExpressionLevel:    "Medium",
		ExerciseRegulation: "Up",
		MetabolicEffect:    "Fatty acid oxidation, insulin sensitivity",
	}

	TNFR1 = AdiposeReceptor{
		Name:               "TNF Receptor 1",
		Type:               "Death receptor",
		AdiposeFraction:    "White",
		CellType:           "Adipocyte, Macrophage",
		Ligands:            []string{"TNF-α"},
		SignalingPathways:  []string{"NF-κB", "JNK", "Caspase cascade"},
		ExpressionLevel:    "Medium",
		ExerciseRegulation: "Down",
		MetabolicEffect:    "Insulin resistance, lipolysis, inflammation",
	}
)

// GetReceptorsByAdiposeType returns receptors primarily expressed in specific adipose tissue types
func GetReceptorsByAdiposeType(adiposeType string) []AdiposeReceptor {
	var receptors []AdiposeReceptor

	switch adiposeType {
	case "White":
		receptors = append(receptors, InsulinReceptor, PPARGamma, TNFR1, LeptinReceptor)
	case "Brown":
		receptors = append(receptors, Beta3AdrenergicReceptor, AdipoR1)
	case "Beige":
		receptors = append(receptors, Beta3AdrenergicReceptor, AdipoR1, PPARGamma)
	default:
		// Return all receptors if adipose type not specified
		receptors = []AdiposeReceptor{
			Beta3AdrenergicReceptor,
			InsulinReceptor,
			PPARGamma,
			LeptinReceptor,
			AdipoR1,
			TNFR1,
		}
	}

	return receptors
}

// GetExerciseResponsiveReceptors returns adipose receptors that respond to exercise
func GetExerciseResponsiveReceptors() []AdiposeReceptor {
	return []AdiposeReceptor{
		Beta3AdrenergicReceptor,
		InsulinReceptor,
		PPARGamma,
		AdipoR1,
	}
}

// CalculateReceptorSensitization estimates receptor sensitivity change
// based on exercise type, intensity, and duration
func CalculateReceptorSensitization(receptor AdiposeReceptor,
	exerciseType string, intensityPercent int, weeks int) float64 {

	// Base change factor
	var changeFactor float64 = 1.0

	// Apply receptor-specific changes
	switch receptor.Name {
	case "β3-Adrenergic Receptor":
		if exerciseType == "HIIT" || exerciseType == "Aerobic" {
			// More responsive to cardio
			intensityEffect := float64(intensityPercent) / 100.0
			timeEffect := float64(weeks) / 8.0 // Normalized to 8 weeks
			changeFactor += 0.3 * intensityEffect * timeEffect
		} else {
			// Less responsive to resistance
			timeEffect := float64(weeks) / 8.0
			changeFactor += 0.1 * timeEffect
		}

	case "Insulin Receptor":
		// Responds to all exercise types
		timeEffect := float64(weeks) / 8.0
		changeFactor += 0.25 * timeEffect

	case "PPAR-γ":
		// More responsive to resistance training
		if exerciseType == "Resistance" {
			timeEffect := float64(weeks) / 8.0
			changeFactor += 0.2 * timeEffect
		} else {
			timeEffect := float64(weeks) / 8.0
			changeFactor += 0.1 * timeEffect
		}

	case "Adiponectin Receptor 1":
		// Responds to all exercise types
		timeEffect := float64(weeks) / 8.0
		changeFactor += 0.15 * timeEffect
	}

	// For downregulated receptors, invert the change
	if receptor.ExerciseRegulation == "Down" {
		inverseFactor := 2.0 - changeFactor
		return inverseFactor
	}

	return changeFactor
}
