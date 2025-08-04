// components/encoder/encoder.go
package encoder

import (
    "github.com/gomezdj/exersomes/molecular_types"
    "github.com/gomezdj/exersomes/components/muscle"
    // Add imports for ligands, receptors, prescription as needed
)

// Encoder defines the contract for encoding exersome cargos
type Encoder interface {
    AddMyokines([]molecular_types.Exerkine)
    AddLigands([]muscle.Ligand)
    AddReceptors([]muscle.Receptor)
    AddPrescription(muscle.Prescription)
    Encode() ([]byte, error)
}

// ExersomeEncoder is a concrete implementation
type ExersomeEncoder struct {
    Myokines     []molecular_types.Exerkine
    Ligands      []muscle.Ligand
    Receptors    []muscle.Receptor
    Prescription *muscle.Prescription
}

// AddMyokines appends myokines to the encoder
func (e *ExersomeEncoder) AddMyokines(myokines []molecular_types.Exerkine) {
    e.Myokines = append(e.Myokines, myokines...)
}

// AddLigands appends ligands to the encoder
func (e *ExersomeEncoder) AddLigands(ligands []muscle.Ligand) {
    e.Ligands = append(e.Ligands, ligands...)
}

// AddReceptors appends receptors to the encoder
func (e *ExersomeEncoder) AddReceptors(receptors []muscle.Receptor) {
    e.Receptors = append(e.Receptors, receptors...)
}

// AddPrescription sets the prescription protocol
func (e *ExersomeEncoder) AddPrescription(prescription muscle.Prescription) {
    e.Prescription = &prescription
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
    encoder.AddMyokines(muscle.GetMuscleExerkines()) // from myokines.go
    // encoder.AddLigands(...) // from ligands.go
    // encoder.AddReceptors(...) // from receptors.go
    // encoder.AddPrescription(...) // from prescription.go
    data, err := encoder.Encode()
    if err != nil {
        // handle error
    }
    // use data
}
