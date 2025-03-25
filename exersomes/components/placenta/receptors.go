package placenta

// Placentokine Receptors represents a signaling molecule presented on placental tissue as a general cellular receptor across tissues
type Receptor struct {
    Name               string
    GeneID             string
    Type               string   // e.g., "G-protein coupled", "Ion channel", "Tyrosine kinase"
    Tissues            []string // Tissues where expressed
    CellularLocation   string   // e.g., "Membrane", "Cytosol", "Nucleus"
    Ligands            []string // Natural and synthetic binding molecules
    SignalingPathways  []string
    StructuralFeatures []string // e.g., "7-transmembrane", "Beta-barrel"
    ExpressionLevel    struct {
        Basal      string // "High", "Medium", "Low"
        Variability string // "Stable", "Dynamic", "Circadian"
    }
    FunctionalDomains []string
    PhysiologicalRole string // Primary function in normal physiology
    ExerciseResponse  struct {
        Expression   string // "Up", "Down", "No change", "Biphasic"
        Sensitivity  string // "Increased", "Decreased", "No change"
        TimeToEffect string // "Acute", "Chronic", "Both"
        Mechanism    string // Brief description of regulatory mechanism
    }
    DiseaseAssociations []string // Related pathologies
    TherapeuticTarget   bool     // Whether it's a known drug target
}

