package placenta

// Placentokine represents a signaling molecule secreted from placental tissue
type Placentokine struct {
	Name               string
	GeneCode           string
	MolecularWeight    float64  // In kDa
	SourceCells        []string // Cell types producing this factor
	TargetOrgans       []string
	ExerciseRegulation string // "Up", "Down", "Biphasic"
	TemporalPattern    string // "Acute", "Chronic", "Both"
	PrimaryEffects     []string
	PlacentalEffect    string // Primary effect on placental function
	BaselineRange      struct {
		Min  float64
		Max  float64
		Unit string
	}
}

// Key placentokines affected by exercise
var (
	Apelin = Placentokine{
		Name:               "Apelin",
		GeneCode:           "APLN",
		MolecularWeight:    8.5,
		SourceCells:        []string{"Syncytiotrophoblasts", "Cytotrophoblasts", "Endothelial cells"},
		TargetOrgans:       []string{"Placenta", "Vasculature", "Adipose", "Fetal tissues"},
		ExerciseRegulation: "Up",
		TemporalPattern:    "Both",
		PrimaryEffects:     []string{"Angiogenesis", "Glucose uptake", "Vasodilation", "Anti-inflammatory"},
		PlacentalEffect:    "Enhanced placental blood flow and nutrient transport",
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  200.0,
			Max:  1200.0,
			Unit: "pg/mL",
		},
	}

	Adiponectin = Placentokine{
		Name:               "Adiponectin",
		GeneCode:           "ADIPOQ",
		MolecularWeight:    30.0,
		SourceCells:        []string{"Syncytiotrophoblasts", "Hofbauer cells"},
		TargetOrgans:       []string{"Placenta", "Liver", "Muscle", "Adipose", "Fetal tissues"},
		ExerciseRegulation: "Up",
		TemporalPattern:    "Chronic",
		PrimaryEffects:     []string{"Insulin sensitization", "Anti-inflammatory", "Fatty acid oxidation"},
		PlacentalEffect:    "Improved nutrient transport, reduced inflammation",
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  5.0,
			Max:  15.0,
			Unit: "μg/mL",
		},
	}

	Leptin = Placentokine{
		Name:               "Leptin",
		GeneCode:           "LEP",
		MolecularWeight:    16.0,
		SourceCells:        []string{"Syncytiotrophoblasts", "Cytotrophoblasts"},
		TargetOrgans:       []string{"Hypothalamus", "Placenta", "Adipose", "Fetal tissues"},
		ExerciseRegulation: "Down",
		TemporalPattern:    "Both",
		PrimaryEffects:     []string{"Energy homeostasis", "Appetite regulation", "Angiogenesis"},
		PlacentalEffect:    "Trophoblast invasion, placental growth",
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  5.0,  // Non-pregnant range is lower
			Max:  40.0, // Increases during pregnancy
			Unit: "ng/mL",
		},
	}

	Irisin = Placentokine{
		Name:               "Irisin",
		GeneCode:           "FNDC5",
		MolecularWeight:    12.0,
		SourceCells:        []string{"Syncytiotrophoblasts", "Decidual cells"},
		TargetOrgans:       []string{"Adipose", "Placenta", "Fetal tissues"},
		ExerciseRegulation: "Up",
		TemporalPattern:    "Both",
		PrimaryEffects:     []string{"Thermogenesis", "Glucose homeostasis", "Energy expenditure"},
		PlacentalEffect:    "Metabolic adaptation, mitochondrial biogenesis",
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  0.2,
			Max:  2.0,
			Unit: "μg/mL",
		},
	}

	SOD3 = Placentokine{
		Name:               "Superoxide dismutase 3",
		GeneCode:           "SOD3",
		MolecularWeight:    32.0,
		SourceCells:        []string{"Trophoblasts", "Placental macrophages"},
		TargetOrgans:       []string{"Placenta", "Vasculature", "Fetal tissues"},
		ExerciseRegulation: "Up",
		TemporalPattern:    "Chronic",
		PrimaryEffects:     []string{"Antioxidant defense", "Free radical scavenging", "Tissue protection"},
		PlacentalEffect:    "Protection against oxidative stress, reduced placental dysfunction",
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  20.0,
			Max:  100.0,
			Unit: "ng/mL",
		},
	}

	Chemerin = Placentokine{
		Name:               "Chemerin",
		GeneCode:           "RARRES2",
		MolecularWeight:    16.0,
		SourceCells:        []string{"Extravillous trophoblasts", "Decidual cells"},
		TargetOrgans:       []string{"Adipose", "Placenta", "Immune cells"},
		ExerciseRegulation: "Down",
		TemporalPattern:    "Chronic",
		PrimaryEffects:     []string{"Immune cell migration", "Adipocyte differentiation", "Glucose metabolism"},
		PlacentalEffect:    "Immune regulation, potential role in preeclampsia when elevated",
		BaselineRange: struct {
			Min  float64
			Max  float64
			Unit string
		}{
			Min:  100.0,
			Max:  300.0,
			Unit: "ng/mL",
		},
	}
)

// GetExerciseUpregulatedPlacentokines returns placentokines that increase with exercise
func GetExerciseUpregulatedPlacentokines() []Placentokine {
	return []Placentokine{Apelin, Adiponectin, Irisin, SOD3}
}

// GetExerciseDownregulatedPlacentokines returns placentokines that decrease with exercise
func GetExerciseDownregulatedPlacentokines() []Placentokine {
	return []Placentokine{Leptin, Chemerin}
}

