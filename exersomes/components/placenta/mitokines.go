package placenta

// Mitokine represents a signaling molecule produced by or in response to mitochondrial activity
type Mitokine struct {
    Name                 string
    Type                 string   // "Peptide", "Metabolite", "Regulator", etc.
    MolecularWeight      float64  // in kDa
    SourceTissues        []string // Primary tissues producing this factor
    TargetTissues        []string // Tissues responding to this factor
    Production           string   // "Mitochondrial stress", "Bioenergetic demand", etc.
    PrimaryPathways      []string // Key signaling pathways involved
    ExerciseRegulation   string   // "Up", "Down", "Biphasic", "Context-dependent"
    ExerciseTimeCourse   string   // "Acute", "Chronic", "Both"
    PhysiologicalEffects []string // Major physiological effects
    BioenergeticRole     string   // Role in energy metabolism
}

// Collection of key mitochondrial signaling factors
var (
    PGC1a = Mitokine{
        Name:              "Peroxisome Proliferator-Activated Receptor Gamma Coactivator 1-Alpha",
        Type:              "Transcriptional Coactivator",
        MolecularWeight:   91.0,
        SourceTissues:     []string{"Skeletal muscle", "Brown adipose tissue", "Heart", "Brain", "Kidney"},
        TargetTissues:     []string{"Self-regulation", "Mitochondria", "Nucleus"},
        Production:         "Exercise-induced AMPK & p38 MAPK activation, β-adrenergic signaling",
        PrimaryPathways:    []string{"Mitochondrial biogenesis", "Oxidative metabolism", "Antioxidant defense"},
        ExerciseRegulation: "Up",
        ExerciseTimeCourse: "Both",
        PhysiologicalEffects: []string{
            "Mitochondrial biogenesis",
            "Metabolic substrate shifting",
            "Antioxidant enzyme upregulation",
            "Angiogenesis",
            "Fiber-type transformation",
        },
        BioenergeticRole: "Master regulator of mitochondrial biogenesis and function",
    },

    TFAM = Mitokine{
        Name:              "Mitochondrial Transcription Factor A",
        Type:              "Transcription Factor",
        MolecularWeight:   29.0,
        SourceTissues:     []string{"All tissues with mitochondria"},
        TargetTissues:     []string{"Mitochondria"},
        Production:         "Downstream of PGC-1α activation",
        PrimaryPathways:    []string{"mtDNA transcription", "mtDNA maintenance"},
        ExerciseRegulation: "Up",
        ExerciseTimeCourse: "Chronic",
        PhysiologicalEffects: []string{
            "mtDNA replication",
            "mtDNA transcription",
            "mtDNA packaging and protection",
            "Regulation of mitochondrial gene expression",
        },
        BioenergeticRole: "Essential for mitochondrial gene expression and mtDNA maintenance",
    },

    FGF21 = Mitokine{
        Name:              "Fibroblast Growth Factor 21",
        Type:              "Hormone",
        MolecularWeight:   22.0,
        SourceTissues:     []string{"Liver", "Muscle", "BAT", "WAT"},
        TargetTissues:     []string{"Liver", "Adipose tissue", "Pancreas", "Brain"},
        Production:         "Mitochondrial stress, unfolded protein response, nutrient deprivation",
        PrimaryPathways:    []string{"β-Klotho/FGFR1", "AMPK", "PGC-1α", "Akt"},
        ExerciseRegulation: "Up",
        ExerciseTimeCourse: "Both",
        PhysiologicalEffects: []string{
            "Glucose homeostasis",
            "Lipid metabolism",
            "Mitochondrial function",
            "Thermogenesis in BAT",
            "Stress resilience",
        },
        BioenergeticRole: "Metabolic stress hormone facilitating adaptations to energetic challenges",
    },

    GDF15 = Mitokine{
        Name:              "Growth Differentiation Factor 15",
        Type:              "Cytokine",
        MolecularWeight:   35.0,
        SourceTissues:     []string{"Muscle", "Heart", "Liver", "Kidney", "Macrophages"},
        TargetTissues:     []string{"Brain", "Systemic targets", "Heart"},
        Production:         "Mitochondrial stress, integrated stress response (ISR), inflammation",
        PrimaryPathways:    []string{"GFRAL-RET", "eIF2α-ATF4", "p53"},
        ExerciseRegulation: "Up",
        ExerciseTimeCourse: "Acute",
        PhysiologicalEffects: []string{
            "Energy intake regulation",
            "Cardioprotection",
            "Anti-inflammatory",
            "Metabolic regulation",
        },
        BioenergeticRole: "Stress-induced mitokine signaling mitochondrial homeostatic challenge",
    },

    Lactate = Mitokine{
        Name:              "Lactate",
        Type:              "Metabolite",
        MolecularWeight:   0.09, // 90 Da
        SourceTissues:     []string{"Skeletal muscle", "Red blood cells", "Brain", "Skin"},
        TargetTissues:     []string{"Liver", "Heart", "Brain", "Other skeletal muscles"},
        Production:         "Anaerobic glycolysis, high-intensity exercise",
        PrimaryPathways:    []string{"MCT transporters", "GPR81 receptor", "Redox signaling"},
        ExerciseRegulation: "Up",
        ExerciseTimeCourse: "Acute",
        PhysiologicalEffects: []string{
            "Alternative fuel source",
            "Signaling molecule",
            "Glucose sparing",
            "Angiogenesis promotion",
            "Gene expression regulation",
        },
        BioenergeticRole: "Metabolic intermediate and signaling molecule in energy substrate shuttling",
    },

    mtROS = Mitokine{
        Name:              "Mitochondrial Reactive Oxygen Species",
        Type:              "Metabolite/Signaling molecules",
        MolecularWeight:   0.0, // Various species
        SourceTissues:     []string{"All tissues with mitochondria"},
        TargetTissues:     []string{"Local and systemic targets", "Nucleus"},
        Production:         "Electron transport chain activity, mitochondrial respiration",
        PrimaryPathways:    []string{"NF-κB", "MAPK", "Nrf2-Keap1", "HIF-1α"},
        ExerciseRegulation: "Biphasic",
        ExerciseTimeCourse: "Both",
        PhysiologicalEffects: []string{
            "Redox signaling",
            "Antioxidant enzyme induction",
            "Mitochondrial hormesis",
            "Mitochondrial biogenesis",
            "Stress adaptation",
        },
        BioenergeticRole: "Mitohormetic signal linking mitochondrial activity to cellular adaptation",
    },

    ATP = Mitokine{
        Name:              "Extracellular ATP",
        Type:              "Nucleotide",
        MolecularWeight:   0.507, // 507 Da
        SourceTissues:     []string{"Exercising muscle", "Stressed cells", "Platelets", "Neurons"},
        TargetTissues:     []string{"Vasculature", "Immune cells", "Neurons", "Various tissues"},
        Production:         "Cell stress, damage, or controlled release",
        PrimaryPathways:    []string{"P2X receptors", "P2Y receptors", "AMPK (after conversion to adenosine)"},
        ExerciseRegulation: "Up",
        ExerciseTimeCourse: "Acute",
        PhysiologicalEffects: []string{
            "Vasodilation",
            "Pain sensation",
            "Immune cell recruitment",
            "Neurotransmission",
            "Exercise hyperemia",
        },
        BioenergeticRole: "Energy currency and extracellular signaling molecule",
    },
)

