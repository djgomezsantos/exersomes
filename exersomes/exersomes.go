package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

// NCBI Response struct definitions
type ESearchResult struct {
	XMLName  xml.Name `xml:"eSearchResult"`
	Count    int      `xml:"Count"`
	RetMax   int      `xml:"RetMax"`
	RetStart int      `xml:"RetStart"`
	IdList   struct {
		Ids []string `xml:"Id"`
	} `xml:"IdList"`
}

type ESummaryResult struct {
	XMLName xml.Name `xml:"eSummaryResult"`
	DocSums []DocSum `xml:"DocSum"`
}

type DocSum struct {
	Id    string `xml:"Id"`
	Items []Item `xml:"Item"`
}

type Item struct {
	Name     string `xml:"Name,attr"`
	Type     string `xml:"Type,attr"`
	Value    string `xml:",chardata"`
	SubItems []Item `xml:"Item"`
}

// Main function
func main() {
	// At the beginning of main()
	if _, err := exec.LookPath("esearch"); err != nil {
		log.Fatal("NCBI E-utilities not found. Please install them: https://www.ncbi.nlm.nih.gov/books/NBK179288/")
	}

	/*
		inputFile := flag.String("input", "exerkines_list.txt", "Input file with gene list")
		outputDir := flag.String("output", "./data/processed_data", "Output directory")
		workers := flag.Int("workers", 5, "Number of concurrent workers")
		flag.Parse()
	*/

	// Load the list of genes/proteins of interest
	geneList := loadInputList("exerkines_list.txt")

	// Process each database type
	fmt.Println("Retrieving Gene References...")
	fetchGeneReferences(geneList)

	fmt.Println("\nRetrieving Protein IDs and Sequences...")
	fetchProteinData(geneList)

	fmt.Println("\nRetrieving Pathway Maps...")
	fetchPathwayMaps(geneList)

	fmt.Println("\nGathering Functional Insights...")
	fetchFunctionalInsights(geneList)
}

func fetchGeneReferences(geneList []string) {
	outputFile, err := os.Create("gene_references.tsv")
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer outputFile.Close()

	// Write header
	outputFile.WriteString("Query\tGene_ID\tSymbol\tDescription\tChromosome\tMapLocation\n")

	// Create a mutex for safe file writing
	var fileMutex sync.Mutex

	// Create progress tracker
	progress := NewProgressTracker(len(geneList))

	// Process genes in parallel with 5 workers
	processGenesInParallel(geneList, 5, func(genes []string, results chan string) {
		for _, gene := range genes {
			// Run esearch with retry and rate limiting
			cmd := exec.Command("sh", "-c",
				fmt.Sprintf("esearch -db gene -query \"%s[Gene Name] AND \"Homo sapiens\"[Organism]\" | efetch -format xml", gene))
			output, err := execWithRetry(cmd, 3) // Try up to 3 times
			if err != nil {
				results <- fmt.Sprintf("Error searching for gene %s: %v\n", gene, err)
				continue
			}

			// Parse XML and write to file...
			var result struct {
				Entrezgenes []struct {
					EntrezgeneTrack struct {
						GeneTrack struct {
							GeneID string `xml:"geneid"`
						} `xml:"Gene-track"`
					} `xml:"Entrezgene_track"`
					EntrezgeneGene struct {
						GeneRef struct {
							GeneDesc  string `xml:"Gene-ref_desc"`
							GeneLocus string `xml:"Gene-ref_locus"`
							GeneMap   struct {
								MapLoc string `xml:"Maps_display-str"`
							} `xml:"Gene-ref_maploc"`
						} `xml:"Gene-ref"`
					} `xml:"Entrezgene_gene"`
					EntrezgeneSource struct {
						BioSource struct {
							SubSource []struct {
								SubtypeName  string `xml:"SubSource_subtype>SubSource_subtype_name"`
								SubtypeValue string `xml:"SubSource_name"`
							} `xml:"BioSource_subtype>SubSource"`
						} `xml:"Entrezgene_source_biosrc"`
					} `xml:"Entrezgene_source"`
				} `xml:"Entrezgene"`
			}

			if err := xml.Unmarshal(output, &result); err != nil {
				fmt.Printf("Error parsing XML for gene %s: %v\n", gene, err)
				continue
			}

			// Write results to file
			for _, eg := range result.Entrezgenes {
				geneID := eg.EntrezgeneTrack.GeneTrack.GeneID
				symbol := eg.EntrezgeneGene.GeneRef.GeneLocus
				description := eg.EntrezgeneGene.GeneRef.GeneDesc
				mapLocation := eg.EntrezgeneGene.GeneRef.GeneMap.MapLoc

				// Find chromosome
				chromosome := "Unknown"
				for _, src := range eg.EntrezgeneSource.BioSource.SubSource {
					if src.SubtypeName == "chromosome" {
						chromosome = src.SubtypeValue
						break
					}
				}

				// Use mutex when writing to file
				fileMutex.Lock()
				outputFile.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\n",
					gene, geneID, symbol, description, chromosome, mapLocation))
				fileMutex.Unlock()
			}

			progress.Increment()
			results <- fmt.Sprintf("Processed gene: %s\n", gene)

		}
	})

	fmt.Printf("\nGene references saved to gene_references.tsv\n")
}

