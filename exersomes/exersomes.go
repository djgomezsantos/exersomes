package main

import (
	"bufio"
	"encoding/xml"
	"exersomes/exersomes/molecular_types"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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
	var molecules []molecular_types.MolecularType

	// Add different molecular types to the list
	met := molecular_types.Metabolite{ID: "1", Name: "Lactate"}
	mirna := molecular_types.MiRNA{ID: "2", Name: "miR-486-5p"}
	prot := molecular_types.Protein{ID: "3", Name: "IL-6"}
	rna := molecular_types.RNA{ID: "4", Name: "PGC-1α mRNA"}

	// Package the molecular cargo into a virtual Exersome!
	exersome := molecular_types.Vesicles{ID: "5", Name: "Muscle-Exersome", Cargo: []molecular_types.MolecularType{met, mirna, prot, rna}}

	// Add them all to the processing queue
	molecules = append(molecules, met, mirna, prot, rna, exersome)

	// Process the molecules
	for _, molecule := range molecules {
		fmt.Printf("Processing [ID: %s] %s...\n", molecule.GetID(), molecule.GetName())

		// Type switch to handle specific molecular processing
		switch m := molecule.(type) {
		case molecular_types.Protein:
			fmt.Printf("  -> Protein identified! Ready for sequence analysis.\n")
		case molecular_types.RNA:
			fmt.Printf("  -> RNA identified! Ready for transcriptomic mapping.\n")
		case molecular_types.Vesicles:
			fmt.Printf("  -> Vesicle identified! Inspecting cargo capacity: %d items.\n", len(m.Cargo))
		default:
			fmt.Printf("  -> Standard molecule logged.\n")
		}
	}

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
	geneList := loadInputList("data/raw_data/exerkines.txt")

	// Process each database type
	fmt.Println("Retrieving RNA References...")
	fetchRNAReferences(geneList)

	fmt.Println("\nRetrieving Protein IDs and Sequences...")
	fetchProteinData(geneList)

	fmt.Println("\nRetrieving RNA (mRNA) Sequences...")
	fetchRNAData(geneList)

	fmt.Println("\nRetrieving Pathway Maps...")
	fetchPathwayMaps(geneList)

	fmt.Println("\nGathering Functional Insights...")
	fetchFunctionalInsights(geneList)
}

