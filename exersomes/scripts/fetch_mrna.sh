#!/bin/bash

# Directory to store sequences
OUTDIR="data/sequences/mRNA"
mkdir -p "$OUTDIR"

# Gene list
GENE_LIST=("ADPN" "ANGPTL4" "ANGPT1" "APLN" "ANGPTL8" "ADIPOR1" "ADIPOR2" "BDNF" "LEP" "DARC" "APLNR" "NTRK2" "AHSG" "CX3CL1" "FGF21" "FGFR1" "FGFR2" "FNDC5" "FST" "GDF15" "GFRAL" "HSPA1A" "IL6" "IL6ST" "IL6R" "IL7" "IL7R" "CXCL8" "CXCR1" "IL13" "IL13RA1" "IL15" "IL15" "IL15" "IL15RA" "IL22" "IGFBP1" "FNDC5" "IGF1" "IGF1R" "METRNL" "LIF" "LIFR" "Ostn" "MSTN" "SPARC" "TGFB1" "TGFB2" "VEGFA" "KDR" "TNF" "TNFRSF1A" "TNF" "TNFRSF1B" "TNFRSF1B" "TNFRSF1A" "SDC4" "PPARGC1A" "PNPLA2" "PRDM16" "BMP8A" "BMP8B" "UCP1" "IL1RN" "CSTB" "GLPD1" "RBP4" "MSTN" "DCX" "S100A10" "CLU" "LEP" "LEPR" "DCN" "CCL2" "CCR2" "NAMPT" "RETN" "ITLN1" "LUM" "TMSB4X" "CNTF" "CNTFR" "INS" "INSR" "VIM" "NRG4" "ERBB3" "PLTP" "IL1A" "IL1B" "IL1R1" "IL4" "IL4R" "IL18" "IL18R1" "GDF11" "ACVR1B" "ACVR2B" "SIRT1" "KL" "CITED4" "CXCL1" "CXCR1" "CSF3" "CTGF" "DPP4" "AGTR1" "GCG" "GRN" "REN" "MME" "CD34" "IL33" "IL1RL1" "FABP4" "ENG" "MCAM" "CD4" "CD8A" "CD14" "ITGA2B" "GP1BA" "SELP" "HLA-DRA" "KIT" "FGF19" "SERPINA12" "RARRES2" "CMKLR1" "NPFFR2" "GPR158" "GPRC6A" "ADIPOQ" "ADIPOR1" "ADIPOR2" "SPX" "FABP4" "IL10" "FNDC5" "ERFE" "TNFRSF11B" "BGLAP" "SPARC" "SOST" "BMP2" "BMP4" "SPP1" "FETUB" "ADIPOR2" "INHBE" "IFNG" "IFNGR1" "VDR" "NFKB1" "PSMD1" "AKT1" "EIF4EBP1" "MAPK3" "MAPK14" "RPS6KB1" "SLC2A4" "AGER" "TRIM63" "FBXO32" "CCL11" "FAS" "ICAM1" "CCL22" "MMP1" "MMP2" "MMP7" "MMP9" "SERPINE1" "CCL18" "CCL5" "TIMP3" "F13A1" "LOXL1" "NRP1" "NRP2" "PPARG" "HCAR1" "GPR81" "ITGA5" "ITGB1" "TLR4" "CAP1" "ITGAV" "ITGB3" "SUCNR1" "PPARGC1A")

# Function using NCBI EDirect if available
fetch_with_edirect() {
  local gene=$1
  esearch -db nucleotide -query "${gene}[Gene] AND Rattus norvegicus[Organism] AND mRNA" \
    | efetch -format fasta > "${OUTDIR}/${gene}_mRNA.fasta"
}

# Fallback using curl if EDirect is not available
fetch_with_curl() {
  local gene=$1
  uid=$(curl -s "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/esearch.fcgi?db=nucleotide&term=${gene}[Gene]+AND+Rattus+norvegicus[Organism]+AND+mRNA&retmode=json" | jq -r '.esearchresult.idlist[0]')
  if [[ -n "$uid" && "$uid" != "null" ]]; then
    curl -s "https://eutils.ncbi.nlm.nih.gov/entrez/eutils/efetch.fcgi?db=nucleotide&id=${uid}&rettype=fasta&retmode=text" > "${OUTDIR}/${gene}_mRNA.fasta"
  else
    echo "No result for $gene"
  fi
}

# Detect whether esearch exists
USE_EDIRECT=true
command -v esearch >/dev/null 2>&1 || USE_EDIRECT=false

# Main loop
for GENE in "${GENE_LIST[@]}"; do
  echo "Fetching mRNA for $GENE..."
  if $USE_EDIRECT; then
    fetch_with_edirect "$GENE"
  else
    fetch_with_curl "$GENE"
  fi
  sleep 0.4
done