// Add this function to process genes concurrently with worker pool
func processGenesInParallel(geneList []string, workerCount int, processFunc func([]string, chan string)) {
	// Create a channel to receive genes to process
	jobs := make(chan string, len(geneList))
	results := make(chan string, len(geneList))

	// Create worker pool
	var wg sync.WaitGroup
	for w := 1; w <= workerCount; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for gene := range jobs {
				// Process the gene
				processFunc([]string{gene}, results)
			}
		}(w)
	}

	// Send jobs to workers
	for _, gene := range geneList {
		jobs <- gene
	}
	close(jobs)

	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Print results
	for result := range results {
		fmt.Print(result)
	}
}

// Add this function to sanitize XML before parsing
func sanitizeXML(data []byte) []byte {
	// Replace common XML entities that cause parsing issues
	sanitized := string(data)

	// Fix &usehistory - common in NCBI responses
	sanitized = strings.ReplaceAll(sanitized, "&usehistory", "&amp;usehistory")

	// Fix other potential issues
	sanitized = strings.ReplaceAll(sanitized, " & ", " &amp; ")

	return []byte(sanitized)
}

// Add retry logic for transient failures
func execWithRetry(cmd *exec.Cmd, maxRetries int) ([]byte, error) {
	var output []byte
	var err error

	for i := 0; i < maxRetries; i++ {
		output, err = cmd.CombinedOutput()
		if err == nil {
			return output, nil
		}

		fmt.Printf("Command failed (attempt %d/%d): %v\n", i+1, maxRetries, err)
		time.Sleep(time.Duration(i*2) * time.Second) // Exponential backoff
	}

	return output, err
}

// Add a progress tracker
type ProgressTracker struct {
	total     int
	completed int
	mu        sync.Mutex
}

func NewProgressTracker(total int) *ProgressTracker {
	return &ProgressTracker{total: total}
}

func (p *ProgressTracker) Increment() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.completed++
	fmt.Printf("\rProgress: %d/%d (%.1f%%)", p.completed, p.total, float64(p.completed)/float64(p.total)*100)
}

/***
// Add rate limiting to avoid overwhelming NCBI servers
func fetchWithRateLimit(cmd *exec.Cmd) ([]byte, error) {
	// Static rate limiter
	time.Sleep(334 * time.Millisecond) // ~3 requests per second
	return cmd.CombinedOutput()
}

// Use streaming XML decoder for large responses
func parseXMLStream(data []byte, handler func(*xml.Decoder) error) error {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	return handler(decoder)
}
***/

