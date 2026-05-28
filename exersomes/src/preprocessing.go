// src/preprocessing.go
package main

import (
	"encoding/csv"
	"exersomes/exersomes/molecular_types"
	"os"
)

func LoadExerkineData(path string) ([]molecular_types.Exerkine, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t' // For TSV files

	// Parse CSV/TSV into Exerkine structs
	// Implementation omitted for brevity
	return nil, nil
}

func NormalizeData(exerkines []molecular_types.Exerkine) []molecular_types.Exerkine {
	// Implement normalization logic
	return exerkines
}