// GetExerciseUpregulatedMitokines returns mitokines that increase with exercise
func GetExerciseUpregulatedMitokines() []Mitokine {
    mitokines := []Mitokine{}
    for _, m := range []Mitokine{PGC1a, TFAM, FGF21, GDF15, Lactate, ATP} {
        if m.ExerciseRegulation == "Up" {
            mitokines = append(mitokines, m)
        }
    }
    return mitokines
}

// CalculateMitokineResponse estimates the change in mitokine levels based on exercise parameters
func CalculateMitokineResponse(mitokine Mitokine, exerciseIntensity float64, 
                             exerciseDuration float64, trainingStatus string) float64 {
    // Base multiplier (1.0 = no change)
    responseMultiplier := 1.0
    
    // Normalized exercise parameters
    intensityFactor := exerciseIntensity / 100.0
    durationFactor := exerciseDuration / 60.0 // Normalized to hours
    
    // Training status effect (untrained individuals often have more pronounced acute responses)
    trainingFactor := 1.0
    if trainingStatus == "Untrained" {
        trainingFactor = 1.3
    } else if trainingStatus == "Highly trained" {
        trainingFactor = 0.8
    }
    
    // Different mitokines have distinct response patterns
    switch mitokine.Name {
    case "Peroxisome Proliferator-Activated Receptor Gamma Coactivator 1-Alpha":
        // PGC-1α shows intensity-dependent response
        if exerciseIntensity > 70 {
            responseMultiplier += 0.8 * intensityFactor * trainingFactor
        } else {
            responseMultiplier += 0.5 * intensityFactor * durationFactor * trainingFactor
        }
        
    case "Mitochondrial Transcription Factor A":
        // TFAM has delayed response, more pronounced with chronic training
        responseMultiplier += 0.3 * intensityFactor * trainingFactor
        
    case "Fibroblast Growth Factor 21":
        // FGF21 responds to metabolic stress and duration
        responseMultiplier += 0.4 * durationFactor * trainingFactor
        if exerciseIntensity > 80 {
            responseMultiplier += 0.3 // High-intensity effect
        }
        
    case "Growth Differentiation Factor 15":
        // GDF15 responds strongly to novel or intense exercise stress
        noveltyFactor := 1.0
        if trainingStatus == "Untrained" {
            noveltyFactor = 1.5
        }
        responseMultiplier += 0.6 * intensityFactor * noveltyFactor
        
    case "Lactate":
        // Lactate rises exponentially above lactate threshold (~60-70% intensity)
        if exerciseIntensity < 60 {
            responseMultiplier += 0.3 * intensityFactor
        } else {
            // Exponential increase above threshold
            responseMultiplier += 0.3 * intensityFactor + 
                                 0.7 * math.Pow((intensityFactor-0.6)/0.4, 2)
        }
        
    case "Mitochondrial Reactive Oxygen Species":
        // mtROS shows U-shaped response curve
        if exerciseIntensity < 50 {
            responseMultiplier += 0.2 * intensityFactor
        } else if exerciseIntensity > 80 {
            responseMultiplier += 0.7 * intensityFactor
        } else {
            responseMultiplier += 0.3 * intensityFactor
        }
        
    case "Extracellular ATP":
        // ATP release is primarily intensity-dependent
        responseMultiplier += 0.5 * intensityFactor * intensityFactor * trainingFactor
    }
    
    return responseMultiplier
}