// Load input list from file
func loadInputList(filename string) []string {
	// Check if file exists, if not create example file
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("Input file %s not found. Creating sample file.\n", filename)
		sampleList := []string{
			"ACVR2A", "ACVR2B", "ADIPOR1", "ADIPOR2", "ADPN", "ADRPIN", "AGE", "AHSG",
			"ANGPTL4", "APLN", "APLNR", "APOER2", "ATP6AP2", "BDNF", "BGLAP", "BMP7",
			"BMP8A", "BMP8B", "BMPR1A", "BMPR2", "CCL2", "CCL11", "CCR1", "CCR2", "CCR3",
			"CCR4", "CCR5", "CD44", "CD209", "ChemR23", "CMLKR1", "CNDP2", "CNTF", "CNTFR",
			"CSF3", "CSF3R", "CTGF", "CX3CL1", "CX3CR1", "CXCL1", "CXCR1", "CXCR2", "DCN",
			"EGFR", "ERBB4", "FAS", "FASR", "FGF1", "FGF2", "FGF19", "FGF21", "FGFR1",
			"FGFR2", "FGFR3", "FGFR4", "GALR2", "GALR3", "GDF8", "GDF11", "GDF15", "GFRAL",
			"gp130", "GPR1", "GPR41", "GPR43", "GPRC6A", "GRP78", "ICAM1", "IFNG", "IFNGR1",
			"IFNGR2", "IGF1", "IGF1R", "IL1", "IL1A", "IL1B", "IL1R1", "IL1R2", "IL1RA",
			"IL4", "IL4R", "IL6", "IL6R", "IL6SR", "IL7", "IL7R", "IL8", "IL10", "IL10R",
			"IL10R2", "IL13", "IL13RA1", "IL13RA2", "IL15", "IL15RA", "IL18", "IL18R1",
			"IL18RAP", "INHBA", "INHBB", "INHBE", "INS", "INSR", "ITGAL", "ITGAM", "ITGAV",
			"ITGB2", "KL", "LECT2", "LIF", "LIFR", "LPS", "LRP1", "LRP4", "LRP5", "LRP6",
			"MCP1", "MDC", "METRNL", "MSTN", "NRG4", "NPR3", "NTF3", "OPN", "OSTN",
			"SPARC", "PGRN", "RAGE", "RANKL", "RANTES", "RARRES2", "RBP4", "REN", "S100A",
			"S100B", "SCFAs", "SEP", "SERPINA12", "SORT1", "SOST", "SPARC", "SPP1", "SPX",
			"STRA6", "TGFB1", "TGFB2", "TGFBR1", "TGFBR2", "TLR4", "TNFA", "TNFR1",
			"TNFR2", "TNFRSF11B", "TRKB", "TRKC", "VDR", "VEGF", "VEGFA", "VEGFR1", "VEGFR2",
		}
		file, err := os.Create(filename)
		if err != nil {
			log.Fatalf("Failed to create sample file: %v", err)
		}
		defer file.Close()

		for _, gene := range sampleList {
			file.WriteString(gene + "\n")
		}
		fmt.Printf("Created sample file with common exerkines. Edit %s to customize your search.\n\n", filename)
	}

	// Read file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open input file: %v", err)
	}
	defer file.Close()

	var list []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		item := strings.TrimSpace(scanner.Text())
		if item != "" {
			list = append(list, item)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	fmt.Printf("Loaded %d items from %s\n", len(list), filename)
	return list
}

// Replace your existing fetchProteinData function with this:
func fetchProteinData(geneList []string) {
	// Create protein info file
	infoFile, err := os.Create("protein_info.tsv")
	if err != nil {
		log.Fatalf("Failed to create protein info file: %v", err)
	}
	defer infoFile.Close()

	// Create FASTA file
	fastaFile, err := os.Create("protein_sequences.fasta")
	if err != nil {
		log.Fatalf("Failed to create FASTA file: %v", err)
	}
	defer fastaFile.Close()

	// Write header for info file
	infoFile.WriteString("Query\tProtein_ID\tAccession\tName\tLength\tMolecular_Weight\n")

	// Create mutexes for safe file writing
	var infoMutex, fastaMutex sync.Mutex

	// Create progress tracker
	progress := NewProgressTracker(len(geneList))

	// Process proteins in parallel with 5 workers
	processGenesInParallel(geneList, 5, func(genes []string, results chan string) {
		for _, gene := range genes {
			// Run esearch with retry and rate limiting
			cmd := exec.Command("sh", "-c",
				fmt.Sprintf("esearch -db protein -query \"%s[Gene Name] AND \"Homo sapiens\"[Organism] AND refseq[Filter]\" | efetch -format xml", gene))
			output, err := execWithRetry(cmd, 3) // Try up to 3 times
			if err != nil {
				results <- fmt.Sprintf("Error searching for protein %s: %v\n", gene, err)
				continue
			}

			// Parse protein data
			var proteinResult struct {
				ProteinList []struct {
					ProtID        string  `xml:"Bioseq_id>Seq-id>Seq-id_other>Seq-id_other_gi"`
					ProtAccession string  `xml:"Bioseq_id>Seq-id>Seq-id_other>Textseq-id>Textseq-id_accession"`
					ProtName      string  `xml:"Bioseq_id>Seq-id>Seq-id_other>Textseq-id>Textseq-id_name"`
					ProtLength    int     `xml:"Bioseq_length"`
					ProtSeqData   string  `xml:"Bioseq_seq-data>Seq-data>Seq-data_iupacaa>IUPACaa"`
					ProtMolWeight float64 `xml:"Bioseq_annot>Seq-annot>Seq-annot_data>Seq-annot_data_ftable>Seq-feat>Seq-feat_ext>User-object>User-field>User-field_data>User-field_data_real"`
				} `xml:"Bioseq-set_seq-set>Bioseq"`
			}

			if err := xml.Unmarshal(sanitizeXML(output), &proteinResult); err != nil {
				results <- fmt.Sprintf("Error parsing protein XML for %s: %v\n", gene, err)
				continue
			}

			// Write results to files
			for _, prot := range proteinResult.ProteinList {
				// Use mutex when writing to info file
				infoMutex.Lock()
				infoFile.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\t%d\t%.2f\n",
					gene, prot.ProtID, prot.ProtAccession, prot.ProtName, prot.ProtLength, prot.ProtMolWeight))
				infoMutex.Unlock()

				// Use mutex when writing to FASTA file
				fastaMutex.Lock()
				fastaFile.WriteString(fmt.Sprintf(">%s|%s|%s|%s\n", gene, prot.ProtID, prot.ProtAccession, prot.ProtName))

				// Get sequence via efetch with retry
				seqCmd := exec.Command("sh", "-c",
					fmt.Sprintf("efetch -db protein -id %s -format fasta", prot.ProtID))
				seqOutput, err := execWithRetry(seqCmd, 3)
				if err != nil {
					fastaMutex.Unlock()
					results <- fmt.Sprintf("Error fetching sequence for %s: %v\n", prot.ProtID, err)
					continue
				}

				// Skip the header line and write the sequence
				seqLines := strings.Split(string(seqOutput), "\n")
				for i := 1; i < len(seqLines); i++ {
					if len(seqLines[i]) > 0 {
						fastaFile.WriteString(seqLines[i] + "\n")
					}
				}
				fastaMutex.Unlock()
			}

			progress.Increment()
			results <- fmt.Sprintf("Processed protein: %s\n", gene)
		}
	})

	fmt.Printf("\nProtein information saved to protein_info.tsv\n")
	fmt.Printf("Protein sequences saved to protein_sequences.fasta\n")
}

