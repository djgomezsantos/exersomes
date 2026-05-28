// exersomes/molecular_types/common.go
package molecular_types

type Exerkine struct {
	Name           string
	Category       string // Protein, Metabolite, miRNA, etc.
	TissueSources  []string
	BiologicalFunc string
	Sequence       string // Chemical formula or AA sequence
	PDBID          string // For proteins
	Receptors      []string
}

type ExerkineNetwork struct {
	Nodes        []Exerkine
	Interactions []Interaction
}

type Interaction struct {
	Ligand   string
	Receptor string
	Tissue   string
	Pathway  string
	Strength float64 // Optional confidence score
}

// ExerciseMolecule represents a common interface for all exercise-responsive molecules
type ExerciseMolecule interface {
	GetName() string
	GetSource() []string
	GetExerciseResponse(exerciseType string, intensity float64, duration int) float64
	GetTargetTissues() []string
	GetBiologicalEffects() []string
}

// MolecularType defines a generic interface for diverse molecular entities
type MolecularType interface {
	GetID() string
	GetName() string
}

// Protein represents a protein-based exerkine or molecule (e.g., IL-6, Irisin)
type Protein struct {
	ID       string
	Name     string
	Sequence string
	PDBID    string
}

func (p Protein) GetID() string   { return p.ID }
func (p Protein) GetName() string { return p.Name }

// RNA represents a messenger RNA or other RNA-based molecule
type RNA struct {
	ID       string
	Name     string
	Sequence string
}

func (r RNA) GetID() string   { return r.ID }
func (r RNA) GetName() string { return r.Name }

// Metabolite represents a metabolic molecule (e.g., Lactate)
type Metabolite struct {
	ID   string
	Name string
}

func (m Metabolite) GetID() string   { return m.ID }
func (m Metabolite) GetName() string { return m.Name }

// MiRNA represents a microRNA molecule
type MiRNA struct {
	ID   string
	Name string
}

func (m MiRNA) GetID() string   { return m.ID }
func (m MiRNA) GetName() string { return m.Name }

// Vesicles represents an extracellular vesicle (e.g., exosome) and its cargo
type Vesicles struct {
	ID    string
	Name  string
	Cargo []MolecularType
}

func (v Vesicles) GetID() string   { return v.ID }
func (v Vesicles) GetName() string { return v.Name }
