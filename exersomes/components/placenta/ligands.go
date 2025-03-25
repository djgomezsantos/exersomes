package placenta

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

var (
    // ...existing ligands...
    
    Glucose = Ligand{
        Name:            "Glucose",
        Class:           "Carbohydrate",
        MolecularWeight: 0.18, // kDa
        SourceTissues:   []string{"Liver", "Intestine", "Kidney"},
        TargetTissues:   []string{"All tissues", "Brain", "Skeletal muscle", "Liver"},
        ReceptorFamilies: []string{"GLUTs", "SGLT"},
        Pathways:         []string{"Insulin signaling", "AMPK", "mTOR"},
        ExerciseResponse: struct {
            AcuteRegulation  string
            ChronicEffect    string
            IntensityFactor  float64
            DurationFactor   float64
            RecoveryTime     float64
        }{
            AcuteRegulation:  "Down",
            ChronicEffect:    "Increased clearance",
            IntensityFactor:  1.4,
            DurationFactor:   1.6,
            RecoveryTime:     2.0, // Hours
        },
        PhysiologicalEffect: "Primary energy substrate, metabolic regulation",
        ClinicalRelevance:   []string{"Diabetes", "Insulin resistance", "Metabolic syndrome", "Exercise performance"},
    },
    
    VEGFA = Ligand{
        Name:            "Vascular Endothelial Growth Factor A",
        Class:           "Growth Factor",
        MolecularWeight: 45.0,
        SourceTissues:   []string{"Skeletal muscle", "Endothelial cells", "Adipose tissue", "Platelets"},
        TargetTissues:   []string{"Endothelial cells", "Vascular smooth muscle", "Skeletal muscle"},
        ReceptorFamilies: []string{"VEGFR-1", "VEGFR-2", "Neuropilins"},
        Pathways:         []string{"MAPK/ERK", "PI3K/Akt", "PLC-γ", "Src"},
        ExerciseResponse: struct {
            AcuteRegulation  string
            ChronicEffect    string
            IntensityFactor  float64
            DurationFactor   float64
            RecoveryTime     float64
        }{
            AcuteRegulation:  "Up",
            ChronicEffect:    "Increased baseline levels",
            IntensityFactor:  1.5,
            DurationFactor:   1.3,
            RecoveryTime:     48.0,
        },
        PhysiologicalEffect: "Angiogenesis, vascular permeability, endothelial cell survival",
        ClinicalRelevance:   []string{"Cardiovascular health", "Peripheral artery disease", "Wound healing", "Cancer"},
    },
    
    Apelin = Ligand{
        Name:            "Apelin",
        Class:           "Peptide",
        MolecularWeight: 3.5, // Varies by isoform
        SourceTissues:   []string{"Heart", "Adipose tissue", "Skeletal muscle", "Placenta", "Brain"},
        TargetTissues:   []string{"Heart", "Vessels", "Adipose tissue", "Brain", "Skeletal muscle"},
        ReceptorFamilies: []string{"APJ"},
        Pathways:         []string{"PI3K-Akt", "AMPK", "ERK1/2", "eNOS"},
        ExerciseResponse: struct {
            AcuteRegulation  string
            ChronicEffect    string
            IntensityFactor  float64
            DurationFactor   float64
            RecoveryTime     float64
        }{
            AcuteRegulation:  "Up",
            ChronicEffect:    "Increased sensitivity",
            IntensityFactor:  1.3,
            DurationFactor:   1.2,
            RecoveryTime:     8.0,
        },
        PhysiologicalEffect: "Vasodilation, cardiac contractility, glucose metabolism, fluid homeostasis",
        ClinicalRelevance:   []string{"Hypertension", "Heart failure", "Obesity", "Diabetes"},
    },
)