func fetchRNAReferences(geneList []string) {
	outputFile, err := os.Create("rna_references.tsv")
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer outputFile.Close()

	// Write header
	outputFile.WriteString("Query\tOrganism\tRNA_Reference\tDescription\n")

	// Create a mutex for safe file writing
	var fileMutex sync.Mutex

	// Create progress tracker
	progress := NewProgressTracker(len(geneList))

	// Process genes in parallel with 2 workers to respect NCBI rate limits
	processGenesInParallel(geneList, 2, func(genes []string, results chan string) {
		for _, gene := range genes {
			// Run esearch with retry and rate limiting
			cmd := exec.Command("sh", "-c",
				fmt.Sprintf("esearch -db nucleotide -query \"%s[Gene Name] AND (\\\"Homo sapiens\\\"[Organism] OR \\\"Rattus norvegicus\\\"[Organism]) AND mRNA[Filter] AND refseq[Filter]\" | efetch -format xml", gene))
			output, err := execWithRetry(cmd, 5) // Try up to 5 times
			if err != nil {
				results <- fmt.Sprintf("Error searching for RNA ref %s: %v\n", gene, err)
				continue
			}

			var rnaResult struct {
				SeqList []struct {
					Accession string `xml:"Bioseq_id>Seq-id>Seq-id_other>Textseq-id>Textseq-id_accession"`
					Title     string `xml:"Bioseq_descr>Seq-descr>Seqdesc>Seqdesc_title"`
				} `xml:"Bioseq-set_seq-set>Bioseq"`
			}

			if err := xml.Unmarshal(sanitizeXML(output), &rnaResult); err != nil {
				results <- fmt.Sprintf("Error parsing XML for RNA ref %s: %v\n", gene, err)
				continue
			}

			// Write results to file
			fileMutex.Lock()
			for _, seq := range rnaResult.SeqList {
				accession := seq.Accession
				title := seq.Title
				var organism string

				if strings.Contains(title, "Homo sapiens") {
					organism = "Human"
				} else if strings.Contains(title, "Rattus norvegicus") {
					organism = "Rat"
				} else {
					organism = "Unknown"
				}

				if accession != "" {
					outputFile.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\n", gene, organism, accession, title))
				}
			}
			fileMutex.Unlock()

			progress.Increment()
			results <- fmt.Sprintf("Processed RNA ref: %s\n", gene)

		}
	})

	fmt.Printf("\nRNA references saved to rna_references.tsv\n")
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
	sanitized = strings.ReplaceAll(sanitized, "&api_key", "&amp;api_key")
	sanitized = strings.ReplaceAll(sanitized, "&db=", "&amp;db=")
	sanitized = strings.ReplaceAll(sanitized, "&term=", "&amp;term=")
	sanitized = strings.ReplaceAll(sanitized, "&retmax=", "&amp;retmax=")
	sanitized = strings.ReplaceAll(sanitized, "&query_key=", "&amp;query_key=")
	sanitized = strings.ReplaceAll(sanitized, "&tool=", "&amp;tool=")
	sanitized = strings.ReplaceAll(sanitized, "&edirect=", "&amp;edirect=")
	sanitized = strings.ReplaceAll(sanitized, "&edirect_os=", "&amp;edirect_os=")
	sanitized = strings.ReplaceAll(sanitized, "&email=", "&amp;email=")
	sanitized = strings.ReplaceAll(sanitized, "&WebEnv", "&amp;WebEnv")
	sanitized = strings.ReplaceAll(sanitized, "&retstart", "&amp;retstart")
	sanitized = strings.ReplaceAll(sanitized, "&rettype", "&amp;rettype")
	sanitized = strings.ReplaceAll(sanitized, "&retmode", "&amp;retmode")
	sanitized = strings.ReplaceAll(sanitized, "&id=", "&amp;id=")
	sanitized = strings.ReplaceAll(sanitized, "&dbfrom=", "&amp;dbfrom=")
	sanitized = strings.ReplaceAll(sanitized, "&cmd=", "&amp;cmd=")
	sanitized = strings.ReplaceAll(sanitized, "&linkname=", "&amp;linkname=")

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
		time.Sleep(time.Duration(i*3+2) * time.Second) // Slower exponential backoff
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

		if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
			log.Fatalf("Failed to create directories for %s: %v", filename, err)
		}

		sampleList := []string{
			"ACVR2A", "ACVR2B", "ADIPOR1", "ADIPOR2", "ADPN", "ADRPIN", "AGE", "AHSG",
			"ANGPTL4", "APLN", "APLNR", "APOER2", "ATP6AP2", "BDNF", "BGLAP", "BMP7",
			"BMP8A", "BMP8B", "BMPR1A", "BMPR2", "CCL2", "CCL11", "CCR1", "CCR2", "CCR3",
			"CCR4", "CCR5", "CCR7", "CD44", "CD209", "ChemR23", "CMLKR1", "CNDP2", "CNTF", "CNTFR",
			"CSF3", "CSF3R", "CTGF", "CX3CL1", "CX3CR1", "CXCL1", "CXCL10", "CXCL11", "CXCL12",
			"CXCR1", "CXCR2", "CXCR3", "CXCR4", "DCN", "DKK1", "EGFR", "ERBB2", "ERBB3", "ERBB4",
			"FAS", "FASR", "FABP3", "FGF1", "FGF2", "FGF19", "FGF21", "FGFR1", "FGFR2", "FGFR3", "FGFR4",
			"FGFRL1", "FST", "FSTL1", "GAS6", "GDF8", "GDF11", "GDF15", "GFRAL", "gp130", "GPR1", "GPR41",
			"GPR43", "GPRC6A", "GRP78", "ICAM1", "IFNG", "IFNGR1", "IFNGR2", "IGF1", "IGF1R", "IL1", "IL1A",
			"IL1B", "IL1R1", "IL1R2", "IL1RN", "IL4", "IL4R", "IL6", "IL6R", "IL6SR", "IL7", "IL7R", "IL8",
			"IL10", "IL10R", "IL10R2", "IL11", "IL13", "IL13RA1", "IL13RA2", "IL15", "IL15RA", "IL17A", "IL17RA",
			"IL18", "IL18R1", "IL18RAP", "IL27RA", "IL33", "INHBA", "INHBB", "INHBE", "INS", "INSR", "ITGAL", "ITGAM",
			"ITGAV", "ITGB2", "KL", "LECT2", "LIF", "LIFR", "LPS", "LRP1", "LRP4", "LRP5", "LRP6", "LCN2", "MCP1", "MDC",
			"METRNL", "MMP2", "MMP9", "MMP14", "MSTN", "MUSK", "MYLK", "NDNF", "NRG1", "NRG4", "NPR3", "NTF3", "OPN", "OSTN",
			"SPARC", "PGRN", "RAGE", "RANKL", "RANTES", "RARRES2", "RBP4", "REN", "S100A", "S100B", "SCFAs", "SEP", "SERPINA12",
			"SORT1", "SOST", "SPARC", "SPP1", "SPX", "SEMA3A", "SEMA3F", "SDC1", "SDC4", "SDC3", "SOD3", "THBS1", "THBS4",
			"TGFB1", "TGFB2", "TGFBR1", "TGFBR2", "TLR4", "TNF", "TNFA", "TNFSF12", "TNFSF14", "TNFR1", "TNFR2", "TNFRSF11B", "TSLP",
			"TRKB", "TRKC", "VDR", "VEGF", "VEGFA", "VEGFB", "VEGFC", "VEGFR1", "VEGFR2", "WNT5A", "ACVR1B", "ACVR1C", "ALK1", "ALK4",
			"ALK5", "BMPR1A", "BMPR1B", "CSF1R", "EPOR", "ERBB3", "ERBB4", "GP130", "IFNGR1", "IFNGR2", "IL1RAP", "INSRR", "MPL",
			"NOTCH1", "NOTCH2", "OSMR", "PDGFRA", "PDGFRB", "ROR1", "ROR2", "TREM2", "ANXA1", "FPR2", "LGALS3", "MRC1", "NLRP3", "PTGER4",
			"SOCS3", "STAT3", "TGFB1", "TNFRSF1A", "TNFRSF1B", "AMPK", "PRKAA1", "CAMK2A", "CAMK2B", "CREB1", "ESRRA", "NRF1", "NFE2L2",
			"PGC1A", "PPARGC1A", "PPARA", "PPARD", "PPARG", "SIRT1", "SIRT3", "TFAM", "UCP2", "UCP3", "CD9", "CD63", "CD81", "HSP90AA1",
			"RAB27A", "RAB27B", "TSG101", "CNTF", "CREB3L1", "EPO", "GDNF", "NRTN", "S100B", "VGF", "ANGPT1", "ANGPT2", "TEK", "NOS3", "HIF1A",
			"PECAM1", "SELP", "VCAM1", "CD34", "HGF", "HGFAC", "ITGA7", "NCAM1", "PAX7", "TGFA", "ANGPTL8", "APOA1", "APOE", "C1QTNF5",
			"CHEMERIN", "FGF19", "RBP4", "SERPINA12",
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
			// Extract official gene symbol if it's in the format "Alias (SYMBOL)"
			if start := strings.Index(item, "("); start != -1 {
				if end := strings.Index(item, ")"); end != -1 && end > start {
					item = strings.TrimSpace(item[start+1 : end])
				}
			}

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

	// Process proteins in parallel with 2 workers to respect NCBI rate limits
	processGenesInParallel(geneList, 2, func(genes []string, results chan string) {
		for _, gene := range genes {
			// Run esearch with retry and rate limiting
			cmd := exec.Command("sh", "-c",
				fmt.Sprintf("esearch -db protein -query \"%s[Gene Name] AND (\\\"Homo sapiens\\\"[Organism] OR \\\"Rattus norvegicus\\\"[Organism]) AND refseq[Filter]\" | efetch -format xml", gene))
			output, err := execWithRetry(cmd, 5) // Try up to 5 times
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
				seqOutput, err := execWithRetry(seqCmd, 5)
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

// Fetch RNA (mRNA) sequences
func fetchRNAData(geneList []string) {
	// Create FASTA file for RNA
	fastaFile, err := os.Create("rna_sequences.fasta")
	if err != nil {
		log.Fatalf("Failed to create RNA FASTA file: %v", err)
	}
	defer fastaFile.Close()

	// Create mutex for safe file writing
	var fastaMutex sync.Mutex

	// Create progress tracker
	progress := NewProgressTracker(len(geneList))

	// Process genes in parallel with 2 workers to respect NCBI rate limits
	processGenesInParallel(geneList, 2, func(genes []string, results chan string) {
		for _, gene := range genes {
			// Delay to respect the rate limit
			time.Sleep(400 * time.Millisecond)

			// Search nucleotide database for mRNA refseqs and fetch fasta directly
			cmd := exec.Command("sh", "-c",
				fmt.Sprintf("esearch -db nucleotide -query \"%s[Gene Name] AND (\\\"Homo sapiens\\\"[Organism] OR \\\"Rattus norvegicus\\\"[Organism]) AND mRNA[Filter] AND refseq[Filter]\" | efetch -format fasta", gene))
			output, err := execWithRetry(cmd, 5) // Try up to 5 times
			if err != nil {
				results <- fmt.Sprintf("Error fetching RNA for %s: %v\n", gene, err)
				continue
			}

			// Validate and write to file safely
			outputStr := string(output)
			if len(outputStr) > 0 && strings.Contains(outputStr, ">") {
				fastaMutex.Lock()
				fastaFile.WriteString(outputStr)
				if !strings.HasSuffix(outputStr, "\n") {
					fastaFile.WriteString("\n")
				}
				fastaMutex.Unlock()
			}

			progress.Increment()
			results <- fmt.Sprintf("Processed RNA for: %s\n", gene)
		}
	})

	fmt.Printf("\nRNA sequences saved to rna_sequences.fasta\n")
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

	var fileMutex sync.Mutex
	progress := NewProgressTracker(len(geneList))

	processGenesInParallel(geneList, 2, func(genes []string, results chan string) {
		for _, gene := range genes {
			// Delay to respect the 3 requests/sec limit
			time.Sleep(400 * time.Millisecond)

			// Search for pathways in NCBI Biosystems (Corrected link direction)
			cmd := exec.Command("sh", "-c",
				fmt.Sprintf("esearch -db gene -query \"%s[Gene Name] AND (\\\"Homo sapiens\\\"[Organism] OR \\\"Rattus norvegicus\\\"[Organism])\" | elink -target biosystems | efetch -format xml", gene))
			output, err := execWithRetry(cmd, 5)
			if err != nil {
				results <- fmt.Sprintf("Error searching for pathways for %s: %v\n", gene, err)
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

			if err := xml.Unmarshal(sanitizeXML(output), &result); err == nil {
				// Write results to file safely
				fileMutex.Lock()
				for _, bs := range result.BioSystems {
					role := "Member"
					if len(bs.BSGenes) > 0 {
						role = bs.BSGenes[0].GeneRole
					}
					outputFile.WriteString(fmt.Sprintf("%s\t%s\t%s\t%s\t%s\n",
						gene, bs.BSID, bs.BSName, bs.BSSource.SourceName, role))
				}
				fileMutex.Unlock()
			}

			// Alternative search in KEGG
			keggCmd := exec.Command("sh", "-c",
				fmt.Sprintf("esearch -db gene -query \"%s[Gene Name] AND (\\\"Homo sapiens\\\"[Organism] OR \\\"Rattus norvegicus\\\"[Organism])\" | elink -target pathway | esummary -format xml", gene))
			keggOutput, err := execWithRetry(keggCmd, 5)
			if err != nil {
				results <- fmt.Sprintf("Error searching KEGG pathways for %s: %v\n", gene, err)
				continue
			}

			var keggResult ESummaryResult
			if err := xml.Unmarshal(sanitizeXML(keggOutput), &keggResult); err == nil {
				fileMutex.Lock()
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
				fileMutex.Unlock()
			}

			progress.Increment()
			results <- fmt.Sprintf("Processed pathways for: %s\n", gene)
		}
	})

	fmt.Printf("\nPathway information saved to pathway_maps.tsv\n")
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

	var fileMutex sync.Mutex
	progress := NewProgressTracker(len(geneList))

	processGenesInParallel(geneList, 2, func(genes []string, results chan string) {
		for _, gene := range genes {
			// Delay to respect the 3 requests/sec limit
			time.Sleep(400 * time.Millisecond)

			// Search PubMed for functional studies
			cmd := exec.Command("sh", "-c",
				fmt.Sprintf("esearch -db pubmed -query \"%s[Gene Name] AND function AND (\\\"Homo sapiens\\\"[Organism] OR \\\"Rattus norvegicus\\\"[Organism] OR human OR rat)\" | efetch -format xml", gene))
			output, err := execWithRetry(cmd, 5)
			if err != nil {
				results <- fmt.Sprintf("Error searching for functional insights for %s: %v\n", gene, err)
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

			if err := xml.Unmarshal(sanitizeXML(output), &result); err == nil {
				fileMutex.Lock()
				for _, article := range result.Articles {
					functionType := "Molecular Function"
					evidence := "Literature"
					var description string

					if len(article.Article.Abstract.AbstractText) > 0 {
						for _, section := range article.Article.Abstract.AbstractText {
							if section.Label == "RESULTS" || section.Label == "CONCLUSION" || section.Label == "CONCLUSIONS" {
								description = section.Text
								break
							}
						}

						if description == "" {
							description = article.Article.Abstract.AbstractText[0].Text
						}
					}

					if description == "" {
						description = article.Article.Title
					}

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
				fileMutex.Unlock()
			}

			// Get Gene Ontology annotations
			goCmd := exec.Command("sh", "-c",
				fmt.Sprintf("esearch -db gene -query \"%s[Gene Name] AND (\\\"Homo sapiens\\\"[Organism] OR \\\"Rattus norvegicus\\\"[Organism])\" | elink -target geneontology | efetch -format xml", gene))
			goOutput, err := execWithRetry(goCmd, 5)
			if err == nil {
				var goResult struct {
					Terms []struct {
						TermID       string `xml:"GO_term_id"`
						TermName     string `xml:"GO_term_name"`
						TermCategory string `xml:"GO_term_category"`
						Evidence     string `xml:"GO_term_evidence"`
						Source       string `xml:"GO_term_source"`
					} `xml:"GoTerm"`
				}

				if err := xml.Unmarshal(sanitizeXML(goOutput), &goResult); err == nil {
					fileMutex.Lock()
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
					fileMutex.Unlock()
				}
			}

			progress.Increment()
			results <- fmt.Sprintf("Processed functional insights for: %s\n", gene)
		}
	})

	fmt.Printf("\nFunctional insights saved to functional_insights.tsv\n")
}
