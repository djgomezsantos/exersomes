// exersomes/molecular_types/common.go
package molecular_types

type Exerkine struct {
    Name            string
    Category        string  // Protein, Metabolite, miRNA, etc.
    TissueSources   []string
    BiologicalFunc  string
    Sequence        string  // Chemical formula or AA sequence
    PDBID          string  // For proteins
    Receptors       []string
}

type ExerkineNetwork struct {
    Nodes         []Exerkine
    Interactions  []Interaction
}

type Interaction struct {
    Ligand      string
    Receptor    string
    Tissue      string
    Pathway     string
    Strength    float64 // Optional confidence score
}

// ExerciseMolecule represents a common interface for all exercise-responsive molecules
type ExerciseMolecule interface {
	GetName() string
	GetSource() []string
	GetExerciseResponse(exerciseType string, intensity float64, duration int) float64
	GetTargetTissues() []string
	GetBiologicalEffects() []string
}
