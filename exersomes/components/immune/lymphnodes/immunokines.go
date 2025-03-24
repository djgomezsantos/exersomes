package lymphnodes

import "time"

// LymphNodeFactor represents a signaling molecule related to lymph node function
type LymphNodeFactor struct {
    Name               string
    Type               string   // "Cytokine", "Chemokine", "Growth factor", etc.
    SourceCells        []string
    TargetCells        []string
    ExerciseRegulation string   // "Up", "Down", "Biphasic"
    PrimaryEffects     []string
    MechanismOfAction  string   // How it affects lymph node function
    ConcentrationRange struct {
        Baseline struct {
            Min  float64
            Max  float64
            Unit string
        }
        PostExercise struct {
            Min  float64
            Max  float64
            Unit string
        }
    }
}

// Key lymph node factors affected by exercise
var (
    CCL21 = LymphNodeFactor{
        Name:               "CCL21 (C-C motif chemokine ligand 21)",
        Type:               "Chemokine",
        SourceCells:        []string{"High endothelial venules", "Stromal cells", "T-zone fibroblasts"},
        TargetCells:        []string{"T cells", "Dendritic cells"},
        ExerciseRegulation: "Up", // Transiently increased with acute exercise
        PrimaryEffects:     []string{"T cell homing", "Dendritic cell migration", "Lymphocyte trafficking"},
        MechanismOfAction:  "Guides lymphocytes into lymph nodes via CCR7 receptor binding",
        ConcentrationRange: struct {
            Baseline struct { Min float64; Max float64; Unit string }
            PostExercise struct { Min float64; Max float64; Unit string }
        }{
            Baseline: struct { Min float64; Max float64; Unit string }{
                Min:  100.0,
                Max:  300.0,
                Unit: "pg/mL",
            },
            PostExercise: struct { Min float64; Max float64; Unit string }{
                Min:  150.0,
                Max:  450.0,
                Unit: "pg/mL",
            },
        },
    }

    CXCL13 = LymphNodeFactor{
        Name:               "CXCL13 (C-X-C motif chemokine ligand 13)",
        Type:               "Chemokine",
        SourceCells:        []string{"Follicular dendritic cells", "B-zone fibroblasts"},
        TargetCells:        []string{"B cells", "T follicular helper cells"},
        ExerciseRegulation: "Up", // Modest increase with intense exercise
        PrimaryEffects:     []string{"B cell homing", "Germinal center organization", "Antibody response facilitation"},
        MechanismOfAction:  "Organizes B cells in follicles via CXCR5 receptor binding",
        ConcentrationRange: struct {
            Baseline struct { Min float64; Max float64; Unit string }
            PostExercise struct { Min float64; Max float64; Unit string }
        }{
            Baseline: struct { Min float64; Max float64; Unit string }{
                Min:  30.0,
                Max:  100.0,
                Unit: "pg/mL",
            },
            PostExercise: struct { Min float64; Max float64; Unit string }{
                Min:  40.0,
                Max:  140.0,
                Unit: "pg/mL",
            },
        },
    }

    IL7 = LymphNodeFactor{
        Name:               "Interleukin-7",
        Type:               "Cytokine",
        SourceCells:        []string{"Stromal cells", "Fibroblastic reticular cells"},
        TargetCells:        []string{"T cells", "B cells", "Lymphoid progenitors"},
        ExerciseRegulation: "Up", // Increased with regular exercise
        PrimaryEffects:     []string{"T cell homeostasis", "Naive T cell survival", "Lymphocyte development"},
        MechanismOfAction:  "Maintains T cell populations and supports lymphocyte differentiation",
        ConcentrationRange: struct {
            Baseline struct { Min float64; Max float64; Unit string }
            PostExercise struct { Min float64; Max float64; Unit string }
        }{
            Baseline: struct { Min float64; Max float64; Unit string }{
                Min:  2.0,
                Max:  8.0,
                Unit: "pg/mL",
            },
            PostExercise: struct { Min float64; Max float64; Unit string }{
                Min:  3.0,
                Max:  12.0,
                Unit: "pg/mL",
            },
        },
    }
)

// CalculateLymphNodeResponse estimates changes in lymph node factors with exercise
func CalculateLymphNodeResponse(factor LymphNodeFactor, 
    exerciseIntensityPercent int, durationMinutes int, chronicTrainingWeeks int) float64 {
    
    // Base response factor (1.0 = no change)
    responseFactor := 1.0
    
    // Exercise parameter factors
    intensityFactor := float64(exerciseIntensityPercent) / 100.0
    durationFactor := float64(durationMinutes) / 60.0 // Normalized to 1 hour
    
    // Training adaptation factor
    adaptationFactor := 1.0
    if chronicTrainingWeeks > 0 {
        adaptationWeeks := float64(chronicTrainingWeeks) / 12.0 // Normalized to 12 weeks
        if adaptationWeeks > 1.0 {
            adaptationWeeks = 1.0 + (0.2 * (adaptationWeeks - 1.0)) // Diminishing returns