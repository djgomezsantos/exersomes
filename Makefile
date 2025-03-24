.PHONY: all build test run analyze visualize clean

all: build run analyze visualize

build:
    @echo "Building exersomes..."
    cd exersomes && go build

test:
    @echo "Running tests..."
    cd exersomes && go test -v

run:
    @echo "Running gene data retrieval..."
    ./exersomes/exersomes

analyze:
    @echo "Running analysis..."
    python3 src/data_analysis.py

visualize:
    @echo "Generating visualizations..."
    python3 src/visualize.py

clean:
    rm -f gene_references.tsv protein_info.tsv protein_sequences.fasta pathway_maps.tsv functional_insights.tsv
    rm -f *.png