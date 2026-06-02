# Exersomes

**Exersomes** is a comprehensive Go-based bioinformatics and simulation framework designed to model, analyze, and predict the molecular responses to exercise. 

The project focuses on **exerkines**—signaling molecules (proteins, metabolites, miRNAs, RNAs) released by various tissues (like skeletal muscle, mitochondria, and placenta) during physical exertion. By integrating real-world biological data with programmatic simulations and a powerful **Protein-Nucleic acid Language Model (PNLM)**, Exersomes aims to map the complex molecular networks of exercise physiology for personalized precision medicine, gene therapy, and disease modeling.

## 🧬 Key Features

- **Molecular Type Modeling**: Strongly-typed Go structs representing diverse exercise-responsive molecules including `Proteins` (e.g., IL-6, Irisin), `Metabolites` (e.g., Lactate), `miRNAs`, `RNAs`, and `Vesicles`.
- **Protein-Nucleic Language Model (PNLM)**: Natively integrates with transformer-based architectures to measure and embed **RNA trajectory velocities**, **protein expression levels**, and lipid/metabolite cross-talk for high-precision exerkine generation.
- **Temporal Dynamics & Circulating Factors**: Calculates physiological half-lives, acute vs. chronic release times, and tracks bloodstream factors (e.g., Epinephrine, Cortisol, Lactate, BDNF) during and post-exercise.
- **Mitochondrial Signaling (Mitokines)**: Simulates the release and downstream effects of mitokines (e.g., PGC-1α, TFAM, FGF21, GDF15, mtROS) in response to bioenergetic demands and mitochondrial stress.
- **Exercise Prescription Simulation**: Models specific workout protocols encompassing intensity and duration to predict quantitative molecular responses over time.
- **Automated NCBI Integration**: Uses NCBI E-utilities to dynamically fetch gene references, protein accessions, FASTA sequences, KEGG/BioSystems pathway maps, and PubMed functional insights for targeted exerkines.
- **Translational Consortium Interoperability**: Capable of integrating structural and spatial molecular responses for mapping to **MoTrPAC**, the **HuBMAP** Human Reference Atlas, and **HTAN** exercise immuno-oncology models.

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
   - `rna_sequences.fasta`: Ready-to-use human and rat mRNA nucleotide sequences.
   - `protein_sequences.fasta`: Amino acid sequences for structural modeling.
   - `protein_info.tsv`: Protein accessions and molecular weights.
   - `rna_references.tsv`: RNA RefSeq accessions and descriptions for human and rat.
   - `pathway_maps.tsv` & `functional_insights.tsv`: Mined functional insights from PubMed and GO.

### 2. Physiological Simulations

Programmatically integrate the additional omes, cardiometabolism, placental components, and organokiome to infer and explain molecular responses:

- **Adipomyokinome**: Use `InferAdipomyokines(exerciseType, intensityPercent, durationMinutes, timePointsCount)` to map the rise and decay of adipomyokines like IL-6 in obesity and cancer.
- **Mitokine Signalling**: Use `CalculateMitokineResponse` to estimate the change in mitokine levels based on exercise parameters and training status for healthy pregnancy and pregnancy complications (gestational diabetes and preclampsia).
- **Organokinome**: Use `CalculateOrganokineResponse` to estimate the change in organokine levels based on exercise parameters, gut microbiota, and training status.

### 3. Language Model (LM) Integrations

Exersomes provides a unified Go client interface to orchestrate state-of-the-art machine learning models for deep molecular analysis:

- **PNMLM (Protein-Nucleic Acid-Metabolite Language Model)**: Generate multi-modal latent embeddings combining RNA velocity, protein expression, and localized lipids. Predict systemic exersome assembly and decode personalized exercise prescriptions.
- **ESM-2 (Evolutionary Scale Modeling)**: Perform zero-shot structural prediction, inter-residue contact mapping, and generate novel viable protein sequences (exerkines/ligands).
- **EDM (Equivariant Diffusion Model)**: Generate highly precise 3D biochemical structures (metabolites, lipids) natively adhering to organic chemoinformatics constraints (C, H, O, N, S).


## 🗂 Project Structure

```text
exersomes/
├── exersomes.go                             # Main NCBI data fetching entrypoint
├── molecular_types/                         # Core exerkine models (miRNA, RNA, Proteins)
├── components/
│   ├── muscle/                              # Skeletal muscle simulation
│   ├── cardiovascular/bloodstream/          # Circulating blood factors & temporal decay
│   └── placenta/                            # Mitokines and mitochondrial signaling
├── models/                                  # Deep learning architecture and clients
│   ├── lm.go                                # PNMLM, ESM-2, and EDM language model clients
│   ├── latentspace.go                       # VAE latent space representations
│   └── encoder.go                           # Exersome multi-omic cargo encoder
└── src/                                     # APIs, data preprocessing, and network analysis
```

## 📄 License

Open-source. Please ensure compliance with NCBI's data usage guidelines when running the automated E-utilities fetching pipeline.

If you use Exersomes in your research, please cite:

Gomez, DJ. et al. (2025). Exersomes: A computational tool for analyzing exercise-induced signaling molecules. [Software]. Available from https://github.com/gomezdj/Exersomes