// CalculateExerciseResponse estimates placentokine changes from exercise
func CalculateExerciseResponse(placentokine Placentokine, exerciseType string,
	intensityPercent int, durationMinutes int, gestationalWeek int, trainingWeeks int) float64 {

	// Base level factor (centered at 1.0 = no change)
	var changeFactor float64 = 1.0

	// Acute response factors
	acuteIntensityEffect := float64(intensityPercent) / 100.0
	acuteDurationEffect := float64(durationMinutes) / 60.0

	// Chronic adaptation factors
	chronicEffect := float64(trainingWeeks) / 12.0 // Normalized to 12 weeks

	// Pregnancy trimester factors (pregnancy progresses over time)
	var trimesterFactor float64 = 1.0
	if gestationalWeek <= 13 {
		trimesterFactor = 0.8 // First trimester
	} else if gestationalWeek <= 26 {
		trimesterFactor = 1.0 // Second trimester
	} else {
		trimesterFactor = 1.2 // Third trimester
	}

	// Different placentokines have distinct response patterns to exercise
	switch placentokine.Name {
	case "Apelin":
		// Responsive to both acute and chronic exercise
		acuteEffect := 1.0 + (0.5 * acuteIntensityEffect * acuteDurationEffect)
		chronicEffect := 1.0 + (0.3 * float64(trainingWeeks) / 12.0)

		if exerciseType == "Aerobic" {
			changeFactor = acuteEffect * chronicEffect * 1.1 * trimesterFactor
		} else {
			changeFactor = acuteEffect * chronicEffect * trimesterFactor
		}

	case "Adiponectin":
		// Primarily responds to chronic exercise
		if trainingWeeks > 3 {
			changeFactor = 1.0 + (0.25 * chronicEffect * trimesterFactor)

			// Exercise type effect
			if exerciseType == "Combined" {
				changeFactor *= 1.2 // Combined training has greater effect
			}
		} else {
			changeFactor = 1.0 + (0.1 * acuteIntensityEffect * trimesterFactor)
		}

	case "Leptin":
		// Decreases with exercise, more with chronic training
		acuteEffect := 1.0 - (0.1 * acuteIntensityEffect * acuteDurationEffect)
		chronicEffect := 1.0 - (0.2 * float64(trainingWeeks) / 12.0)

		// Exercise intensity matters more for leptin
		if intensityPercent > 70 {
			changeFactor = acuteEffect * chronicEffect * 0.9 * trimesterFactor
		} else {
			changeFactor = acuteEffect * chronicEffect * trimesterFactor
		}

		// Floor at 40% reduction
		if changeFactor < 0.6 {
			changeFactor = 0.6
		}

	case "Irisin":
		// Strongly responds to resistance training
		if exerciseType == "Resistance" || exerciseType == "Combined" {
			acuteEffect := 1.0 + (0.6 * acuteIntensityEffect * acuteDurationEffect)
			chronicEffect := 1.0 + (0.3 * float64(trainingWeeks) / 12.0)
			changeFactor = acuteEffect * chronicEffect * trimesterFactor
		} else {
			acuteEffect := 1.0 + (0.3 * acuteIntensityEffect * acuteDurationEffect)
			chronicEffect := 1.0 + (0.15 * float64(trainingWeeks) / 12.0)
			changeFactor = acuteEffect * chronicEffect * trimesterFactor
		}

	case "Superoxide dismutase 3":
		// Increases with regular training, modest acute effect
		if trainingWeeks > 2 {
			changeFactor = 1.0 + (0.3 * chronicEffect * trimesterFactor)
		} else {
			changeFactor = 1.0 + (0.1 * acuteIntensityEffect * trimesterFactor)
		}

	case "Chemerin":
		// Decreases with regular exercise
		if trainingWeeks > 4 {
			changeFactor = 1.0 - (0.2 * chronicEffect * trimesterFactor)
			if changeFactor < 0.7 { // Floor at 30% reduction
				changeFactor = 0.7
			}
		} else {
			changeFactor = 1.0 - (0.05 * acuteIntensityEffect * trimesterFactor)
		}
	}

	return changeFactor
}

// PredictPlacentokineResponse estimates placentokine levels after an exercise protocol
func PredictPlacentokineResponse(exerciseType string, intensityPercent int,
	durationMinutes int, gestationalWeek int, trainingWeeks int,
	hasGestationalDiabetes bool) map[string]float64 {

	// Get all placentokines
	placentokines := append(GetExerciseUpregulatedPlacentokines(), GetExerciseDownregulatedPlacentokines()...)

	// Calculate responses
	responses := make(map[string]float64)

	for _, placentokine := range placentokines {
		// Calculate baseline level (middle of reference range)
		baseline := (placentokine.BaselineRange.Min + placentokine.BaselineRange.Max) / 2.0

		// Adjust baseline for gestational diabetes if applicable
		if hasGestationalDiabetes {
			switch placentokine.Name {
			case "Adiponectin":
				baseline *= 0.7 // Decreased in GDM
			case "Leptin":
				baseline *= 1.5 // Elevated in GDM
			case "Apelin":
				baseline *= 1.3 // Elevated in GDM
			case "Chemerin":
				baseline *= 1.4 // Elevated in GDM
			}
		}

		// Calculate response factor
		responseFactor := CalculateExerciseResponse(placentokine, exerciseType,
			intensityPercent, durationMinutes, gestationalWeek, trainingWeeks)

		// Calculate predicted level
		responses[placentokine.Name] = baseline * responseFactor
	}

	return responses
}
