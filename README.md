# Exersomes

A computational tool for analyzing exerkines (signaling molecules released from exercise).

## Overview

Exersomes combines Go-based data retrieval with Python analysis to study exercise-induced signaling molecules. The project provides high-throughput processing of gene/protein sequences and facilitates pathway analysis of exerkines, their ligands, and receptors. Using bioinformatics, this project aims to facilitate the exploration and generation of biologically relevant sequences that could aid in understanding gene expression during physical activity.

## Features
- Fast concurrent retrieval of gene/protein data from NCBI databases
- High-throughput processing and BLAST similarity searches
- Network analysis and pathway visualization
- Gene and protein sequence generation using transformer models
- Support for variable-length RNA sequence analysis
- Fetches gene data and annotations from NCBI and Ensembl databases
- Performs high-throughput processing of gene sequences and utilizes BLAST for similarity searches
- Generates gene and protein sequences using a DGM transformer model
- Supports padding and attention masking for variable-length RNA sequences


## Requirements

### Go Components
- Go 1.16+
- NCBI Entrez Direct utilities

### Python Components
- Python 3.8+
- Python packages: pandas, matplotlib, seaborn, networkx

## Installation

1. Install Go dependencies:

```
cd exersomes go mod tidy
```

2. Install Python dependencies:

```
pip install -r requirements.txt
```

3. Install NCBI Entrez Direct utilities:

```
sh -c "$(curl -fsSL https://ftp.ncbi.nlm.nih.gov/entrez/entrezdirect/install-edirect.sh)"
```

## Usage

1. Run the complete workflow:

```
make all
```

2. Run individual steps:

```make build``` # Build Go code make run # Fetch data from NCBI 
```make analyze ``` # Run Python analysis make test # Run unit tests



## Output Files

- `gene_references.tsv`: Basic gene information
- `protein_info.tsv`: Protein details and properties
- `protein_sequences.fasta`: Protein sequences in FASTA format
- `pathway_maps.tsv`: Gene pathway associations
- `functional_insights.tsv`: Functional annotations from literature and GO

### Components:

- **Go Module**: Fast concurrent retrieval of gene/protein data from NCBI
- **Python Analysis**: Processing and machine learning on the retrieved data
- **Visualization**: Network analysis and pathway visualization

## Project Structure

. 
├── LICENSE 
├── README.md 
├── data 
│ ├── processed_data 
│ │ ├── exerkines_list.txt 
│ │ └── gene_references.tsv 
│ └── raw_data │ 
├── exersome_gene_ids.txt 
│ ├── gene_ids.txt 
│ ├── gene_ids_with_names.txt 
│ └── target_gene_list.txt 
├── exersomes │ ├── components │ │ ├── RNA │ │ ├── cf-exerkines │ │ ├── miRNA │ │ ├── protein │ │ └── vesicles │ ├── exersomes.go │ └── go.mod ├── models │ ├── model.py │ └── train_model.py ├── scripts │ ├── BLAST_results.py │ ├── fetch_gene_info.py │ └── retrieve_DNA.py ├── setup.py ├── src │ ├── init.py │ ├── api.py │ ├── main.py │ ├── model.py │ └── preprocessing.py └── tests


## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Citation

If you use Exersomes in your research, please cite:

Gomez, DJ. et al. (2025). Exersomes: A computational tool for analyzing exercise-induced signaling molecules. [Software]. Available from https://github.com/djgomez8/Exersomes