// Additional key exercise-responsive receptors
var (
    // ...existing receptors...
    
    GLUT1 = Receptor{
        Name:            "Glucose Transporter 1",
        Family:          "SLC2 Transporter",
        Type:            "Transporter",
        LigandBindings:  []string{"Glucose"},
        ExpressionSites: []string{"Red blood cells", "Brain", "Blood-brain barrier", "Endothelial cells"},
        DownstreamPathways: []string{"Basal glucose uptake"},
        ExerciseAdaptation: struct {
            AcuteEffect      string
            ChronicEffect    string
            AdaptationTime   float64
            MagnitudeFactor  float64
        }{
            AcuteEffect:      "No change",
            ChronicEffect:    "Mild upregulation",
            AdaptationTime:   6.0,
            MagnitudeFactor:  1.2,
        },
        PhysiologicalRoles: []string{"Basal glucose uptake", "Glucose transport across barriers", "Hypoxic metabolism"},
        Antagonists:        []string{"Cytochalasin B", "Phloretin"},
        Agonists:           []string{},
    },
    
    GLUT4 = Receptor{
        Name:            "Glucose Transporter 4",
        Family:          "SLC2 Transporter",
        Type:            "Transporter",
        LigandBindings:  []string{"Glucose"},
        ExpressionSites: []string{"Skeletal muscle", "Adipose tissue", "Heart"},
        DownstreamPathways: []string{"Insulin-stimulated glucose uptake", "Contraction-stimulated glucose uptake"},
        ExerciseAdaptation: struct {
            AcuteEffect      string
            ChronicEffect    string
            AdaptationTime   float64
            MagnitudeFactor  float64
        }{
            AcuteEffect:      "Translocation",
            ChronicEffect:    "Upregulation",
            AdaptationTime:   2.0,
            MagnitudeFactor:  1.8,
        },
        PhysiologicalRoles: []string{"Insulin-stimulated glucose disposal", "Exercise-induced glucose uptake", "Glycogen synthesis"},
        Antagonists:        []string{"Indinavir", "Cytochalasin B"},
        Agonists:           []string{"Exercise", "AICAR", "Metformin (indirect)"},
    },
    
    GLUT2 = Receptor{
        Name:            "Glucose Transporter 2",
        Family:          "SLC2 Transporter",
        Type:            "Transporter",
        LigandBindings:  []string{"Glucose", "Fructose"},
        ExpressionSites: []string{"Liver", "Pancreatic beta cells", "Intestine", "Kidney"},
        DownstreamPathways: []string{"Glucose sensing", "Bidirectional glucose transport"},
        ExerciseAdaptation: struct {
            AcuteEffect      string
            ChronicEffect    string
            AdaptationTime   float64
            MagnitudeFactor  float64
        }{
            AcuteEffect:      "No change",
            ChronicEffect:    "Moderate upregulation",
            AdaptationTime:   4.0,
            MagnitudeFactor:  1.3,
        },
        PhysiologicalRoles: []string{"Glucose sensing", "Hepatic glucose uptake/release", "Intestinal glucose absorption"},
        Antagonists:        []string{"Phlorizin (weak)"},
        Agonists:           []string{},
    },
    
    B3AR = Receptor{
        Name:            "Beta-3 Adrenergic Receptor",
        Family:          "G Protein-Coupled Receptor",
        Type:            "GPCR",
        LigandBindings:  []string{"Norepinephrine", "Epinephrine"},
        ExpressionSites: []string{"Adipose tissue (brown and white)", "Bladder", "Skeletal muscle"},
        DownstreamPathways: []string{"cAMP-PKA", "p38 MAPK", "AMPK", "PGC-1α"},
        ExerciseAdaptation: struct {
            AcuteEffect      string
            ChronicEffect    string
            AdaptationTime   float64
            MagnitudeFactor  float64
        }{
            AcuteEffect:      "Activation",
            ChronicEffect:    "Upregulation",
            AdaptationTime:   3.0,
            MagnitudeFactor:  1.5,
        },
        PhysiologicalRoles: []string{"Thermogenesis", "Lipolysis", "Energy expenditure", "Browning of white adipose tissue"},
        Antagonists:        []string{"SR59230A", "L-748,337"},
        Agonists:           []string{"Mirabegron", "CL-316,243"},
    },
    
    APJ = Receptor{
        Name:            "Apelin Receptor (APJ)",
        Family:          "G Protein-Coupled Receptor",
        Type:            "GPCR",
        LigandBindings:  []string{"Apelin", "Elabela/Toddler"},
        ExpressionSites: []string{"Heart", "Vasculature", "Adipose tissue", "Skeletal muscle", "CNS"},
        DownstreamPathways: []string{"Gi/Go", "PI3K-Akt", "AMPK", "ERK1/2", "eNOS"},
        ExerciseAdaptation: struct {
            AcuteEffect      string
            ChronicEffect    string
            AdaptationTime   float64
            MagnitudeFactor  float64
        }{
            AcuteEffect:      "Activation",
            ChronicEffect:    "Upregulation",
            AdaptationTime:   4.0,
            MagnitudeFactor:  1.4,
        },
        PhysiologicalRoles: []string{"Vasodilation", "Cardiac contractility", "Fluid homeostasis", "Glucose metabolism"},
        Antagonists:        []string{"ML221", "F13A"},
        Agonists:           []string{"MM07", "CMF-019"},
    },
    
    FATP1 = Receptor{
        Name:            "Fatty Acid Transport Protein 1",
        Family:          "SLC27 Transporter",
        Type:            "Transporter",
        LigandBindings:  []string{"Long-chain fatty acids"},
        ExpressionSites: []string{"Adipose tissue", "Heart", "Skeletal muscle"},
        DownstreamPathways: []string{"Fatty acid uptake", "Acyl-CoA synthetase"},
        ExerciseAdaptation: struct {
            AcuteEffect      string
            ChronicEffect    string
            AdaptationTime   float64
            MagnitudeFactor  float64
        }{
            AcuteEffect:      "Translocation",
            ChronicEffect:    "Upregulation",
            AdaptationTime:   3.0,
            MagnitudeFactor:  1.6,
        },
        PhysiologicalRoles: []string{"Fatty acid uptake", "Lipid metabolism", "Energy production"},
        Antagonists:        []string{"Lipofermata"},
        Agonists:           []string{},
    },
    
    CD36 = Receptor{
        Name:            "CD36/Fatty Acid Translocase",
        Family:          "Scavenger Receptor",
        Type:            "Transporter/Receptor",
        LigandBindings:  []string{"Long-chain fatty acids", "Oxidized LDL", "Thrombospondin"},
        ExpressionSites: []string{"Skeletal muscle", "Heart", "Adipose tissue", "Macrophages"},
        DownstreamPathways: []string{"Fatty acid uptake", "Src family kinases", "MAPK", "NFκB"},
        ExerciseAdaptation: struct {
            AcuteEffect      string
            ChronicEffect    string
            AdaptationTime   float64
            MagnitudeFactor  float64
        }{
            AcuteEffect:      "Translocation",
            ChronicEffect:    "Upregulation",
            AdaptationTime:   2.5,
            MagnitudeFactor:  1.7,
        },
        PhysiologicalRoles: []string{"Fatty acid uptake", "Lipid metabolism", "Muscle fuel selection", "Inflammation"},
        Antagonists:        []string{"SSO", "AP5055"},
        Agonists:           []string{},
    },
)