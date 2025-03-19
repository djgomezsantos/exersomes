import requests
from Bio import Entrez, SeqIO
import pandas as pd
import time
import sys

# Set your email for NCBI E-utilities
Entrez.email = "your_email@example.com"

def fetch_exerkine_genes(gene_names):
    gene_ids = []
    for gene in gene_names:
        try:
            handle = Entrez.esearch(db="gene", term=f"{gene} exerkines", retmax=10)
            record = Entrez.read(handle)
            gene_ids.extend(record["IdList"])
            handle.close()
            time.sleep(1)  # Delay to avoid hitting the rate limit
        except Exception as e:
            print(f"Error fetching gene ID for {gene}: {e}")
    return gene_ids

def fetch_gene_details(gene_ids):
    gene_details = []
    for gene_id in gene_ids:
        try:
            handle = Entrez.efetch(db="gene", id=gene_id, rettype="xml", retmode="text")
            gene_record = handle.read()
            gene_details.append(gene_record)  # You can parse this XML response for details
            handle.close()
            time.sleep(1)  # Delay to avoid hitting the rate limit
        except Exception as e:
            print(f"Error fetching details for Gene ID {gene_id}: {e}")
    return gene_details

def save_results(gene_ids, file_name):
    df = pd.DataFrame({'Gene_ID': gene_ids})
    df.to_csv(file_name, index=False)

def main(gene_names):
    # Fetch IDs for specified genes
    genes_found = fetch_exerkine_genes(gene_names)
    print("Fetched Gene IDs:", genes_found)

    # Fetch details for the found gene IDs
    gene_details = fetch_gene_details(genes_found)
    
    # Save the results to a CSV file
    save_results(genes_found, 'exerkine_genes.csv')

if __name__ == "__main__":
    # You can pass gene names as command-line arguments
    if len(sys.argv) > 1:
        gene_names = sys.argv[1:]  # Accepts gene names as command-line arguments
    else:
        # Default gene names related to exerkines
        gene_names = [
            "IL6",    # Interleukin 6
            "FNDC5",  # Fibronectin type III domain containing 5
            "BDKRB2", # Bradykinin receptor B2
            # Add more genes of interest...
        ]
    
    main(gene_names)