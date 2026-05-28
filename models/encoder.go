// components/encoder/encoder.go
package encoder

import (
	"exersomes/exersomes/components/muscle"
	"exersomes/exersomes/molecular_types"
	// Add imports for ligands, receptors, prescription as needed
)

// Encoder defines the contract for encoding exersome cargos
type Encoder interface {
	AddExerkines([]molecular_types.Exerkine)
	AddLigands([]muscle.Ligand)
	AddReceptors([]muscle.Receptor)
	AddMetabolites([]molecular_types.ExerciseMetabolite)
	AddPrescription(muscle.ExercisePrescription)
	SetOmicsData(rnaVelocities map[string]RNAVelocity, proteinExpr map[string]float64, lipids map[string]float64)
	Encode() ([]byte, error)
}

// ExersomeEncoder is a concrete implementation for multi-omic embedding
type ExersomeEncoder struct {
	Exerkines          []molecular_types.Exerkine
	Ligands            []muscle.Ligand
	Receptors          []muscle.Receptor
	Metabolites        []molecular_types.ExerciseMetabolite
	Prescription       *muscle.ExercisePrescription
	RNATrajectories    map[string]RNAVelocity // MoTrPAC RNA velocity embedding support
	ProteinExpressions map[string]float64     // Protein expression level measurements
	Lipids             map[string]float64     // Lipidomics interaction tracking
}

// AddExerkines appends exerkines to the encoder
func (e *ExersomeEncoder) AddExerkines(exerkines []molecular_types.Exerkine) {
	e.Exerkines = append(e.Exerkines, exerkines...)
}

// AddLigands appends ligands to the encoder
func (e *ExersomeEncoder) AddLigands(ligands []muscle.Ligand) {
	e.Ligands = append(e.Ligands, ligands...)
}

// AddReceptors appends receptors to the encoder
func (e *ExersomeEncoder) AddReceptors(receptors []muscle.Receptor) {
	e.Receptors = append(e.Receptors, receptors...)
}

// AddMetabolites appends metabolites for metabolic/lipid cross-talk modeling
func (e *ExersomeEncoder) AddMetabolites(metabolites []molecular_types.ExerciseMetabolite) {
	e.Metabolites = append(e.Metabolites, metabolites...)
}

// AddPrescription sets the prescription protocol
func (e *ExersomeEncoder) AddPrescription(prescription muscle.ExercisePrescription) {
	e.Prescription = &prescription
}

// SetOmicsData injects advanced multi-omics expression and trajectory data
func (e *ExersomeEncoder) SetOmicsData(rna map[string]RNAVelocity, prot map[string]float64, lipids map[string]float64) {
	e.RNATrajectories = rna
	e.ProteinExpressions = prot
	e.Lipids = lipids
}

// Encode serializes the Exersome cargo (example: to JSON)
func (e *ExersomeEncoder) Encode() ([]byte, error) {
	// Use your preferred encoding method (e.g., json.Marshal)
	// Compose a struct or map containing all fields
	return nil, nil // Implement encoding logic
}

// Example usage:
func Example() {
	encoder := &ExersomeEncoder{}
	encoder.AddExerkines(muscle.GetExerkines()) // from exerkines.go
	// encoder.AddLigands(...) // from ligands.go
	// encoder.AddReceptors(...) // from receptors.go
	// encoder.AddMetabolites(...) // from metabolites.go
	// encoder.AddPrescription(muscle.HypertrophyProtocol)
	data, err := encoder.Encode()
	if err != nil {
		// handle error
	}
	// use data
}
