package pancreas

// PancreaticExerkine represents a signaling molecule derived from pancreatic tissue
type PancreaticExerkine struct {
	Name               string
	CellOrigin         string // Beta, Alpha, Delta, PP, or Epsilon cells
	TargetOrgans       []string
	ExerciseRegulation string  // How exercise affects its secretion
	MolecularWeight    float64 // In kDa
	HalfLifeMinutes    float64
	Function           string
	ActionMechanism    string
}

// Key pancreatic exerkines and their exercise responses
var (
	Insulin = PancreaticExerkine{
		Name:               "Insulin",
		CellOrigin:         "Beta cells",
		TargetOrgans:       []string{"Liver", "Muscle", "Adipose"},
		ExerciseRegulation: "Biphasic", // Acute decrease during exercise, increased sensitivity after
		MolecularWeight:    5.8,
		HalfLifeMinutes:    4.0,
		Function:           "Glucose uptake, protein synthesis, lipogenesis",
		ActionMechanism:    "Receptor tyrosine kinase activation, GLUT4 translocation",
	}

	Glucagon = PancreaticExerkine{
		Name:               "Glucagon",
		CellOrigin:         "Alpha cells",
		TargetOrgans:       []string{"Liver", "Adipose"},
		ExerciseRegulation: "Up", // Increased during exercise
		MolecularWeight:    3.5,
		HalfLifeMinutes:    6.0,
		Function:           "Glucose production, lipolysis",
		ActionMechanism:    "GPCR activation, cAMP signaling",
	}

	PancreaticPolypeptide = PancreaticExerkine{
		Name:               "Pancreatic Polypeptide",
		CellOrigin:         "PP cells",
		TargetOrgans:       []string{"GI tract", "Brain"},
		ExerciseRegulation: "Up", // Increased during prolonged exercise
		MolecularWeight:    4.2,
		HalfLifeMinutes:    7.0,
		Function:           "Appetite regulation, digestive enzyme secretion",
		ActionMechanism:    "Y4 receptor binding",
	}

	Amylin = PancreaticExerkine{
		Name:               "Amylin",
		CellOrigin:         "Beta cells",
		TargetOrgans:       []string{"Brain", "GI tract"},
		ExerciseRegulation: "Biphasic", // Similar to insulin
		MolecularWeight:    3.9,
		HalfLifeMinutes:    12.0,
		Function:           "Satiety, gastric emptying, glycemic control",
		ActionMechanism:    "Amylin receptor binding, area postrema activation",
	}
)

// GetExerciseResponsiveExerkines returns pancreatic exerkines that respond to exercise
func GetExerciseResponsiveExerkines() []PancreaticExerkine {
	return []PancreaticExerkine{Insulin, Glucagon, PancreaticPolypeptide, Amylin}
}

// CalculateExerkineResponseRatio estimates the ratio change in exerkine levels
// based on exercise intensity and duration
func CalculateExerkineResponseRatio(exerkine PancreaticExerkine,
	intensityPercent float64, durationMinutes float64) float64 {

	// Different exerkines have different response patterns to exercise
	switch exerkine.Name {
	case "Insulin":
		// During acute exercise, insulin decreases (0.5-0.7x baseline)
		if durationMinutes < 60 {
			return 0.7 - (0.2 * intensityPercent / 100)
		}
		// After exercise, insulin sensitivity increases
		return 1.0

	case "Glucagon":
		// Increases with intensity and duration
		return 1.0 + (intensityPercent/100)*(durationMinutes/60)

	case "Pancreatic Polypeptide":
		// Primarily responds to duration
		return 1.0 + (0.3 * durationMinutes / 60)

	case "Amylin":
		// Similar to insulin but less pronounced
		if durationMinutes < 60 {
			return 0.8 - (0.1 * intensityPercent / 100)
		}
		return 1.0

	default:
		return 1.0
	}
}
