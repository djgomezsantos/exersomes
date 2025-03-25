package placenta

// Hepatokine represents a signaling molecule secreted from liver tissue
type Hepatokine struct {
    Name              string
    GeneID            string
    MolecularWeight   float64  // In kDa
    ExerciseResponse  string   // "Up", "Down", "Biphasic", "Unknown"
    TemporalPattern   string   // "Acute", "Chronic", "Both", "Unknown"
    SecretionTriggers []string // Factors triggering secretion from liver
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

// Collection of key exercise-responsive hepatokines
var (
    FGF21 = Hepatokine{
        Name:             "Fibroblast Growth Factor 21 (Liver-derived)",
        GeneID:           "FGF21",
        MolecularWeight:  22.0,
        ExerciseResponse: "Up",
        TemporalPattern:  "Both",
        SecretionTriggers: []string{
            "Fasting", 
            "High-intensity exercise",
            "Endurance exercise",
            "PPARα activation",
            "Protein restriction",
            "Mitochondrial stress",
        },
        TargetTissues: []string{"Adipose tissue", "Brain", "Pancreas", "Skeletal muscle", "Heart"},
        PrimaryEffects: []string{
            "Glucose homeostasis",
            "Lipid metabolism",
            "Ketogenesis",
            "Insulin sensitivity",
            "Energy expenditure",
            "Weight regulation",
        },
        SignalingPathways: []string{"FGFR1c/β-Klotho", "ERK1/2", "PI3K/Akt", "AMPK", "SIRT1"},
        ClinicalRelevance: []string{
            "Type 2 diabetes",
            "Obesity",
            "Non-alcoholic fatty liver disease",
            "Metabolic syndrome",
            "Cardiovascular disease",
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

    FetuinA = Hepatokine{
        Name:             "Fetuin-A",
        GeneID:           "AHSG",
        MolecularWeight:  64.0,
        ExerciseResponse: "Down",
        TemporalPattern:  "Chronic",
        SecretionTriggers: []string{
            "Palmitate exposure",
            "Inflammatory cytokines",
            "Hepatic steatosis",
            "Insulin resistance",
        },
        TargetTissues: []string{"Skeletal muscle", "Adipose tissue", "Pancreas", "Kidney"},
        PrimaryEffects: []string{
            "Insulin receptor inhibition",
            "TLR4 activation",
            "Adipose tissue inflammation",
            "Calcium phosphate inhibition",
            "Vascular calcification prevention",
        },
        SignalingPathways: []string{"TLR4/NF-κB", "JNK", "Insulin receptor/IRS-1"},
        ClinicalRelevance: []string{
            "Type 2 diabetes",
            "Insulin resistance",
            "Non-alcoholic fatty liver disease",
            "Cardiovascular disease",
            "Chronic kidney disease",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  200.0,
            Max:  600.0,
            Unit: "μg/mL",
        },
    },

    ANGPTL4 = Hepatokine{
        Name:             "Angiopoietin-like Protein 4",
        GeneID:           "ANGPTL4",
        MolecularWeight:  45.0,
        ExerciseResponse: "Up",
        TemporalPattern:  "Both",
        SecretionTriggers: []string{
            "Fasting",
            "Exercise",
            "Fatty acids",
            "PPARα/δ activation",
            "Hypoxia",
            "Glucocorticoids",
        },
        TargetTissues: []string{"Adipose tissue", "Skeletal muscle", "Heart", "Vasculature", "Intestine"},
        PrimaryEffects: []string{
            "Lipoprotein lipase inhibition",
            "Triglyceride metabolism",
            "Fat storage regulation",
            "Angiogenesis modulation",
            "Energy homeostasis",
            "Lipid partitioning",
        },
        SignalingPathways: []string{"PPAR signaling", "HIF-1α", "AKT/mTOR"},
        ClinicalRelevance: []string{
            "Dyslipidemia",
            "Obesity",
            "Type 2 diabetes",
            "Cardiovascular disease",
            "Exercise metabolism",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  2.0,
            Max:  8.0,
            Unit: "ng/mL",
        },
    },

    SelP = Hepatokine{
        Name:             "Selenoprotein P",
        GeneID:           "SELENOP",
        MolecularWeight:  42.0,
        ExerciseResponse: "Down",
        TemporalPattern:  "Chronic",
        SecretionTriggers: []string{
            "High selenium intake",
            "Inflammatory cytokines",
            "Hyperglycemia",
            "Oxidative stress",
        },
        TargetTissues: []string{"Brain", "Testes", "Skeletal muscle", "Pancreas"},
        PrimaryEffects: []string{
            "Selenium transport",
            "Antioxidant action",
            "AMPK inhibition",
            "Insulin signaling disruption",
            "Skeletal muscle insulin resistance",
        },
        SignalingPathways: []string{"AMPK", "PGC-1α", "Insulin/IRS", "ROS pathways"},
        ClinicalRelevance: []string{
            "Type 2 diabetes",
            "Insulin resistance",
            "Metabolic syndrome",
            "Sarcopenia",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  2.0,
            Max:  6.0,
            Unit: "μg/mL",
        },
    },

    Follistatin = Hepatokine{
        Name:             "Follistatin",
        GeneID:           "FST",
        MolecularWeight:  35.0,
        ExerciseResponse: "Up",
        TemporalPattern:  "Acute",
        SecretionTriggers: []string{
            "Acute exercise",
            "Resistance training",
            "Inflammatory cytokines",
            "Hepatic stress",
        },
        TargetTissues: []string{"Skeletal muscle", "Adipose tissue", "Pancreas", "Gonads"},
        PrimaryEffects: []string{
            "Myostatin antagonism",
            "Muscle hypertrophy",
            "TGF-β inhibition",
            "Insulin secretion modulation",
            "Glucose homeostasis",
        },
        SignalingPathways: []string{"Activin/Myostatin", "SMAD2/3", "Akt/mTOR", "MAPK"},
        ClinicalRelevance: []string{
            "Sarcopenia",
            "Cachexia",
            "Type 2 diabetes",
            "Muscle wasting disorders",
            "Reproductive disorders",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  1.0,
            Max:  3.0,
            Unit: "ng/mL",
        },
    },

    IGFBP1 = Hepatokine{
        Name:             "Insulin-like Growth Factor Binding Protein 1",
        GeneID:           "IGFBP1",
        MolecularWeight:  25.0,
        ExerciseResponse: "Up",
        TemporalPattern:  "Acute",
        SecretionTriggers: []string{
            "Fasting",
            "Acute exercise",
            "Insulin deficiency",
            "Glucocorticoids",
            "Inflammatory cytokines",
        },
        TargetTissues: []string{"Liver", "Skeletal muscle", "Adipose tissue", "Vasculature"},
        PrimaryEffects: []string{
            "IGF-1 bioavailability regulation",
            "Glucose metabolism",
            "Cell growth modulation",
            "Insulin sensitivity",
            "Cell survival",
        },
        SignalingPathways: []string{"IGF-1R", "Integrin", "FAK", "mTOR"},
        ClinicalRelevance: []string{
            "Type 2 diabetes",
            "Insulin resistance",
            "Growth disorders",
            "Exercise adaptation",
            "Metabolic health",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  20.0,
            Max:  50.0,
            Unit: "ng/mL",
        },
    },

    Hepassocin = Hepatokine{
        Name:             "Hepassocin",
        GeneID:           "FGL1",
        MolecularWeight:  35.0,
        ExerciseResponse: "Up",
        TemporalPattern:  "Chronic",
        SecretionTriggers: []string{
            "Liver regeneration signals",
            "Chronic exercise",
            "Hepatic stress",
            "Low-grade inflammation",
        },
        TargetTissues: []string{"Liver", "Adipose tissue", "Skeletal muscle"},
        PrimaryEffects: []string{
            "Hepatocyte proliferation",
            "Liver regeneration",
            "Fat accumulation reduction",
            "Insulin sensitivity improvement",
            "Mitochondrial function",
        },
        SignalingPathways: []string{"EGFR", "Akt", "STAT3", "ERK1/2"},
        ClinicalRelevance: []string{
            "Non-alcoholic fatty liver disease",
            "Liver injury",
            "Type 2 diabetes",
            "Metabolic syndrome",
        },
        BaselineRange: struct {
            Min  float64
            Max  float64
            Unit string
        }{
            Min:  10.0,
            Max:  40.0,
            Unit: "ng/mL",
        },
    },
)

// GetExerciseResponsiveHepatokines returns a slice of hepatokines that respond to exercise
func GetExerciseResponsiveHepatokines() []Hepatokine {
    return []Hepatokine{
        FGF21,
        FetuinA,
        ANGPTL4,
        SelP,
        Follistatin,
        IGFBP1,
        Hepassocin,
    }
}

// GetHepatokineByName returns a hepatokine by its name
func GetHepatokineByName(name string) (Hepatokine, bool) {
    for _, hepatokine := range GetExerciseResponsiveHepatokines() {
        if hepatokine.Name == name {
            return hepatokine, true
        }
    }
    return Hepatokine{}, false
}

// FilterHepatokinesByExerciseResponse returns hepatokines with a specific exercise response
func FilterHepatokinesByExerciseResponse(response string) []Hepatokine {
    var filtered []Hepatokine
    for _, hepatokine := range GetExerciseResponsiveHepatokines() {
        if hepatokine.ExerciseResponse == response {
            filtered = append(filtered, hepatokine)
        }
    }
    return filtered
}

// CalculateAcuteExerciseResponse predicts the relative change in hepatokine levels 
// after an acute exercise bout based on intensity and duration
func CalculateAcuteExerciseResponse(hepatokine Hepatokine, 
    intensityPercent int, durationMinutes int) float64 {
    
    // Default no change
    responseMultiplier := 1.0
    
    // Only process if hepatokine responds acutely to exercise
    if hepatokine.TemporalPattern != "Acute" && hepatokine.TemporalPattern != "Both" {
        return responseMultiplier
    }
    
    // Calculate intensity factor (0.0-2.0 scale)
    intensityFactor := float64(intensityPercent) / 50.0
    
    // Calculate duration factor (normalized to 60 minutes)
    durationFactor := float64(durationMinutes) / 60.0
    if durationFactor > 2.0 {
        durationFactor = 2.0 // Cap very long exercise
    }
    
    // Calculate response based on specific hepatokines
    switch hepatokine.Name {
    case "Fibroblast Growth Factor 21 (Liver-derived)":
        // Strong response to high intensity exercise
        responseMultiplier = 1.0 + (intensityFactor * 1.8 * durationFactor)
        
    case "Angiopoietin-like Protein 4":
        // Responds strongly to endurance exercise
        if intensityPercent < 70 {
            responseMultiplier = 1.0 + (0.7 * durationFactor)
        } else {
            responseMultiplier = 1.0 + (intensityFactor * 0.9 * durationFactor)
        }
        
    case "Follistatin":
        // Responds more strongly to high intensity
        responseMultiplier = 1.0 + (intensityFactor * intensityFactor * durationFactor)
        
    case "Insulin-like Growth Factor Binding Protein 1":
        // Rapid response to exercise, especially with longer duration
        responseMultiplier = 1.0 + (intensityFactor * 0.8 * durationFactor * 1.2)
        
    default:
        // Generic modest response for other hepatokines
        responseMultiplier = 1.0 + (intensityFactor * 0.4 * durationFactor)
    }
    
    return responseMultiplier
}