package bloodstream

// ImmuneCell represents a circulating immune cell affected by exercise
type ImmuneCell struct {
	Name          string
	BaselineCount struct {
		Mean  float64
		Range [2]float64
		Unit  string
	}
	ExerciseResponse struct {
		AcuteChange       string  // Direction of change during exercise
		AcuteMagnitude    float64 // Fold change during exercise
		RecoveryTime      float64 // Hours to return to baseline
		ChronicAdaptation string  // Long-term training adaptation
	}
	PrimaryFunctions []string
	ExerciseMediated []string // Exercise-mediated effects
	Receptors        []string // Key receptors expressed
	Secreted         []string // Key factors secreted
}

// Key immune cells affected by exercise
var (
	Neutrophils = ImmuneCell{
		Name: "Neutrophils",
		BaselineCount: struct {
			Mean  float64
			Range [2]float64
			Unit  string
		}{
			Mean:  4000.0,
			Range: [2]float64{1500.0, 7000.0},
			Unit:  "cells/μL",
		},
		ExerciseResponse: struct {
			AcuteChange       string
			AcuteMagnitude    float64
			RecoveryTime      float64
			ChronicAdaptation string
		}{
			AcuteChange:       "Increase",
			AcuteMagnitude:    2.5, // 2.5-fold increase during intense exercise
			RecoveryTime:      6.0, // Hours
			ChronicAdaptation: "Attenuated response",
		},
		PrimaryFunctions: []string{
			"Phagocytosis", "Degranulation", "NETosis",
			"Pathogen killing", "Inflammatory response",
		},
		ExerciseMediated: []string{
			"Enhanced phagocytosis", "Delayed apoptosis",
			"ROS production", "Tissue repair signals",
		},
		Receptors: []string{
			"CXCR1", "CXCR2", "Fc receptors", "TLRs", "Complement receptors",
		},
		Secreted: []string{
			"Myeloperoxidase", "Elastase", "ROS", "NETs", "Cytokines",
		},
	}

	Monocytes = ImmuneCell{
		Name: "Monocytes",
		BaselineCount: struct {
			Mean  float64
			Range [2]float64
			Unit  string
		}{
			Mean:  500.0,
			Range: [2]float64{200.0, 800.0},
			Unit:  "cells/μL",
		},
		ExerciseResponse: struct {
			AcuteChange       string
			AcuteMagnitude    float64
			RecoveryTime      float64
			ChronicAdaptation string
		}{
			AcuteChange:       "Increase",
			AcuteMagnitude:    1.5, // 1.5-fold increase
			RecoveryTime:      4.0, // Hours
			ChronicAdaptation: "Anti-inflammatory phenotype",
		},
		PrimaryFunctions: []string{
			"Phagocytosis", "Antigen presentation",
			"Cytokine production", "Tissue macrophage precursor",
		},
		ExerciseMediated: []string{
			"Phenotype shift (M1→M2)", "Enhanced tissue infiltration",
			"Anti-inflammatory cytokine secretion",
		},
		Receptors: []string{
			"CCR2", "CX3CR1", "CD14", "CD16", "TLRs",
		},
		Secreted: []string{
			"IL-10", "IL-1ra", "IL-6", "TNF-α", "TGF-β",
		},
	}

	NaturalKillerCells = ImmuneCell{
		Name: "Natural Killer Cells",
		BaselineCount: struct {
			Mean  float64
			Range [2]float64
			Unit  string
		}{
			Mean:  200.0,
			Range: [2]float64{100.0, 400.0},
			Unit:  "cells/μL",
		},
		ExerciseResponse: struct {
			AcuteChange       string
			AcuteMagnitude    float64
			RecoveryTime      float64
			ChronicAdaptation string
		}{
			AcuteChange:       "Increase",
			AcuteMagnitude:    3.0, // 3-fold increase during intense exercise
			RecoveryTime:      3.0, // Hours
			ChronicAdaptation: "Enhanced cytotoxicity",
		},
		PrimaryFunctions: []string{
			"Cytotoxicity against infected/tumor cells",
			"Cytokine production", "Immune surveillance",
		},
		ExerciseMediated: []string{
			"Enhanced cytotoxicity", "Increased mobilization",
			"Improved surveillance", "Anti-tumor activity",
		},
		Receptors: []string{
			"CD16", "CD56", "KIRs", "NKG2D", "Interleukin receptors",
		},
		Secreted: []string{
			"IFN-γ", "TNF-α", "Perforin", "Granzymes", "GM-CSF",
		},
	}

	CD4THelper = ImmuneCell{
		Name: "CD4+ T Helper Cells",
		BaselineCount: struct {
			Mean  float64
			Range [2]float64
			Unit  string
		}{
			Mean:  800.0,
			Range: [2]float64{500.0, 1500.0},
			Unit:  "cells/μL",
		},
		ExerciseResponse: struct {
			AcuteChange       string
			AcuteMagnitude    float64
			RecoveryTime      float64
			ChronicAdaptation string
		}{
			AcuteChange:       "Increase",
			AcuteMagnitude:    1.5, // 1.5-fold increase
			RecoveryTime:      2.0, // Hours
			ChronicAdaptation: "Th1/Th2 balance shift",
		},
		PrimaryFunctions: []string{
			"Cytokine secretion", "B-cell help",
			"Macrophage activation", "Inflammatory regulation",
		},
		ExerciseMediated: []string{
			"Shift towards anti-inflammatory phenotype",
			"Reduced Th17/increased Treg", "Enhanced memory formation",
		},
		Receptors: []string{
			"CD4", "CD28", "TCR", "IL receptors", "Chemokine receptors",
		},
		Secreted: []string{
			"IL-2", "IL-4", "IL-10", "IFN-γ", "TNF-α",
		},
	}

	CD8TCytotoxic = ImmuneCell{
		Name: "CD8+ Cytotoxic T Cells",
		BaselineCount: struct {
			Mean  float64
			Range [2]float64
			Unit  string
		}{
			Mean:  600.0,
			Range: [2]float64{300.0, 1000.0},
			Unit:  "cells/μL",
		},
		ExerciseResponse: struct {
			AcuteChange       string
			AcuteMagnitude    float64
			RecoveryTime      float64
			ChronicAdaptation string
		}{
			AcuteChange:       "Increase",
			AcuteMagnitude:    2.0, // 2-fold increase
			RecoveryTime:      3.0, // Hours
			ChronicAdaptation: "Enhanced memory compartment",
		},
		PrimaryFunctions: []string{
			"Cytotoxicity against virus-infected cells",
			"Tumor cell killing", "Memory formation",
		},
		ExerciseMediated: []string{
			"Increased mobilization", "Enhanced cytotoxicity",
			"Improved viral clearance", "Expanded memory pool",
		},
		Receptors: []string{
			"CD8", "CD28", "TCR", "IL receptors", "CXCR3",
		},
		Secreted: []string{
			"Perforin", "Granzymes", "IFN-γ", "TNF-α", "IL-2",
		},
	}

	BLymphocytes = ImmuneCell{
		Name: "B Lymphocytes",
		BaselineCount: struct {
			Mean  float64
			Range [2]float64
			Unit  string
		}{
			Mean:  200.0,
			Range: [2]float64{100.0, 400.0},
			Unit:  "cells/μL",
		},
		ExerciseResponse: struct {
			AcuteChange       string
			AcuteMagnitude    float64
			RecoveryTime      float64
			ChronicAdaptation string
		}{
			AcuteChange:       "Increase",
			AcuteMagnitude:    1.3, // 1.3-fold increase
			RecoveryTime:      2.0, // Hours
			ChronicAdaptation: "Enhanced antibody response",
		},
		PrimaryFunctions: []string{
			"Antibody production", "Antigen presentation",
			"Cytokine secretion", "Memory formation",
		},
		ExerciseMediated: []string{
			"Increased mobilization", "Enhanced antibody production",
			"Improved vaccination response",
		},
		Receptors: []string{
			"BCR", "CD19", "CD20", "CD40", "TLRs",
		},
		Secreted: []string{
			"Antibodies", "IL-6", "IL-10", "TNF-α", "Lymphotoxin",
		},
	}
)

