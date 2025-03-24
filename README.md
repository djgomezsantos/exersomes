# Exersomes
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
