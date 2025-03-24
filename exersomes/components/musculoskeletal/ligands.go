package musculoskeletal

import "time"

// Ligand represents a signaling molecule that binds to receptors
type Ligand struct {
	Name                string
	Type                string   // e.g., "Cytokine", "Growth Factor", "Hormone"
	MolecularWeight     float64  // In kDa
	SourceTissues       []string // Tissues that produce this ligand
	TargetReceptors     []string // Receptors this ligand binds to
	HalfLifeMinutes     float64  // Half-life in circulation
	ExerciseResponsive  bool
	ExerciseRegulation  string  // "Up", "Down", or "Biphasic"
	ResponseThreshold   float64 // Exercise intensity threshold (% VO2max or 1RM)
	PeakTimeMinutes     float64 // Time to peak after exercise onset
	BaselineConc        Range   // Normal range in circulation
	ExerciseInducedConc Range   // Concentration after exercise
}

// Range represents a concentration range
type Range struct {
	Min  float64
	Max  float64
	Unit string // e.g., "pg/mL", "ng/mL"
}

// CalculateExerciseResponse predicts the concentration after exercise
func (l *Ligand) CalculateExerciseResponse(intensity float64, duration time.Duration) float64 {
	if !l.ExerciseResponsive || intensity < l.ResponseThreshold {
		return l.BaselineConc.Min
	}

	// Simple model: response is proportional to exercise intensity and duration
	// More sophisticated models would consider more variables
	durationFactor := float64(duration.Minutes()) / 60.0
	intensityFactor := intensity / 100.0

	baseline := (l.BaselineConc.Min + l.BaselineConc.Max) / 2
	maxIncrease := l.ExerciseInducedConc.Max - baseline

	response := baseline + (maxIncrease * intensityFactor * durationFactor)

	// Cap at maximum value
	if response > l.ExerciseInducedConc.Max {
		response = l.ExerciseInducedConc.Max
	}

	return response
}

// Common musculoskeletal ligands involved in exercise response
var (
	IL6 = Ligand{
		Name:                "Interleukin-6",
		Type:                "Cytokine",
		MolecularWeight:     21.0,
		SourceTissues:       []string{"Muscle", "Adipose", "Immune Cells"},
		TargetReceptors:     []string{"IL-6 Receptor"},
		HalfLifeMinutes:     15.0,
		ExerciseResponsive:  true,
		ExerciseRegulation:  "Up",
		ResponseThreshold:   50.0, // Moderate intensity needed
		PeakTimeMinutes:     120.0,
		BaselineConc:        Range{Min: 1.0, Max: 10.0, Unit: "pg/mL"},
		ExerciseInducedConc: Range{Min: 10.0, Max: 100.0, Unit: "pg/mL"},
	}

	IGF1 = Ligand{
		Name:                "Insulin-like Growth Factor 1",
		Type:                "Growth Factor",
		MolecularWeight:     7.6,
		SourceTissues:       []string{"Liver", "Muscle"},
		TargetReceptors:     []string{"IGF-1 Receptor"},
		HalfLifeMinutes:     360.0, // 6 hours
		ExerciseResponsive:  true,
		ExerciseRegulation:  "Biphasic", // Acute decrease, chronic increase
		ResponseThreshold:   70.0,       // Higher intensity needed
		PeakTimeMinutes:     2880.0,     // 48 hours for chronic adaptation
		BaselineConc:        Range{Min: 100.0, Max: 300.0, Unit: "ng/mL"},
		ExerciseInducedConc: Range{Min: 150.0, Max: 400.0, Unit: "ng/mL"},
	}

	Myostatin = Ligand{
		Name:                "Myostatin (GDF-8)",
		Type:                "TGF-β Family",
		MolecularWeight:     25.0,
		SourceTissues:       []string{"Muscle"},
		TargetReceptors:     []string{"Activin Receptor Type-2B"},
		HalfLifeMinutes:     480.0, // 8 hours
		ExerciseResponsive:  true,
		ExerciseRegulation:  "Down", // Exercise decreases myostatin
		ResponseThreshold:   75.0,   // Higher intensity needed
		PeakTimeMinutes:     1440.0, // 24 hours
		BaselineConc:        Range{Min: 2.0, Max: 8.0, Unit: "ng/mL"},
		ExerciseInducedConc: Range{Min: 1.0, Max: 5.0, Unit: "ng/mL"},
	}

	IL15 = Ligand{
		Name:                "Interleukin-15",
		Type:                "Cytokine",
		MolecularWeight:     14.0,
		SourceTissues:       []string{"Muscle", "Immune Cells"},
		TargetReceptors:     []string{"IL-15 Receptor"},
		HalfLifeMinutes:     45.0,
		ExerciseResponsive:  true,
		ExerciseRegulation:  "Up",
		ResponseThreshold:   65.0,  // Moderate-high intensity
		PeakTimeMinutes:     180.0, // 3 hours
		BaselineConc:        Range{Min: 1.5, Max: 5.0, Unit: "pg/mL"},
		ExerciseInducedConc: Range{Min: 3.0, Max: 15.0, Unit: "pg/mL"},
	}
)

// GetExerciseResponsiveLigands returns a list of ligands that respond to exercise
func GetExerciseResponsiveLigands() []Ligand {
	return []Ligand{IL6, IGF1, Myostatin, IL15}
}

// PredictLigandResponse calculates expected ligand levels for a given exercise protocol
func PredictLigandResponse(prescription ExercisePrescription) map[string]float64 {
	response := make(map[string]float64)
	intensity := float64(prescription.IntensityPercent1RM)
	duration := time.Duration(prescription.SetsCount*prescription.RepsPerSet*4) * time.Second

	ligands := GetExerciseResponsiveLigands()
	for _, ligand := range ligands {
		response[ligand.Name] = ligand.CalculateExerciseResponse(intensity, duration)
	}

	return response
}