// Fetch pathway maps
func fetchPathwayMaps(geneList []string) {
	outputFile, err := os.Create("pathway_maps.tsv")
	if err != nil {
		log.Fatalf("Failed to create pathway file: %v", err)
	}
	defer outputFile.Close()

	// Write header
	outputFile.WriteString("Gene\tPathway_ID\tPathway_Name\tPathway_Source\tGene_Role\n")

	for _, gene := range geneList {
		fmt.Printf("Fetching pathways for: %s\n", gene)

		// Search for pathways in NCBI Biosystems
		cmd := exec.Command("sh", "-c",
			fmt.Sprintf("esearch -db biosystems -query \"%s[Gene Name] AND \"Homo sapiens\"[Organism]\" | elink -target gene | efetch -format xml", gene))
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error searching for pathways for %s: %v\n", gene, err)
			continue
		}

		// Parse the XML
		var result struct {
			BioSystems []struct {
				BSID     string `xml:"System_id"`
				BSName   string `xml:"System_name"`
				BSSource struct {
					SourceName string `xml:"BioSource_name"`
				} `xml:"System_source>BioSource"`
				BSGenes []struct {
					GeneRole string `xml:"System-links_db-memberships_value"`
				} `xml:"System_ent>System-ent_genes>System-links_db-memberships"`
			} `xml:"Biosystem"`
		}

		if err := xml.Unmarshal(output, &result); err != nil {
			fmt.Printf("Error parsing pathway XML for %s: %v\n", gene, err)
			continue
		}

		// Write results to file
		for _, bs := range result.BioSystems {
			role := "Member"
			if len(bs.BSGenes) > 0 {
				role = bs.BSGenes[0].GeneRole
			}

			outputFile.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\t%s\n",
				gene, bs.BSID, bs.BSName, bs.BSSource.SourceName, role))
		}

		// Alternative search in KEGG
		keggCmd := exec.Command("sh", "-c",
			fmt.Sprintf("esearch -db gene -query \"%s[Gene Name] AND \"Homo sapiens\"[Organism]\" | elink -target pathway | esummary -format xml", gene))
		keggOutput, err := keggCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error searching KEGG pathways for %s: %v\n", gene, err)
			continue
		}

		var keggResult ESummaryResult
		if err := xml.Unmarshal(keggOutput, &keggResult); err != nil {
			fmt.Printf("Error parsing KEGG XML for %s: %v\n", gene, err)
			continue
		}

		for _, docsum := range keggResult.DocSums {
			pathwayID := docsum.Id
			var pathwayName, pathwaySource string

			for _, item := range docsum.Items {
				if item.Name == "Full_name_E" {
					pathwayName = item.Value
				}
				if item.Name == "Source" {
					pathwaySource = item.Value
				}
			}

			outputFile.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\t%s\n",
				gene, pathwayID, pathwayName, pathwaySource, "Member"))
		}
	}
	fmt.Printf("Pathway information saved to pathway_maps.tsv\n")
}

