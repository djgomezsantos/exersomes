package ligands

// Ligand represents a signaling molecule that binds to a receptor
type Ligand struct {
	Name               string
	Receptor           string
	SignalingPathway   string
	BiologicalFunction string
}

// Exerkine ligands
var (
	VEGF = Ligand{
		Name:               "Vascular Endothelial Growth Factor",
		Receptor:           "VEGFR",
		SignalingPathway:   "PI3K-Akt",
		BiologicalFunction: "Promotes angiogenesis",
	}
	SPARC = Ligand{
		Name:               "Secreted Protein Acidic and Cysteine Rich",
		Receptor:           "Multiple ECM proteins",
		SignalingPathway:   "Integrin-mediated",
		BiologicalFunction: "Bone mineralization and collagen binding",
	}

	SOST = Ligand{
		Name:               "Sclerostin",
		Receptor:           "LRP5/6",
		SignalingPathway:   "Wnt/β-catenin (inhibitor)",
		BiologicalFunction: "Inhibits bone formation",
	}

	BMP2 = Ligand{
		Name:               "Bone Morphogenetic Protein 2",
		Receptor:           "BMP2R, BMPR1A",
		SignalingPathway:   "SMAD",
		BiologicalFunction: "Induces bone and cartilage formation",
	}

	BMP4 = Ligand{
		Name:               "Bone Morphogenetic Protein 4",
		Receptor:           "BMP2R, BMPR1A",
		SignalingPathway:   "SMAD",
		BiologicalFunction: "Regulates bone and cartilage development",
	}

	SPP1 = Ligand{
		Name:               "Secreted Phosphoprotein 1",
		Receptor:           "Integrins, CD44",
		SignalingPathway:   "Integrin-mediated",
		BiologicalFunction: "Bone remodeling and immune regulation",
	}

	// Add more receptors
)
