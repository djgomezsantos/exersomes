package common

// ExerciseMolecule represents a common interface for all exercise-responsive molecules
type ExerciseMolecule interface {
	GetName() string
	GetSource() []string
	GetExerciseResponse(exerciseType string, intensity float64, duration int) float64
	GetTargetTissues() []string
	GetBiologicalEffects() []string
}
