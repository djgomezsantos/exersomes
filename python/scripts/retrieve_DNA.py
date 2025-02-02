from Bio import Entrez, SeqIO

# Set your email
Entrez.email = "your_email@example.com"

def fetch_dna_sequence(gene_id):
    handle = Entrez.efetch(db="nucleotide", id=gene_id, rettype="gb", retmode="text")
    record = SeqIO.read(handle, "genbank")
    handle.close()
    return str(record.seq)

# Example usage: Fetch sequence for a specific gene ID
gene_id = "NM_001301717"  # Replace with your specific gene ID
dna_sequence = fetch_dna_sequence(gene_id)
print(dna_sequence)