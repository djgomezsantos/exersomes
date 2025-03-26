package bone

// BoneOsteokine represents a bone-derived signaling molecule
type BoneOsteokine struct {
	Name               string
	TargetOrgans       []string
	ResponseToExercise string
	SignalingPathway   string
	PeakTimeMinutes    int
}

// Example bone osteokines
var (
	OSTN = BoneOsteokine{
		Name:               "Osteocrin",
		TargetOrgans:       []string{"Muscle", "Brain"},
		ResponseToExercise: "Increases with mechanical loading",
		SignalingPathway:   "cGMP-PKG",
		PeakTimeMinutes:    60,
	}

	BGLAP = BoneOsteokine{
		Name:               "Osteocalcin",
		TargetOrgans:       []string{"Pancreas", "Adipose", "Brain", "Muscle"},
		ResponseToExercise: "Increases with bone remodeling",
		SignalingPathway:   "GPRC6A",
		PeakTimeMinutes:    120,
	}

	// Add more bone osteokines
)
