package placenta

// Adipokine represents a signaling molecule secreted from white adipose tissue
type Adipokine struct {
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

// Collection of key exercise-responsive adipokines
var (
    Adiponectin = Adipokine{
        Name:             "Adiponectin",
        GeneID:           "ADIPOQ",
        MolecularWeight:  30.0,
        ExerciseResponse: "Up",
        TemporalPattern:  "Chronic",
        SecretionTriggers: []string{
            "Weight loss",
            "Aerobic training",
            "PPAR-γ activation",
            "Caloric restriction",
        },
        TargetTissues: []string{"Liver", "Skeletal muscle", "Pancreas", "Brain", "Heart"},
        PrimaryEffects: []string{
            "Increased insulin sensitivity",
            "Enhanced fatty acid oxidation",
            "Anti-inflammatory action",
            "AMPK activation",
        },
        SignalingPathways: []string{"AMPK", "PPAR-α", "Ceramidase", "NF-κB"},
        ClinicalRelevance: []string{
            "Type 2 diabetes",
            "Cardiovascular disease",
            "Obesity",
            "Metabolic syndrome",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  5.0,
            Max:  15.0,
            Unit: "μg/mL",
        },
    },

    Leptin = Adipokine{
        Name:             "Leptin",
        GeneID:           "LEP",
        MolecularWeight:  16.0,
        ExerciseResponse: "Down",
        TemporalPattern:  "Both",
        SecretionTriggers: []string{
            "Increased fat mass",
            "Insulin",
            "Glucocorticoids",
            "Inflammatory cytokines",
        },
        TargetTissues: []string{"Hypothalamus", "Pancreas", "Liver", "Skeletal muscle", "Immune cells"},
        PrimaryEffects: []string{
            "Appetite suppression",
            "Energy expenditure",
            "Glucose metabolism",
            "Immune modulation",
        },
        SignalingPathways: []string{"JAK-STAT", "MAPK", "PI3K", "AMPK"},
        ClinicalRelevance: []string{
            "Obesity",
            "Leptin resistance",
            "Inflammation",
            "Reproductive disorders",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  1.0,
            Max:  15.0,
            Unit: "ng/mL",
        },
    },

    Visfatin = Adipokine{
        Name:             "Visfatin/NAMPT",
        GeneID:           "NAMPT",
        MolecularWeight:  52.0,
        ExerciseResponse: "Up",
        TemporalPattern:  "Both",
        SecretionTriggers: []string{
            "Exercise",
            "Inflammatory stimuli",
            "Hyperglycemia",
        },
        TargetTissues: []string{"Pancreas", "Liver", "Skeletal muscle", "Immune cells"},
        PrimaryEffects: []string{
            "NAD biosynthesis",
            "Insulin secretion",
            "SIRT1 activation",
            "Pro-inflammatory action",
        },
        SignalingPathways: []string{"SIRT1", "NF-κB", "MAPK", "PI3K-Akt"},
        ClinicalRelevance: []string{
            "Type 2 diabetes",
            "Obesity",
            "Inflammation",
            "Aging",
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

    Omentin = Adipokine{
        Name:             "Omentin",
        GeneID:           "ITLN1",
        MolecularWeight:  40.0,
        ExerciseResponse: "Up",
        TemporalPattern:  "Chronic",
        SecretionTriggers: []string{
            "Weight loss",
            "Exercise training",
            "Caloric restriction",
        },
        TargetTissues: []string{"Vasculature", "Adipose tissue", "Skeletal muscle"},
        PrimaryEffects: []string{
            "Insulin sensitization",
            "Anti-inflammatory",
            "Vasodilation",
            "Glucose uptake",
        },
        SignalingPathways: []string{"AMPK", "Akt", "eNOS"},
        ClinicalRelevance: []string{
            "Insulin resistance",
            "Obesity",
            "Cardiovascular disease",
            "Polycystic ovary syndrome",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  250.0,
            Max:  800.0,
            Unit: "ng/mL",
        },
    },

    Chemerin = Adipokine{
        Name:             "Chemerin",
        GeneID:           "RARRES2",
        MolecularWeight:  16.0,
        ExerciseResponse: "Down",
        TemporalPattern:  "Chronic",
        SecretionTriggers: []string{
            "Inflammation",
            "Adipogenesis",
            "TNF-α",
            "Insulin",
        },
        TargetTissues: []string{"Adipose tissue", "Liver", "Skeletal muscle", "Immune cells"},
        PrimaryEffects: []string{
            "Adipogenesis",
            "Immune cell migration",
            "Glucose metabolism",
            "Insulin sensitivity",
        },
        SignalingPathways: []string{"CMKLR1/ChemR23", "GPR1", "MAPK", "PI3K-Akt"},
        ClinicalRelevance: []string{
            "Obesity",
            "Metabolic syndrome",
            "Type 2 diabetes",
            "Inflammatory disorders",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  100.0,
            Max:  300.0,
            Unit: "ng/mL",
        },
    },

    Apelin = Adipokine{
        Name:             "Apelin",
        GeneID:           "APLN",
        MolecularWeight:  3.5, // Varies by isoform
        ExerciseResponse: "Up",
        TemporalPattern:  "Both",
        SecretionTriggers: []string{
            "Hypoxia",
            "Insulin",
            "Exercise",
            "TNF-α",
        },
        TargetTissues: []string{"Heart", "Vasculature", "Skeletal muscle", "Adipose tissue"},
        PrimaryEffects: []string{
            "Vasodilation",
            "Insulin sensitivity",
            "Glucose uptake",
            "Cardiac contractility",
        },
        SignalingPathways: []string{"APJ receptor", "MAPK", "PI3K-Akt", "eNOS", "AMPK"},
        ClinicalRelevance: []string{
            "Hypertension",
            "Heart failure",
            "Obesity",
            "Type 2 diabetes",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  200.0,
            Max:  1200.0,
            Unit: "pg/mL",
        },
    },
)

// GetExerciseUpregulatedAdipokines returns adipokines that increase with exercise
func GetExerciseUpregulatedAdipokines() []Adipokine {
    return []Adipokine{Adiponectin, Visfatin, Omentin, Apelin}
}

// GetExerciseDownregulatedAdipokines returns adipokines that decrease with exercise
func GetExerciseDownregulatedAdipokines() []Adipokine {
    return []Adipokine{Leptin, Chemerin}
}

// CalculateAdipokineFoldChange estimates the change in adipokine levels based on exercise
func CalculateAdipokineFoldChange(adipokine Adipokine, exerciseIntensity float64, 
                                 durationMinutes float64, bodyFatPercent float64, 
                                 trainingWeeks int) float64 {
    // Base fold change (1.0 = no change)
    foldChange := 1.0
    
    // Normalize intensity and duration
    intensityFactor := exerciseIntensity / 100.0
    durationFactor := durationMinutes / 60.0
    
    // Body composition influence - adipokines are strongly affected by body fat
    fatMassFactor := 1.0
    if bodyFatPercent > 25 {
        fatMassFactor = 1.0 + ((bodyFatPercent - 25) / 15)
    } else {
        fatMassFactor = 1.0 - ((25 - bodyFatPercent) / 25)
    }
    if fatMassFactor < 0.6 {
        fatMassFactor = 0.6
    }
    
    // Chronic adaptation factor
    chronicFactor := float64(trainingWeeks) / 12.0
    if chronicFactor > 1.0 {
        chronicFactor = 1.0
    }
    
    // Calculate response based on specific adipokine
    switch adipokine.Name {
    case "Adiponectin":
        // Adiponectin increases with chronic training, more with lower body fat
        acuteChange := 1.0 + (intensityFactor * 0.05) // Minimal acute effect
        chronicChange := 1.0 + (chronicFactor * 0.3 / fatMassFactor) // Greater effect with lower fat
        foldChange = acuteChange * chronicChange
        
    case "Leptin":
        // Leptin decreases with exercise, both acutely and chronically
        acuteChange := 1.0 - (intensityFactor * durationFactor * 0.15)
        chronicChange := 1.0 - (chronicFactor * 0.25 * fatMassFactor) // Greater decrease with higher fat
        foldChange = acuteChange * chronicChange
        if foldChange < 0.5 {
            foldChange = 0.5 // Limit the decrease
        }
        
    case "Visfatin/NAMPT":
        // Visfatin increases with exercise
        acuteChange := 1.0 + (intensityFactor * 0.2)
        chronicChange := 1.0 + (chronicFactor * 0.15)
        foldChange = acuteChange * chronicChange
        
    case "Omentin":
        // Omentin increases with chronic exercise, especially with fat loss
        acuteChange := 1.0 + (intensityFactor * 0.05) // Small acute effect
        chronicChange := 1.0 + (chronicFactor * 0.2 / fatMassFactor) // Greater with lower fat
        foldChange = acuteChange * chronicChange
        
    case "Chemerin":
        // Chemerin decreases with exercise training
        acuteChange := 1.0 - (intensityFactor * 0.05) // Small acute effect
        chronicChange := 1.0 - (chronicFactor * 0.2 * fatMassFactor) // Greater decrease with higher fat
        foldChange = acuteChange * chronicChange
        if foldChange < 0.6 {
            foldChange = 0.6 // Limit the decrease
        }
        
    case "Apelin":
        // Apelin increases with exercise
        acuteChange := 1.0 + (intensityFactor * durationFactor * 0.3)
        chronicChange := 1.0 + (chronicFactor * 0.15)
        foldChange = acuteChange * chronicChange
    }
    
    return foldChange
}

// PredictAdipokineMetabolicEffects models the metabolic effects of changes in adipokine levels
func PredictAdipokineMetabolicEffects(adipokine Adipokine, 
                                     foldChange float64) map[string]float64 {
    // Initialize effects map
    effects := make(map[string]float64)
    
    // Calculate effect magnitude (normalized change from baseline)
    var effectMagnitude float64
    if foldChange >= 1.0 {
        effectMagnitude = foldChange - 1.0 // For increases
    } else {
        effectMagnitude = 1.0 - foldChange // For decreases
    }
    
    // Scale effect to 0-10 range for consistency
    scaledEffect := effectMagnitude * 10
    if scaledEffect > 10.0 {
        scaledEffect = 10.0
    }
    
    // Direction of change
    increasing := foldChange >= 1.0
    
    // Assign effects based on adipokine and direction of change
    switch adipokine.Name {
    case "Adiponectin":
        if increasing {
            effects["Insulin sensitivity"] = scaledEffect * 0.8
            effects["Fatty acid oxidation"] = scaledEffect * 0.7
            effects["Anti-inflammatory action"] = scaledEffect * 0.6
            effects["Glucose uptake"] = scaledEffect * 0.5
        } else {
            effects["Insulin resistance"] = scaledEffect * 0.8
            effects["Lipid accumulation"] = scaledEffect * 0.7
            effects["Inflammatory state"] = scaledEffect * 0.6
        }
        
    case "Leptin":
        if increasing {
            effects["Appetite suppression"] = scaledEffect * 0.7
            effects["Energy expenditure"] = scaledEffect * 0.6
            effects["Inflammation"] = scaledEffect * 0.5
        } else {
            effects["Improved leptin sensitivity"] = scaledEffect * 0.7
            effects["Reduced inflammation"] = scaledEffect * 0.6
            effects["Metabolic flexibility"] = scaledEffect * 0.5
        }
        
    case "Visfatin/NAMPT":
        if increasing {
            effects["NAD biosynthesis"] = scaledEffect * 0.8
            effects["SIRT1 activation"] = scaledEffect * 0.7
            effects["Insulin secretion"] = scaledEffect * 0.5
            effects["Inflammatory response"] = scaledEffect * 0.4 // Can be pro-inflammatory
        } else {
            effects["Reduced NAD availability"] = scaledEffect * 0.8
            effects["Decreased SIRT1 activity"] = scaledEffect * 0.7
        }
        
    case "Omentin":
        if increasing {
            effects["Insulin sensitivity"] = scaledEffect * 0.8
            effects["Anti-inflammatory action"] = scaledEffect * 0.7
            effects["Vasodilation"] = scaledEffect * 0.6
            effects["Glucose uptake"] = scaledEffect * 0.5
        } else {
            effects["Insulin resistance"] = scaledEffect * 0.8
            effects["Inflammatory state"] = scaledEffect * 0.7
            effects["Vascular dysfunction"] = scaledEffect * 0.6
        }
        
    case "Chemerin":
        if increasing {
            effects["Immune cell recruitment"] = scaledEffect * 0.8
            effects["Adipogenesis"] = scaledEffect * 0.7
            effects["Insulin resistance"] = scaledEffect * 0.6
            effects["Inflammatory response"] = scaledEffect * 0.7
        } else {
            effects["Reduced inflammation"] = scaledEffect * 0.8
            effects["Improved insulin sensitivity"] = scaledEffect * 0.7
            effects["Decreased immune infiltration"] = scaledEffect * 0.6
        }
        
    case "Apelin":
        if increasing {
            effects["Glucose uptake"] = scaledEffect * 0.8
            effects["Vasodilation"] = scaledEffect * 0.7
            effects["Insulin sensitivity"] = scaledEffect * 0.6
            effects["Cardiac function"] = scaledEffect * 0.5
            effects["Nitric oxide production"] = scaledEffect * 0.7
        } else {
            effects["Impaired glucose handling"] = scaledEffect * 0.8
            effects["Reduced vasodilation"] = scaledEffect * 0.7
            effects["Decreased cardiac function"] = scaledEffect * 0.5
        }
    }
    
    return effects
}