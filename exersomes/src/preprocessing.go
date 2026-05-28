// src/preprocessing.go
package main

import (
	"encoding/csv"
	"exersomes/exersomes/molecular_types"
	"os"
	"strings"
)

func LoadExerkineData(path string) ([]molecular_types.Exerkine, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'         // For TSV files
	reader.FieldsPerRecord = -1 // Allow a variable number of fields per row

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var exerkines []molecular_types.Exerkine

	// Parse TSV into Exerkine structs (Targeting protein_info.tsv format)
	// Header: Query, Protein_ID, Accession, Name, Length, Molecular_Weight
	for i, record := range records {
		if i == 0 || len(record) < 4 {
			continue // Skip header row or malformed rows
		}

		exerkine := molecular_types.Exerkine{
			Name:           record[0], // Query / Gene Symbol
			Category:       "Protein",
			PDBID:          record[2], // Accession ID
			BiologicalFunc: record[3], // Descriptive protein name
			TissueSources:  []string{},
			Receptors:      []string{},
		}
		exerkines = append(exerkines, exerkine)
	}

	return exerkines, nil
}

// EnrichWithFunctionalInsights parses functional_insights.tsv to append text-mined data
func EnrichWithFunctionalInsights(exerkines []molecular_types.Exerkine, path string) ([]molecular_types.Exerkine, error) {
	file, err := os.Open(path)
	if err != nil {
		return exerkines, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		return exerkines, err
	}

	// Header: Gene, Function_Type, Description, Evidence, Reference_PMID
	for i, record := range records {
		if i == 0 || len(record) < 3 {
			continue
		}

		gene, funcType, description := record[0], record[1], record[2]

		// Map insights to the corresponding Exerkine by Name
		for j := range exerkines {
			if exerkines[j].Name == gene && !strings.Contains(exerkines[j].BiologicalFunc, description) {
				exerkines[j].BiologicalFunc += "; [" + funcType + "] " + description
			}
		}
	}

	return exerkines, nil
}

// SystemCatalog holds the dynamically routed molecules for each organ system
type SystemCatalog struct {
	MuscleExerkines       []molecular_types.Exerkine
	NeuralFactors         []molecular_types.Exerkine
	CardiovascularFactors []molecular_types.Exerkine
	SystemicReceptors     []molecular_types.Exerkine
}

// BuildSystemCatalog uses text-mining on the functional insights to route molecules to components
func BuildSystemCatalog(exerkines []molecular_types.Exerkine) SystemCatalog {
	catalog := SystemCatalog{}

	for _, ex := range exerkines {
		desc := strings.ToLower(ex.BiologicalFunc)
		name := strings.ToUpper(ex.Name)

		// 1. Identify Receptors (Ends with 'R' like IGF1R, or mentions receptor)
		if strings.HasSuffix(name, "R") || strings.HasSuffix(name, "R1") || strings.HasSuffix(name, "R2") || strings.Contains(desc, "receptor") {
			ex.Category = "Receptor"
			catalog.SystemicReceptors = append(catalog.SystemicReceptors, ex)
			continue // Keep receptors in their own pool for cellular surface binding
		}

		// 2. Route to Neural Component
		if strings.Contains(desc, "brain") || strings.Contains(desc, "neuron") || strings.Contains(desc, "synap") || strings.Contains(desc, "cognitive") {
			catalog.NeuralFactors = append(catalog.NeuralFactors, ex)
		}

		// 3. Route to Muscle Component
		if strings.Contains(desc, "muscle") || strings.Contains(desc, "myocyte") || strings.Contains(desc, "contract") || strings.Contains(desc, "hypertrophy") {
			catalog.MuscleExerkines = append(catalog.MuscleExerkines, ex)
		}

		// 4. Route to Cardiovascular/Bloodstream Component
		if strings.Contains(desc, "blood") || strings.Contains(desc, "cardiac") || strings.Contains(desc, "vascular") || strings.Contains(desc, "angiogen") {
			catalog.CardiovascularFactors = append(catalog.CardiovascularFactors, ex)
		}
	}

	return catalog
}

func NormalizeData(exerkines []molecular_types.Exerkine) []molecular_types.Exerkine {
	// Implement normalization logic
	return exerkines
}