// PredictMitokineSignaling estimates the downstream effects of mitokine signaling
func PredictMitokineSignaling(mitokine Mitokine, responseLevel float64) map[string]float64 {
    // Initialize map of pathway activations
    pathwayActivation := make(map[string]float64)
    
    // Base activation is proportional to response level
    baseActivation := (responseLevel - 1.0) * 10.0 // Scale to 0-10 for typical exercise responses
    if baseActivation < 0 {
        baseActivation = 0
    }
    
    // Set pathway-specific activation based on the mitokine
    switch mitokine.Name {
    case "Peroxisome Proliferator-Activated Receptor Gamma Coactivator 1-Alpha":
        pathwayActivation["Mitochondrial biogenesis"] = baseActivation * 1.0
        pathwayActivation["Oxidative metabolism"] = baseActivation * 0.9
        pathwayActivation["Antioxidant defense"] = baseActivation * 0.7
        pathwayActivation["Angiogenesis"] = baseActivation * 0.5
        
    case "Mitochondrial Transcription Factor A":
        pathwayActivation["mtDNA replication"] = baseActivation * 0.9
        pathwayActivation["mtDNA transcription"] = baseActivation * 1.0
        pathwayActivation["Mitochondrial biogenesis"] = baseActivation * 0.8
        
    case "Fibroblast Growth Factor 21":
        pathwayActivation["Glucose uptake"] = baseActivation * 0.8
        pathwayActivation["Fatty acid oxidation"] = baseActivation * 0.7
        pathwayActivation["Energy expenditure"] = baseActivation * 0.6
        pathwayActivation["Mitochondrial function"] = baseActivation * 0.5
        
    case "Growth Differentiation Factor 15":
        pathwayActivation["Integrated stress response"] = baseActivation * 0.9
        pathwayActivation["Anti-inflammatory signaling"] = baseActivation * 0.6
        pathwayActivation["Appetite regulation"] = baseActivation * 0.5
        
    case "Lactate":
        pathwayActivation["Gluconeogenesis"] = baseActivation * 0.8
        pathwayActivation["GPR81 signaling"] = baseActivation * 0.7
        pathwayActivation["BDNF expression"] = baseActivation * 0.5
        pathwayActivation["Angiogenesis"] = baseActivation * 0.4
        
    case "Mitochondrial Reactive Oxygen Species":
        pathwayActivation["Nrf2 activation"] = baseActivation * 0.9
        pathwayActivation["Antioxidant enzyme induction"] = baseActivation * 0.8
        pathwayActivation["MAPK signaling"] = baseActivation * 0.6
        pathwayActivation["Inflammation"] = baseActivation * 0.5
        
    case "Extracellular ATP":
        pathwayActivation["Purinergic signaling"] = baseActivation * 1.0
        pathwayActivation["Vasodilation"] = baseActivation * 0.8
        pathwayActivation["Inflammation"] = baseActivation * 0.6
    }
    
    return pathwayActivation
}