// CalculateImmuneResponse predicts immune cell counts during/after exercise
func CalculateImmuneResponse(
	cell ImmuneCell,
	exerciseIntensity float64, // 0-100%
	exerciseDuration float64, // Minutes
	timeFromStart float64, // Minutes from exercise start
) float64 {
	// Calculate baseline
	baseline := cell.BaselineCount.Mean

	// Intensity factor (0-1)
	intensityFactor := exerciseIntensity / 100.0

	// Scale the response by exercise intensity and duration
	// More intense exercise = stronger immune response
	responseFactor := intensityFactor * (1.0 + (exerciseDuration / 60.0)) // Normalized to 1 hour

	// Cap response factor for reasonable physiological response
	if responseFactor > 2.0 {
		responseFactor = 2.0
	}

	// Adjust magnitude based on response factor
	adjustedMagnitude := 1.0 + ((cell.ExerciseResponse.AcuteMagnitude - 1.0) * responseFactor)

	// Different phases of immune response
	if timeFromStart <= exerciseDuration {
		// During exercise phase - progressive increase
		progressRatio := timeFromStart / exerciseDuration
		return baseline * (1.0 + ((adjustedMagnitude - 1.0) * progressRatio))

	} else if timeFromStart <= (exerciseDuration + 30.0) {
		// Early recovery phase (first 30 minutes) - often peak for some cell types
		if cell.Name == "Neutrophils" {
			// Neutrophils often continue to increase post-exercise
			return baseline * (adjustedMagnitude + 0.5)
		} else {
			return baseline * adjustedMagnitude
		}

	} else {
		// Recovery phase - exponential decline
		recoveryTimeHours := cell.ExerciseResponse.RecoveryTime
		recoveryTimeMinutes := recoveryTimeHours * 60.0

		timeIntoRecovery := timeFromStart - exerciseDuration - 30.0
		recoveryRatio := timeIntoRecovery / recoveryTimeMinutes

		if recoveryRatio >= 1.0 {
			return baseline // Fully recovered
		}

		// Exponential decay back to baseline
		return baseline + ((baseline*adjustedMagnitude)-baseline)*(1.0-recoveryRatio)
	}
}

