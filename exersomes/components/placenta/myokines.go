package placenta

// Myokine represents a signaling molecule secreted from skeletal muscle
type Myokine struct {
    Name              string
    GeneID            string
    MolecularWeight   float64  // In kDa
    ExerciseResponse  string   // "Up", "Down", "Biphasic"
    TemporalPattern   string   // "Acute", "Chronic", "Both"
    SecretionTriggers []string // Factors triggering the release
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

// Collection of key exercise-induced myokines
var (
    IL6 = Myokine{
        Name:             "Interleukin-6",
        GeneID:           "IL6",
        MolecularWeight:  21.0,
        ExerciseResponse: "Up",
        TemporalPattern:  "Acute",
        SecretionTriggers: []string{
            "Glycogen depletion",
            "Calcium signaling",
            "ROS",
            "Muscle damage",
        },
        TargetTissues: []string{"Liver", "Adipose tissue", "Immune cells", "Brain", "Muscle itself"},
        PrimaryEffects: []string{
            "Hepatic glucose production",
            "Lipolysis in adipose tissue",
            "Anti-inflammatory effects",
            "Insulin secretion modulation",
        },
        SignalingPathways: []string{"JAK-STAT", "MAPK", "PI3K-Akt"},
        ClinicalRelevance: []string{
            "Metabolic disorders",
            "Inflammation",
            "Insulin sensitivity",
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

    BDNF = Myokine{
        Name:             "Brain-Derived Neurotrophic Factor",
        GeneID:           "BDNF",
        MolecularWeight:  13.6,
        ExerciseResponse: "Up",
        TemporalPattern:  "Both",
        SecretionTriggers: []string{
            "AMPK activation",
            "PGC-1α induction",
            "Calcium signaling",
            "Aerobic exercise",
        },
        TargetTissues: []string{"Brain", "Spinal cord", "Peripheral nerves", "Muscle itself"},
        PrimaryEffects: []string{
            "Neuroplasticity",
            "Cognitive function",
            "Fat oxidation",
            "Muscle regeneration",
        },
        SignalingPathways: []string{"TrkB", "p75NTR", "MAPK", "PI3K-Akt"},
        ClinicalRelevance: []string{
            "Depression",
            "Neurodegenerative diseases",
            "Cognitive decline",
            "Metabolic health",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  8.0,
            Max:  25.0,
            Unit: "ng/mL",
        },
    },

    Irisin = Myokine{
        Name:             "Irisin",
        GeneID:           "FNDC5",
        MolecularWeight:  12.0,
        ExerciseResponse: "Up",
        TemporalPattern:  "Both",
        SecretionTriggers: []string{
            "PGC-1α activation",
            "Endurance exercise",
            "Cold exposure",
            "Resistance training",
        },
        TargetTissues: []string{"White adipose tissue", "Bone", "Brain", "Liver"},
        PrimaryEffects: []string{
            "Browning of white fat",
            "Energy expenditure",
            "Bone formation",
            "Neuroprotection",
        },
        SignalingPathways: []string{"P38 MAPK", "ERK", "AMPK"},
        ClinicalRelevance: []string{
            "Obesity",
            "Type 2 diabetes",
            "Osteoporosis",
            "Neurodegeneration",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  0.2,
            Max:  2.0,
            Unit: "μg/mL",
        },
    },

    MyostatinInh = Myokine{
        Name:             "Myostatin Inhibitors (Follistatin, FSTL-1)",
        GeneID:           "FST/FSTL1",
        MolecularWeight:  35.0, // Follistatin
        ExerciseResponse: "Up",
        TemporalPattern:  "Both",
        SecretionTriggers: []string{
            "Resistance exercise",
            "Muscle damage",
            "Eccentric contractions",
        },
        TargetTissues: []string{"Skeletal muscle", "Adipose tissue", "Liver"},
        PrimaryEffects: []string{
            "Myostatin inhibition",
            "Muscle hypertrophy",
            "Satellite cell activation",
            "Fat metabolism",
        },
        SignalingPathways: []string{"Activin receptor", "SMAD2/3", "mTOR"},
        ClinicalRelevance: []string{
            "Sarcopenia",
            "Muscle wasting diseases",
            "Aging",
            "Obesity",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  0.5,
            Max:  2.0,
            Unit: "ng/mL",
        },
    },

    Decorin = Myokine{
        Name:             "Decorin",
        GeneID:           "DCN",
        MolecularWeight:  90.0,
        ExerciseResponse: "Up",
        TemporalPattern:  "Chronic",
        SecretionTriggers: []string{
            "Resistance exercise",
            "Mechanical loading",
            "Muscle remodeling",
        },
        TargetTissues: []string{"Muscle ECM", "Myoblasts", "Connective tissue", "Tumor cells"},
        PrimaryEffects: []string{
            "Myostatin inhibition",
            "ECM remodeling",
            "Muscle hypertrophy",
            "Anti-fibrotic",
        },
        SignalingPathways: []string{"Myostatin", "TGF-β", "EGFR", "IGF-1R", "VEGFR2"},
        ClinicalRelevance: []string{
            "Muscle adaptation",
            "Fibrosis",
            "Tumor suppression",
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

    IL15 = Myokine{
        Name:             "Interleukin-15",
        GeneID:           "IL15",
        MolecularWeight:  14.0,
        ExerciseResponse: "Up",
        TemporalPattern:  "Both",
        SecretionTriggers: []string{
            "Resistance exercise",
            "Muscle contraction",
            "Strength training",
        },
        TargetTissues: []string{"Skeletal muscle", "Adipose tissue", "Immune cells"},
        PrimaryEffects: []string{
            "Muscle anabolism",
            "Fat mass reduction",
            "Oxidative metabolism",
            "Immune cell development",
        },
        SignalingPathways: []string{"JAK-STAT", "MAPK", "PI3K-Akt"},
        ClinicalRelevance: []string{
            "Sarcopenia",
            "Obesity",
            "Inflammation",
            "Metabolic diseases",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  1.0,
            Max:  5.0,
            Unit: "pg/mL",
        },
    },

    Apelin = Myokine{
        Name:             "Apelin",
        GeneID:           "APLN",
        MolecularWeight:  3.5, // Varies by isoform
        ExerciseResponse: "Up",
        TemporalPattern:  "Both",
        SecretionTriggers: []string{
            "Endurance exercise",
            "AMPK activation",
            "PGC-1α signaling",
        },
        TargetTissues: []string{"Cardiovascular system", "Skeletal muscle", "Adipose tissue", "Brain"},
        PrimaryEffects: []string{
            "Glucose utilization",
            "Mitochondrial biogenesis",
            "Vasodilation",
            "Insulin sensitivity",
        },
        SignalingPathways: []string{"APJ receptor", "AMPK", "PI3K-Akt", "eNOS"},
        ClinicalRelevance: []string{
            "Cardiovascular health",
            "Insulin resistance",
            "Hypertension",
            "Obesity",
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

// GetEnduranceExerciseMyokines returns myokines primarily upregulated by endurance exercise
func GetEnduranceExerciseMyokines() []Myokine {
    return []Myokine{IL6, BDNF, Irisin, Apelin}
}

// GetResistanceExerciseMyokines returns myokines primarily upregulated by resistance exercise
func GetResistanceExerciseMyokines() []Myokine {
    return []Myokine{MyostatinInh, Decorin, IL15}
}

// CalculateMyokineResponse estimates myokine changes during/after exercise
func CalculateMyokineResponse(myokine Myokine, exerciseType string, 
                            intensityPercent int, durationMinutes int, 
                            trainingStatus string) float64 {
    // Base fold change (1.0 = no change from baseline)
    foldChange := 1.0
    
    // Convert parameters to normalized values
    intensity := float64(intensityPercent) / 100.0
    duration := float64(durationMinutes) / 60.0 // Duration in hours
    
    // Training status modifier (untrained individuals often show greater acute responses)
    trainingMod := 1.0
    if trainingStatus == "Untrained" {
        trainingMod = 1.4
    } else if trainingStatus == "Highly trained" {
        trainingMod = 0.8
    }
    
    // Exercise type modifier
    exerciseMod := 1.0
    
    // Calculate response based on myokine and exercise characteristics
    switch myokine.Name {
    case "Interleukin-6":
        // IL-6 strongly responds to glycogen-depleting endurance exercise
        if exerciseType == "Endurance" {
            exerciseMod = 1.5
            foldChange = 1.0 + (intensity * duration * 4.0 * exerciseMod * trainingMod)
        } else if exerciseType == "HIIT" {
            exerciseMod = 1.3
            foldChange = 1.0 + (intensity * duration * 3.0 * exerciseMod * trainingMod)
        } else if exerciseType == "Resistance" {
            exerciseMod = 0.7
            foldChange = 1.0 + (intensity * 1.5 * exerciseMod * trainingMod)
        }
        
    case "Brain-Derived Neurotrophic Factor":
        // BDNF responds best to moderate-to-high intensity aerobic exercise
        if exerciseType == "Endurance" || exerciseType == "HIIT" {
            exerciseMod = 1.3
            foldChange = 1.0 + (intensity * 0.8 * exerciseMod * trainingMod)
        } else {
            exerciseMod = 0.6
            foldChange = 1.0 + (intensity * 0.4 * exerciseMod * trainingMod)
        }
        
    case "Irisin":
        // Irisin responds to both resistance and endurance, especially high intensity
        if intensity > 0.7 { // High intensity effect
            foldChange = 1.0 + (intensity * 0.7 * trainingMod)
        } else {
            foldChange = 1.0 + (intensity * duration * 0.4 * trainingMod)
        }
        
    case "Myostatin Inhibitors (Follistatin, FSTL-1)":
        // Follistatin responds best to resistance training
        if exerciseType == "Resistance" {
            exerciseMod = 1.5
            foldChange = 1.0 + (intensity * 1.2 * exerciseMod * trainingMod)
        } else if exerciseType == "HIIT" {
            exerciseMod = 1.0
            foldChange = 1.0 + (intensity * 0.7 * exerciseMod * trainingMod)
        } else {
            exerciseMod = 0.5
            foldChange = 1.0 + (intensity * 0.3 * exerciseMod * trainingMod)
        }
        
    case "Decorin":
        // Decorin responds primarily to resistance training
        if exerciseType == "Resistance" {
            exerciseMod = 1.4
            foldChange = 1.0 + (intensity * 0.6 * exerciseMod * trainingMod)
        } else {
            exerciseMod = 0.3
            foldChange = 1.0 + (intensity * 0.2 * exerciseMod * trainingMod)
        }
        
    case "Interleukin-15":
        // IL-15 responds best to resistance exercise
        if exerciseType == "Resistance" {
            exerciseMod = 1.3
            foldChange = 1.0 + (intensity * 0.8 * exerciseMod * trainingMod)
        } else if exerciseType == "HIIT" {
            exerciseMod = 0.7
            foldChange = 1.0 + (intensity * 0.5 * exerciseMod * trainingMod)
        } else {
            exerciseMod = 0.4
            foldChange = 1.0 + (intensity * 0.3 * exerciseMod * trainingMod)
        }
        
    case "Apelin":
        // Apelin responds to both types but better to endurance
        if exerciseType == "Endurance" || exerciseType == "HIIT" {
            exerciseMod = 1.2
            foldChange = 1.0 + (intensity * duration * 0.7 * exerciseMod * trainingMod)
        } else {
            exerciseMod = 0.7
            foldChange = 1.0 + (intensity * 0.5 * exerciseMod * trainingMod)
        }
    }
    
    return foldChange
}

// PredictMyokineEffects models the downstream physiological effects of myokine increases
func PredictMyokineEffects(myokine Myokine, 
                          responseMagnitude float64) map[string]float64 {
    
    // Initialize effects map
    effects := make(map[string]float64)
    
    // Calculate effect magnitude (normalized response above baseline)
    effectMagnitude := responseMagnitude - 1.0
    if effectMagnitude < 0 {
        effectMagnitude = 0
    }
    
    // Scale effect to 0-10 range for consistency
    scaledEffect := effectMagnitude * 10
    
    // Assign effects based on myokine
    switch myokine.Name {
    case "Interleukin-6":
        effects["Hepatic glucose production"] = scaledEffect * 0.8
        effects["Lipolysis"] = scaledEffect * 0.7
        effects["Anti-inflammatory signaling"] = scaledEffect * 0.6
        effects["Insulin sensitivity"] = scaledEffect * 0.5
        
    case "Brain-Derived Neurotrophic Factor":
        effects["Neuroplasticity"] = scaledEffect * 0.9
        effects["Cognitive function"] = scaledEffect * 0.8
        effects["Fat oxidation"] = scaledEffect * 0.5
        effects["Neurotrophic support"] = scaledEffect * 0.7
        
    case "Irisin":
        effects["WAT browning"] = scaledEffect * 0.7
        effects["Energy expenditure"] = scaledEffect * 0.8
        effects["Glucose homeostasis"] = scaledEffect * 0.6
        effects["Bone mineralization"] = scaledEffect * 0.5
        
    case "Myostatin Inhibitors (Follistatin, FSTL-1)":
        effects["Muscle hypertrophy"] = scaledEffect * 0.9
        effects["Satellite cell activation"] = scaledEffect * 0.8
        effects["Myostatin inhibition"] = scaledEffect * 1.0
        effects["Anti-fibrotic action"] = scaledEffect * 0.6
        
    case "Decorin":
        effects["Myostatin binding"] = scaledEffect * 0.8
        effects["ECM remodeling"] = scaledEffect * 0.9
        effects["Hypertrophy signaling"] = scaledEffect * 0.7
        effects["Collagen organization"] = scaledEffect * 0.8
        
    case "Interleukin-15":
        effects["Muscle anabolism"] = scaledEffect * 0.7
        effects["Adipose tissue reduction"] = scaledEffect * 0.6
        effects["Oxidative metabolism"] = scaledEffect * 0.5
        effects["NK cell development"] = scaledEffect * 0.4
        
    case "Apelin":
        effects["Glucose uptake"] = scaledEffect * 0.8
        effects["Mitochondrial biogenesis"] = scaledEffect * 0.6
        effects["Vasodilation"] = scaledEffect * 0.7
        effects["Cardiac function"] = scaledEffect * 0.6
        effects["Insulin sensitivity"] = scaledEffect * 0.7
    }
    
    return effects
}