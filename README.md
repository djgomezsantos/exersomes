# Exersomes

A computational tool for analyzing exerkines (signaling molecules released during exercise). This project combines Go-based data retrieval with Python analysis to study exerkines (signaling molecules released during exercise).

## Requirements

- Go 1.16+
- Python 3.8+
- NCBI Entrez Direct utilities
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

Exersomes is a nucleotide and peptide/protein sequence generation tool focused on generating gene and protein sequences related to exercise-induced exerkines, ligands, and receptors. Using bioinformatics, this project aims to facilitate the exploration and generation of biologically relevant sequences that could aid in understanding gene expression during physical activity.

## Table of Contents
- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Data Structure](#data-structure)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Features
- Fetches gene data and annotations from NCBI and Ensembl databases.
- Performs high-throughput processing of gene sequences and utilizes BLAST for similarity searches.
- Generates gene and protein sequences using a DGM transformer model.
- Supports padding and attention masking for variable-length RNA sequences.

### Getting Started:

1. Run the Go data retrieval:


## Requirements
- Python 3.7 or higher
- Libraries: 
  - `openai`
  - `transformers`
  - `torch`
  - `numpy`
  - `pandas`
  - `biopython`
  - `requests`

You can install the required libraries using pip:

```bash
pip install openai transformers torch numpy pandas biopython requests

2. Run the Python analysis:

python -m src.main

