# Exersomes

**Exersomes** is a comprehensive Go-based bioinformatics and simulation framework designed to model, analyze, and predict the molecular responses to exercise. 

The project focuses on **exerkines**—signaling molecules (proteins, metabolites, miRNAs, RNAs) released by various tissues (like skeletal muscle, mitochondria, and placenta) during physical exertion. By integrating real-world biological data with programmatic simulations and machine learning, Exersomes aims to map the complex molecular networks of exercise physiology.

## 🧬 Key Features

- **Molecular Type Modeling**: Strongly-typed Go structs representing diverse exercise-responsive molecules including `Proteins` (e.g., IL-6, Irisin), `Metabolites` (e.g., Lactate), `miRNAs`, `RNAs`, and `Vesicles`.
- **Temporal Dynamics & Circulating Factors**: Calculates physiological half-lives, acute vs. chronic release times, and tracks bloodstream factors (e.g., Epinephrine, Cortisol, Lactate, BDNF) during and post-exercise.
- **Mitochondrial Signaling (Mitokines)**: Simulates the release and downstream effects of mitokines (e.g., PGC-1α, TFAM, FGF21, GDF15, mtROS) in response to bioenergetic demands and mitochondrial stress.
- **Exercise Prescription Simulation**: Models specific workout protocols encompassing intensity and duration to predict quantitative molecular responses over time.
- **Automated NCBI Integration**: Uses NCBI E-utilities to dynamically fetch gene references, protein accessions, FASTA sequences, KEGG/BioSystems pathway maps, and PubMed functional insights for targeted exerkines.
- **Latent Space Modeling (Transformer)**: Includes architectures for encoding exersome cargos (myokines, ligands, receptors) into a latent space to analyze and reconstruct molecular signatures.

## 🛠 Prerequisites

- **Go** 1.24.1 or higher
- **NCBI E-utilities (Entrez Direct)**: Required for the Go-based data fetching pipeline (`esearch`, `efetch`, `elink`).
  - **macOS / Linux Installation**: Run the official setup script in your terminal:
    ```bash
    sh -c "$(curl -fsSL https://ftp.ncbi.nlm.nih.gov/entrez/entrezdirect/install-edirect.sh)"
    ```
  - *Note: After installation, follow the terminal prompt to add the `edirect` folder to your system's `PATH`. Your terminal needs to be able to find these commands globally for the Go application to work.*

## 🚀 Installation & Setup

1. Clone the repository and navigate to the project directory.
2. Run the project directly without building a permanent executable:
   ```bash
   go run ./exersomes/exersomes.go
   ```

## 💻 Usage

### 1. Retrieving RNA & Protein Sequences (NCBI Pipeline)

To automatically fetch the FASTA sequences and biological data for your exerkines, ligands, and receptors:

1. **Edit the Input File**: Open `data/raw_data/exerkines.txt` and ensure your target genes are listed (one per line). You can include exerkines (e.g., `IL6`), ligands (e.g., `IGF1`), and their corresponding receptors (e.g., `IL6R`, `IGF1R`).
2. **Run the Pipeline**: Execute the main Go program from your terminal:
   ```bash
   go run ./exersomes/exersomes.go
   ```
3. **Collect Outputs**: The program uses parallel workers to query NCBI databases and will generate several files in your root directory:
   - `rna_sequences.fasta`: Ready-to-use human mRNA nucleotide sequences.
   - `protein_sequences.fasta`: Amino acid sequences for structural modeling.
   - `protein_info.tsv`: Protein accessions and molecular weights.
   - `gene_references.tsv`: Gene IDs, symbols, and chromosome mapping.
   - `pathway_maps.tsv` & `functional_insights.tsv`: Mined functional insights from PubMed and GO.

### 2. Physiological Simulations

Programmatically integrate the cardiovascular and placental components to predict molecular responses:

- **Bloodstream Factors**: Use `PredictCirculatingProfileDuringExercise(exerciseType, intensityPercent, durationMinutes, timePointsCount)` to map the rise and decay of hormones like Epinephrine.
- **Mitokine Signalling**: Use `CalculateMitokineResponse` to estimate the change in mitokine levels based on exercise parameters and training status.

## 🗂 Project Structure

```text
exersomes/
├── exersomes.go                             # Main NCBI data fetching entrypoint
├── molecular_types/                         # Core exerkine models (miRNA, RNA, Proteins)
├── components/
│   ├── muscle/                              # Skeletal muscle simulation
│   ├── cardiovascular/bloodstream/          # Circulating blood factors & temporal decay
│   └── placenta/                            # Mitokines and mitochondrial signaling
├── models/                                  # ML models (Encoder)
└── src/                                     # APIs, data preprocessing, and network analysis
```

## 📄 License

Open-source. Please ensure compliance with NCBI's data usage guidelines when running the automated E-utilities fetching pipeline.
