// components/muscle/muscle.go
package muscle

import (
	"exersomes/exersomes/molecular_types"
)

// Placeholder function
func MuscleFunction() string {
	return "muscle working"
}

// Muscle represents a muscle organ with exerkine secretions
type Muscle struct {
	Name      string
	Type      string   // e.g. "skeletal", "cardiac", "smooth"
	FiberType []string // e.g. ["Type I", "Type IIa", "Type IIx"]
	Exerkines []molecular_types.Exerkine
}

// NewSkeletalMuscle returns a Muscle with common skeletal properties
func NewSkeletalMuscle(name string) Muscle {
	return Muscle{
		Name:      name,
		Type:      "skeletal",
		FiberType: []string{"Type I", "Type IIa", "Type IIx"},
		Exerkines: GetMuscleExerkines(),
	}
}

// Example muscles
var (
	Deltoid         = NewSkeletalMuscle("Deltoid")
	BicepsBrachii   = NewSkeletalMuscle("Biceps Brachii")
	TricepsBrachii  = NewSkeletalMuscle("Triceps Brachii")
	AbominalMuscles = NewSkeletalMuscle("Abdominal Muscles")
	VastusLateralis = NewSkeletalMuscle("Vastus Lateralis")
	Gastrocnemius   = NewSkeletalMuscle("Gastrocnemius")
)