// GetImmuneCellsAffectedByExercise returns all immune cells modulated by exercise
func GetImmuneCellsAffectedByExercise() []ImmuneCell {
	return []ImmuneCell{
		Neutrophils,
		Monocytes,
		NaturalKillerCells,
		CD4THelper,
		CD8TCytotoxic,
		BLymphocytes,
	}
}

// PredictImmuneCellTimeseriesForExercise generates a time series of immune responses
func PredictImmuneCellTimeseriesForExercise(
	exerciseIntensity float64, // % of max
	exerciseDuration float64, // Minutes
	recoveryPeriod float64, // Minutes to monitor post-exercise
	timepoints int, // Number of timepoints to simulate
) map[string][]float64 {
	// Calculate total timeline duration
	totalDuration := exerciseDuration + recoveryPeriod

	// Calculate time interval between points
	interval := totalDuration / float64(timepoints-1)

	// Get immune cells
	cells := GetImmuneCellsAffectedByExercise()

	// Initialize results
	results := make(map[string][]float64)

	// Generate time series for each cell type
	for _, cell := range cells {
		cellValues := make([]float64, timepoints)

		for i := 0; i < timepoints; i++ {
			timeFromStart := float64(i) * interval
			cellValues[i] = CalculateImmuneResponse(cell, exerciseIntensity,
				exerciseDuration, timeFromStart)
		}

		results[cell.Name] = cellValues
	}

	return results
}

// PredictImmuneIndexes calculates integrated measures of immune function
func PredictImmuneIndexes(
	exerciseIntensity float64,
	exerciseDuration float64,
	trainingStatus string, // "Untrained", "Moderately trained", "Highly trained"
) map[string]float64 {
	indexes := make(map[string]float64)

	// Calculate neutrophil:lymphocyte ratio (inflammatory status indicator)
	baselineNLR := Neutrophils.BaselineCount.Mean / (CD4THelper.BaselineCount.Mean +
		CD8TCytotoxic.BaselineCount.Mean)

	// Calculate exercise effect on NLR
	exerciseIntensityFactor := exerciseIntensity / 100.0
	durationFactor := exerciseDuration / 60.0 // Normalized to 1 hour

	// NLR increases with intensity and duration
	acuteNLR := baselineNLR * (1.0 + exerciseIntensityFactor*durationFactor)

	// Training status affects both baseline and response
	trainingFactor := 1.0
	switch trainingStatus {
	case "Untrained":
		trainingFactor = 1.3 // Stronger inflammatory response
	case "Moderately trained":
		trainingFactor = 1.0
	case "Highly trained":
		trainingFactor = 0.7 // Attenuated inflammatory response
	}

	acuteNLR *= trainingFactor
	indexes["Neutrophil:Lymphocyte Ratio"] = acuteNLR

	// Calculate immune system vigilance (NK cell function)
	vigilance := 100.0 // Baseline score

	// Acute exercise enhances vigilance
	vigilanceBoost := 20.0 * exerciseIntensityFactor
	if exerciseDuration > 90.0 {
		// But very prolonged exercise can reduce it
		vigilanceBoost -= (exerciseDuration - 90.0) / 30.0 * 10.0
	}

	// Training status affects NK cell response
	switch trainingStatus {
	case "Untrained":
		vigilanceBoost *= 0.8 // Less efficient response
	case "Highly trained":
		vigilanceBoost *= 1.2 // More efficient response
	}

	indexes["Immune Vigilance"] = vigilance + vigilanceBoost

	// Calculate inflammation resolution capacity
	resolutionCapacity := 100.0 // Baseline score

	// Moderate exercise improves, extreme exercise temporarily impairs
	if exerciseIntensity <= 75.0 {
		resolutionCapacity += 15.0 * (exerciseIntensity / 75.0)
	} else {
		// High-intensity effect depends on duration
		if exerciseDuration <= 30.0 {
			// Short HIIT improves resolution
			resolutionCapacity += 20.0
		} else {
			// Prolonged high-intensity temporarily impairs resolution
			resolutionCapacity -= (exerciseDuration - 30.0) / 15.0 * 5.0
		}
	}

	// Training status strongly affects resolution capacity
	switch trainingStatus {
	case "Untrained":
		resolutionCapacity *= 0.8
	case "Moderately trained":
		resolutionCapacity *= 1.0
	case "Highly trained":
		resolutionCapacity *= 1.3
	}

	indexes["Inflammation Resolution"] = resolutionCapacity

	return indexes
}
