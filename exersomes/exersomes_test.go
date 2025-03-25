package main

import (
	"os"
	"testing"
)

// Test XML sanitization
func TestSanitizeXML(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"<root>test &usehistory</root>", "<root>test &amp;usehistory</root>"},
		{"<root>test & value</root>", "<root>test &amp; value</root>"},
		{"<root>normal xml</root>", "<root>normal xml</root>"},
	}

	for _, test := range tests {
		result := string(sanitizeXML([]byte(test.input)))
		if result != test.expected {
			t.Errorf("sanitizeXML(%s) = %s; expected %s", test.input, result, test.expected)
		}
	}
}

// Test loading input list
func TestLoadInputList(t *testing.T) {
	// Create a temporary file
	filename := "test_genes.txt"
	testContent := "GENE1\nGENE2\n\nGENE3  \n"

	err := os.WriteFile(filename, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(filename)

	// Test loading
	genes := loadInputList(filename)

	// Verify results
	if len(genes) != 3 {
		t.Errorf("Expected 3 genes, got %d", len(genes))
	}

	expectedGenes := []string{"GENE1", "GENE2", "GENE3"}
	for i, gene := range expectedGenes {
		if genes[i] != gene {
			t.Errorf("Expected gene %s at position %d, got %s", gene, i, genes[i])
		}
	}
}

// Test progress tracker
func TestProgressTracker(t *testing.T) {
	progress := NewProgressTracker(10)

	if progress.total != 10 {
		t.Errorf("Expected total 10, got %d", progress.total)
	}

	if progress.completed != 0 {
		t.Errorf("Expected completed 0, got %d", progress.completed)
	}

	progress.Increment()

	if progress.completed != 1 {
		t.Errorf("Expected completed 1, got %d", progress.completed)
	}
}

// Mock test for process genes in parallel
func TestProcessGenesInParallel(t *testing.T) {
	genes := []string{"GENE1", "GENE2", "GENE3"}
	processed := make(map[string]bool)

	processGenesInParallel(genes, 2, func(genes []string, results chan string) {
		for _, gene := range genes {
			processed[gene] = true
			results <- gene
		}
	})

	for _, gene := range genes {
		if !processed[gene] {
			t.Errorf("Gene %s was not processed", gene)
		}
	}
}
