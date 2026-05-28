// components/muscle/muscle.go
package muscle

import (
	"exersomes/exersomes/molecular_types"
)

// Placeholder function
func MuscleFunction() string {
	return "muscle working"
}

// Myocyte represents a single virtual muscle cell (The Packager)
type Myocyte struct {
	IntracellularCargo []molecular_types.MolecularType
}

// PackageExersome simulates a cell wrapping up-regulated molecules into a vesicle
func (c *Myocyte) PackageExersome(exerciseIntensity float64) molecular_types.Vesicles {
	// In a real simulation, intensity/prescription dictates which cargo gets loaded
	return molecular_types.Vesicles{
		ID:    "EV-MUSCLE-01",
		Name:  "Myocyte-Derived Exersome",
		Cargo: c.IntracellularCargo,
	}
}

// Muscle represents a muscle organ with exerkine secretions
type Muscle struct {
	Name      string
	Type      string   // e.g. "skeletal", "cardiac", "smooth"
	FiberType []string // e.g. ["Type I", "Type IIa", "Type IIx"]
	CellCount int64    // Scaling factor for the virtual organ
	Exerkines []molecular_types.Exerkine
}

// SimulateOrganSecretion scales a single cell's exersome output to the whole organ level
func (m *Muscle) SimulateOrganSecretion(myocyte Myocyte, exerciseIntensity float64) (molecular_types.Vesicles, int64) {
	// 1. The representative cell packages the exersome
	baseExersome := myocyte.PackageExersome(exerciseIntensity)

	// 2. Scale the output by the number of active cells recruited by this exercise intensity
	recruitmentFactor := exerciseIntensity / 100.0
	totalVesiclesSecreted := int64(float64(m.CellCount) * recruitmentFactor)

	return baseExersome, totalVesiclesSecreted
}

// SkeletalMuscle returns a Muscle representing the unified skeletal muscle system
func SkeletalMuscle(estimatedTotalCells int64) Muscle {
	return Muscle{
		Name:      "Skeletal Muscle System",
		Type:      "skeletal",
		FiberType: []string{"Type I", "Type IIa", "Type IIx"},
		CellCount: estimatedTotalCells,
		Exerkines: GetMuscleExerkines(),
	}
}
