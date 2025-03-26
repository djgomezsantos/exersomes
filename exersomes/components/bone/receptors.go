package receptors

// Receptor represents a protein that binds to a ligand
type Receptor struct {
	Name               string
	Ligand             string
	SignalingPathway   string
	BiologicalFunction string
}

// Example receptors
var (
	VEGFR = Receptor{
		Name:               "Vascular Endothelial Growth Factor Receptor",
		Ligand:             "VEGF",
		SignalingPathway:   "PI3K-Akt",
		BiologicalFunction: "Mediates angiogenesis",
	}

	TNFRSF11B = Receptor{
		Name:               "TNF Receptor Superfamily Member 11B",
		Ligand:             "RANKL",
		SignalingPathway:   "RANK/RANKL/OPG",
		BiologicalFunction: "Inhibits osteoclastogenesis",
	}

	BMP2R = Receptor{
		Name:               "Bone Morphogenetic Protein Receptor Type 2",
		Ligand:             "BMP2, BMP4",
		SignalingPathway:   "SMAD",
		BiologicalFunction: "Regulates bone development and repair",
	}

	// Add more receptors
)
