package placenta

// Baptokine represents a signaling molecule secreted from brown/beige adipose tissue
type Baptokine struct {
    Name              string
    GeneID            string
    MolecularWeight   float64  // In kDa
    ExerciseResponse  string   // "Up", "Down", "Biphasic", "Unknown"
    TemporalPattern   string   // "Acute", "Chronic", "Both", "Unknown"
    SecretionTriggers []string // Factors triggering secretion
    TargetTissues     []string
    PrimaryEffects    []string
    SignalingPathways []string
    ClinicalRelevance []string
    BaselineRange     struct {
        Min  float64
        Max  float64
        Unit string
    }
}

// Collection of key exercise-responsive baptokines
var (
    NRG4 = Baptokine{
        Name:             "Neuregulin 4",
        GeneID:           "NRG4",
        MolecularWeight:  26.0,
        ExerciseResponse: "Up",
        TemporalPattern:  "Chronic",
        SecretionTriggers: []string{
            "Cold exposure",
            "β-adrenergic stimulation",
            "Exercise training",
            "PGC-1α activation",
        },
        TargetTissues: []string{"Liver", "Pancreas", "Skeletal muscle"},
        PrimaryEffects: []string{
            "Reduced hepatic lipogenesis",
            "Improved energy metabolism",
            "Attenuated liver injury",
            "Insulin secretion modulation",
        },
        SignalingPathways: []string{"ErbB3/ErbB4", "PI3K-Akt", "MAPK"},
        ClinicalRelevance: []string{
            "Non-alcoholic fatty liver disease",
            "Type 2 diabetes",
            "Obesity",
            "Insulin resistance",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  0.5,
            Max:  3.0,
            Unit: "ng/mL",
        },
    },

    DiHOME = Baptokine{
        Name:             "12,13-diHOME",
        GeneID:           "EPHX1/2", // Produced by epoxide hydrolases
        MolecularWeight:  0.313,     // Small lipid
        ExerciseResponse: "Up",
        TemporalPattern:  "Acute",
        SecretionTriggers: []string{
            "Exercise",
            "Cold exposure",
            "β-adrenergic stimulation",
        },
        TargetTissues: []string{"Skeletal muscle", "Heart", "Liver"},
        PrimaryEffects: []string{
            "Fatty acid uptake in muscle",
            "Improved exercise capacity",
            "Enhanced substrate utilization",
            "Mitochondrial function",
        },
        SignalingPathways: []string{"PPAR-α", "CD36/FATP1"},
        ClinicalRelevance: []string{
            "Exercise performance",
            "Metabolic health",
            "Substrate utilization disorders",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  20.0,
            Max:  100.0,
            Unit: "nM",
        },
    },

    Metrnl = Baptokine{
        Name:             "Meteorin-like/Subfatin",
        GeneID:           "METRNL",
        MolecularWeight:  29.0,
        ExerciseResponse: "Up",
        TemporalPattern:  "Both",
        SecretionTriggers: []string{
            "PGC-1α4 activation",
            "Resistance exercise",
            "Cold exposure",
            "Adipocyte browning",
        },
        TargetTissues: []string{"Adipose tissue", "Immune cells", "Liver"},
        PrimaryEffects: []string{
            "Adipose tissue browning",
            "Alternative macrophage activation",
            "Energy expenditure",
            "Heat production",
        },
        SignalingPathways: []string{"IL-4/STAT6", "AMPK", "PGC-1α"},
        ClinicalRelevance: []string{
            "Obesity",
            "Insulin resistance",
            "Type 2 diabetes",
            "Inflammation",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  100.0,
            Max:  500.0,
            Unit: "pg/mL",
        },
    },

    CXCL14 = Baptokine{
        Name:             "CXCL14/BRAK",
        GeneID:           "CXCL14",
        MolecularWeight:  9.4,
        ExerciseResponse: "Up",
        TemporalPattern:  "Chronic",
        SecretionTriggers: []string{
            "Cold exposure",
            "Exercise training",
            "Browning of adipose tissue",
        },
        TargetTissues: []string{"Adipose tissue", "Immune cells", "Muscle"},
        PrimaryEffects: []string{
            "Macrophage recruitment",
            "Adipose browning",
            "Energy expenditure",
            "Glucose uptake",
        },
        SignalingPathways: []string{"CXCR4", "MAPK", "PI3K-Akt"},
        ClinicalRelevance: []string{
            "Obesity",
            "Insulin resistance",
            "Type 2 diabetes",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  10.0,
            Max:  100.0,
            Unit: "pg/mL",
        },
    },

    BMP8b = Baptokine{
        Name:             "Bone Morphogenetic Protein 8b",
        GeneID:           "BMP8B",
        MolecularWeight:  34.0,
        ExerciseResponse: "Up",
        TemporalPattern:  "Chronic",
        SecretionTriggers: []string{
            "Cold exposure",
            "Exercise",
            "Norepinephrine",
            "Sympathetic stimulation",
        },
        TargetTissues: []string{"Brown adipose tissue", "White adipose tissue", "Hypothalamus"},
        PrimaryEffects: []string{
            "Thermogenic activation",
            "Adipocyte browning",
            "Sympathetic nerve sensitivity",
            "Energy expenditure",
        },
        SignalingPathways: []string{"p38 MAPK", "SMAD", "AMPK"},
        ClinicalRelevance: []string{
            "Obesity",
            "Metabolic disorders",
            "Energy expenditure disorders",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  1.0,
            Max:  10.0,
            Unit: "ng/mL",
        },
    },

    FGF21bat = Baptokine{
        Name:             "Fibroblast Growth Factor 21 (BAT-derived)",
        GeneID:           "FGF21",
        MolecularWeight:  22.0,
        ExerciseResponse: "Up",
        TemporalPattern:  "Both",
        SecretionTriggers: []string{
            "Cold exposure",
            "Exercise",
            "Norepinephrine",
            "PPARγ activation",
            "ATF2 signaling",
        },
        TargetTissues: []string{"Adipose tissue", "Brain", "Liver", "Pancreas", "Skeletal muscle"},
		PrimaryEffects: []string{
            "Glucose homeostasis",
            "Fatty acid oxidation",
            "Ketogenesis",
            "Energy expenditure",
            "Thermogenic activation",
            "Insulin sensitization",
        },
        SignalingPathways: []string{"FGFR1/β-Klotho", "ERK1/2", "AMPK", "SIRT1", "PGC-1α"},
        ClinicalRelevance: []string{
            "Type 2 diabetes",
            "Obesity",
            "Non-alcoholic fatty liver disease",
            "Cardiovascular disease",
            "Exercise adaptation",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  50.0,
            Max:  300.0,
            Unit: "pg/mL",
        },
    },
	IL6bat = Baptokine{
        Name:             "Interleukin-6 (BAT-derived)",
        GeneID:           "IL6",
        MolecularWeight:  21.0,
        ExerciseResponse: "Up",
        TemporalPattern:  "Acute",
        SecretionTriggers: []string{
            "Exercise",
            "Cold exposure",
            "β3-adrenergic stimulation",
            "NF-κB signaling",
        },
        TargetTissues: []string{"Liver", "Skeletal muscle", "Adipose tissue", "Immune cells"},
        PrimaryEffects: []string{
            "Glucose metabolism",
            "Insulin sensitivity",
            "Fatty acid oxidation",
            "Anti-inflammatory effects",
            "Hepatic glucose production",
        },
        SignalingPathways: []string{"JAK/STAT3", "MAPK", "PI3K-Akt"},
        ClinicalRelevance: []string{
            "Obesity",
            "Type 2 diabetes",
            "Inflammation",
            "Exercise adaptation",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  1.0,
            Max:  10.0,
            Unit: "pg/mL",
        },
    },

    SLIT2C = Baptokine{
        Name:             "SLIT2-C Fragment",
        GeneID:           "SLIT2",
        MolecularWeight:  50.0,
        ExerciseResponse: "Up",
        TemporalPattern:  "Chronic",
        SecretionTriggers: []string{
            "Cold exposure",
            "Exercise training",
            "Thermogenic activation",
            "PKA signaling",
        },
        TargetTissues: []string{"White adipose tissue", "Beige adipocytes", "Vasculature"},
        PrimaryEffects: []string{
            "Thermogenic programming",
            "Adipocyte beiging",
            "Mitochondrial biogenesis",
            "Energy expenditure",
            "Angiogenesis",
        },
        SignalingPathways: []string{"PKA", "PRDM16", "PGC-1α", "UCP1"},
        ClinicalRelevance: []string{
            "Obesity",
            "Metabolic syndrome",
            "Type 2 diabetes",
            "Thermogenesis disorders",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  0.1,
            Max:  2.0,
            Unit: "ng/mL",
        },
    },

    BAT_Adiponectin = Baptokine{
        Name:             "Adiponectin (BAT-derived)",
        GeneID:           "ADIPOQ",
        MolecularWeight:  30.0,
        ExerciseResponse: "Up",
        TemporalPattern:  "Chronic",
        SecretionTriggers: []string{
            "Cold exposure",
            "Exercise training",
            "PPAR activation",
            "Sympathetic stimulation",
        },
        TargetTissues: []string{"Liver", "Skeletal muscle", "Adipose tissue", "Brain"},
        PrimaryEffects: []string{
            "Insulin sensitivity",
            "Fatty acid oxidation",
            "Anti-inflammatory activity",
            "Thermogenic activation",
            "Glucose uptake enhancement",
        },
        SignalingPathways: []string{"AdipoR1/R2", "AMPK", "PPAR-α", "p38 MAPK", "PGC-1α"},
        ClinicalRelevance: []string{
            "Insulin resistance",
            "Type 2 diabetes",
            "Obesity",
            "Cardiovascular disease",
            "Non-alcoholic fatty liver disease",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  5.0,
            Max:  30.0,
            Unit: "μg/mL",
        },
    },
)

