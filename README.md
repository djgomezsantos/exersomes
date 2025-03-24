# Exersomes
Exersomes is a nucleotide sequence generation tool focused on generating RNA sequences related to exercise-induced exerkines, ligands, and receptors. Using bioinformatics, this project aims to facilitate the exploration and generation of biologically relevant RNA sequences that could aid in understanding gene expression during physical activity. Exersomes generates exerkines, ligands, and receptors. ExerkineNet is tailored to include the essential details concerning your RNA sequence generation work using a transformer model or VAE:

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
- Generates RNA sequences using GPT-4 via OpenAI API.
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
```

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/gomezdj/exerkinenet.git
   ```
2. Navigate to the project directory:
   ```bash
   cd exerkinenet/python
   ```

## Data Structure
The project organizes data into two main directories under `data/`:

- **raw_data/**: Contains unprocessed nucleotide sequences and gene annotations.
- **processed_data/**: Contains normalized RNA sequences, attention masks, padded sequences, and features for model training.

## Usage
1. Fetch gene information and annotations:
   - Run `fetch_gene_info.py` to retrieve sequences from NCBI.
   ```bash
   python fetch_gene_info.py
   ```

2. Generate RNA sequences:
   - Use the `generate_sequence()` function to utilize GPT-4 for generating RNA sequences.
   ```python
   generated_sequence = generate_sequence("Your input RNA sequence or prompt here")
   ```

3. Adjust parameters in the function as needed for model performance.

## Contributing
Contributions are welcome! Please feel free to submit a pull request or open an issue for any suggestions or improvements.

1. Fork the repository.
2. Create a new feature branch (`git checkout -b feature/new-feature`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/new-feature`).
5. Open a pull request.

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

### Notes:
- **Project Name**: Make sure to customize the project name and GitHub repository link to match your actual project.
- **Requirements**: Update the list of required libraries if you have any additional dependencies.
- **Usage Section**: Be sure to include relevant instructions or examples for your methods, focusing on how to effectively utilize the GPT-4 API for sequence generation.
- **Contribution Guidelines**: Feel free to alter this section according to your project's contribution policies.

This README provides a comprehensive overview of your project and should serve well to inform users about the purpose, setup, and usage of your RNA sequence generation tool.
