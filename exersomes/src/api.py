from Bio import Entrez

Entrez.email = "your_email@example.com"
handle = Entrez.esearch(db="gene", term="IL6 exerkines")
record = Entrez.read(handle)
handle.close()
gene_ids = record["IdList"]

import requests

# Example: Get information about a gene by ID
response = requests.get("https://rest.ensembl.org/lookup/id/ENSG00000139618?content-type=application/json")
gene_info = response.json()
print(gene_info)

from transformers import GPT4LMHeadModel, GPT2Tokenizer

tokenizer = GPT2Tokenizer.from_pretrained('model_name')
model = GPT4LMHeadModel.from_pretrained('model_name')
inputs = tokenizer("Example input", return_tensors='pt')
outputs = model.generate(inputs['input_ids'])