// GetExerciseResponsiveBaptokines returns a slice of baptokines that respond to exercise
func GetExerciseResponsiveBaptokines() []Baptokine {
    return []Baptokine{
        NRG4,
        DiHOME,
        Metrnl,
        CXCL14,
        BMP8b,
        FGF21bat,
        IL6bat,
        SLIT2C,
        BAT_Adiponectin,
    }
}

// GetBaptokineByName returns a baptokine by its name
func GetBaptokineByName(name string) (Baptokine, bool) {
    for _, baptokine := range GetExerciseResponsiveBaptokines() {
        if baptokine.Name == name {
            return baptokine, true
        }
    }
    return Baptokine{}, false
}

// FilterBaptokinesByExerciseResponse returns baptokines with a specific exercise response
func FilterBaptokinesByExerciseResponse(response string) []Baptokine {
    var filtered []Baptokine
    for _, baptokine := range GetExerciseResponsiveBaptokines() {
        if baptokine.ExerciseResponse == response {
            filtered = append(filtered, baptokine)
        }
    }
    return filtered
}

// CalculateAcuteExerciseResponse predicts the relative change in baptokine levels 
// after an acute exercise bout based on intensity and duration
func CalculateAcuteExerciseResponse(baptokine Baptokine, 
    intensityPercent int, durationMinutes int) float64 {
    
    // Default no change
    responseMultiplier := 1.0
    
    // Only process if baptokine responds acutely to exercise
    if baptokine.TemporalPattern != "Acute" && baptokine.TemporalPattern != "Both" {
        return responseMultiplier
    }
    
    // Calculate intensity factor (0.0-2.0 scale)
    intensityFactor := float64(intensityPercent) / 50.0
    
    // Calculate duration factor (normalized to 60 minutes)
    durationFactor := float64(durationMinutes) / 60.0
    if durationFactor > 2.0 {
        durationFactor = 2.0 // Cap very long exercise
    }
    
    // Calculate response based on specific baptokines
    switch baptokine.Name {
    case "12,13-diHOME":
        // Strongly responsive to acute exercise, especially higher intensities
        responseMultiplier = 1.0 + (intensityFactor * 1.5 * durationFactor)
        
    case "Fibroblast Growth Factor 21 (BAT-derived)":
        // Moderate acute response that increases with duration
        responseMultiplier = 1.0 + (intensityFactor * 0.8 * durationFactor)
        
    case "Interleukin-6 (BAT-derived)":
        // Strong acute response to high intensity
        responseMultiplier = 1.0 + (intensityFactor * intensityFactor * 0.7 * durationFactor)
        
    case "Meteorin-like/Subfatin":
        // Modest acute response
        responseMultiplier = 1.0 + (intensityFactor * 0.5 * durationFactor)
        
    default:
        // Generic modest response for other baptokines
        responseMultiplier = 1.0 + (intensityFactor * 0.3 * durationFactor)
    }
    
    return responseMultiplier
}