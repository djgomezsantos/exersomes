package pancreas

// PancreaticReceptor represents a receptor expressed in pancreatic tissue
type PancreaticReceptor struct {
	Name               string
	Type               string   // Receptor family
	CellType           string   // Cell type expressing the receptor
	Ligands            []string // Molecules that bind this receptor
	SignalingPathway   []string
	ExpressionLevel    string // "High", "Medium", "Low"
	ExerciseRegulation string // How exercise affects receptor expression
	Function           string // Physiological function
}

// Key pancreatic receptors affected by exercise
var (
	InsulinReceptor = PancreaticReceptor{
		Name:               "Insulin Receptor",
		Type:               "Receptor tyrosine kinase",
		CellType:           "Beta, Alpha, Delta",
		Ligands:            []string{"Insulin"},
		SignalingPathway:   []string{"PI3K/Akt", "MAPK/ERK"},
		ExpressionLevel:    "Medium",
		ExerciseRegulation: "Up", // Exercise increases sensitivity
		Function:           "Glucose sensing, beta cell proliferation",
	}

	GLP1R = PancreaticReceptor{
		Name:               "GLP-1 Receptor",
		Type:               "G-protein coupled receptor",
		CellType:           "Beta",
		Ligands:            []string{"GLP-1"},
		SignalingPathway:   []string{"cAMP/PKA", "Ca2+ signaling"},
		ExpressionLevel:    "High",
		ExerciseRegulation: "Up", // Exercise increases expression
		Function:           "Potentiation of glucose-stimulated insulin secretion",
	}

	AdrenergicReceptorBeta2 = PancreaticReceptor{
		Name:               "β2-Adrenergic Receptor",
		Type:               "G-protein coupled receptor",
		CellType:           "Alpha, Beta",
		Ligands:            []string{"Epinephrine", "Norepinephrine"},
		SignalingPathway:   []string{"cAMP/PKA"},
		ExpressionLevel:    "Medium",
		ExerciseRegulation: "Down", // Exercise can decrease sensitivity (desensitization)
		Function:           "Regulation of insulin and glucagon secretion during exercise",
	}

	AMPKR = PancreaticReceptor{
		Name:               "AMPK",
		Type:               "Metabolic sensor",
		CellType:           "Beta, Alpha",
		Ligands:            []string{"AMP", "ADP"},
		SignalingPathway:   []string{"LKB1", "CaMKK2"},
		ExpressionLevel:    "Medium",
		ExerciseRegulation: "Up", // Exercise activates AMPK
		Function:           "Energy sensing, beta cell function regulation",
	}
)

// GetExerciseResponsiveReceptors returns pancreatic receptors that respond to exercise
func GetExerciseResponsiveReceptors() []PancreaticReceptor {
	return []PancreaticReceptor{
		InsulinReceptor,
		GLP1R,
		AdrenergicReceptorBeta2,
		AMPKR,
	}
}

// GetReceptorsByExerciseType returns receptors primarily affected by specific exercise type
func GetReceptorsByExerciseType(exerciseType string) []PancreaticReceptor {
	var receptors []PancreaticReceptor

	switch exerciseType {
	case "Aerobic":
		receptors = append(receptors, GLP1R, AdrenergicReceptorBeta2, AMPKR)
	case "HIIT":
		receptors = append(receptors, AdrenergicReceptorBeta2, AMPKR)
	case "Resistance":
		receptors = append(receptors, InsulinReceptor, AMPKR)
	default:
		receptors = GetExerciseResponsiveReceptors()
	}

	return receptors
}