// Fetch functional insights
func fetchFunctionalInsights(geneList []string) {
	outputFile, err := os.Create("functional_insights.tsv")
	if err != nil {
		log.Fatalf("Failed to create functional insights file: %v", err)
	}
	defer outputFile.Close()

	// Write header
	outputFile.WriteString("Gene\tFunction_Type\tDescription\tEvidence\tReference_PMID\n")

	for _, gene := range geneList {
		fmt.Printf("Fetching functional insights for: %s\n", gene)

		// Search PubMed for functional studies
		cmd := exec.Command("sh", "-c",
			fmt.Sprintf("esearch -db pubmed -query \"%s[Gene Name] AND function AND (\"Homo sapiens\"[Organism] OR human)\" | efetch -format xml", gene))
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error searching for functional insights for %s: %v\n", gene, err)
			continue
		}

		// Parse the XML
		var result struct {
			Articles []struct {
				PMID    string `xml:"MedlineCitation>PMID"`
				Article struct {
					Title    string `xml:"ArticleTitle"`
					Abstract struct {
						AbstractText []struct {
							Label string `xml:"Label,attr"`
							Text  string `xml:",chardata"`
						} `xml:"AbstractText"`
					} `xml:"Abstract"`
				} `xml:"MedlineCitation>Article"`
				KeywordList struct {
					Keywords []string `xml:"Keyword"`
				} `xml:"MedlineCitation>KeywordList"`
			} `xml:"PubmedArticle"`
		}

		if err := xml.Unmarshal(output, &result); err != nil {
			fmt.Printf("Error parsing functional XML for %s: %v\n", gene, err)
			continue
		}

		// Write results to file
		for _, article := range result.Articles {
			// Extract function type and evidence from abstract sections or keywords
			functionType := "Molecular Function"
			evidence := "Literature"

			var description string

			// Try to get functional description from abstract
			if len(article.Article.Abstract.AbstractText) > 0 {
				for _, section := range article.Article.Abstract.AbstractText {
					if section.Label == "RESULTS" || section.Label == "CONCLUSION" || section.Label == "CONCLUSIONS" {
						description = section.Text
						break
					}
				}

				// If no specific section found, use the first section
				if description == "" && len(article.Article.Abstract.AbstractText) > 0 {
					description = article.Article.Abstract.AbstractText[0].Text
				}
			}

			// If still no description, use article title
			if description == "" {
				description = article.Article.Title
			}

			// Look for functional keywords
			for _, keyword := range article.KeywordList.Keywords {
				lowerKeyword := strings.ToLower(keyword)
				if strings.Contains(lowerKeyword, "signal") || strings.Contains(lowerKeyword, "pathway") {
					functionType = "Signaling"
				} else if strings.Contains(lowerKeyword, "metabol") {
					functionType = "Metabolism"
				} else if strings.Contains(lowerKeyword, "immune") || strings.Contains(lowerKeyword, "inflamm") {
					functionType = "Immune Regulation"
				} else if strings.Contains(lowerKeyword, "exercis") || strings.Contains(lowerKeyword, "muscle") {
					functionType = "Exercise Response"
				}
			}

			outputFile.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\t%s\n",
				gene, functionType, description, evidence, article.PMID))
		}

		// Get Gene Ontology annotations
		goCmd := exec.Command("sh", "-c",
			fmt.Sprintf("esearch -db gene -query \"%s[Gene Name] AND \"Homo sapiens\"[Organism]\" | elink -target geneontology | efetch -format xml", gene))
		goOutput, err := goCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error searching GO terms for %s: %v\n", gene, err)
			continue
		}

		var goResult struct {
			Terms []struct {
				TermID       string `xml:"GO_term_id"`
				TermName     string `xml:"GO_term_name"`
				TermCategory string `xml:"GO_term_category"`
				Evidence     string `xml:"GO_term_evidence"`
				Source       string `xml:"GO_term_source"`
			} `xml:"GoTerm"`
		}

		if err := xml.Unmarshal(goOutput, &goResult); err != nil {
			fmt.Printf("Error parsing GO XML for %s: %v\n", gene, err)
			continue
		}

		for _, term := range goResult.Terms {
			functionType := term.TermCategory
			if functionType == "Function" {
				functionType = "Molecular Function"
			} else if functionType == "Process" {
				functionType = "Biological Process"
			} else if functionType == "Component" {
				functionType = "Cellular Component"
			}

			outputFile.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\t%s\n",
				gene, functionType, term.TermName, term.Evidence, term.Source))
		}
	}
	fmt.Printf("Functional insights saved to functional_insights.tsv\n")
}
