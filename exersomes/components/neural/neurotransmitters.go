package neural

// Neurotransmitter represents a chemical messenger in the nervous system affected by exercise
type Neurotransmitter struct {
	Name               string
	Class              string // "Monoamine", "Amino acid", "Peptide", etc.
	PrimaryFunctions   []string
	SynthesisFrom      string
	KeyEnzymes         []string // Enzymes involved in synthesis/metabolism
	PrimaryRegions     []string // Brain regions with significant transmission
	ExerciseRegulation string   // "Up", "Down", "Biphasic", "Complex"
	AcuteResponse      string   // Acute exercise effect
	ChronicAdaptation  string   // Effect of regular training
	ReleaseTimeframe   string   // When during exercise it peaks
	RecoveryTimeframe  string   // How long elevated after exercise
	MoodEffect         string
	CognitiveEffect    string
	RelatedPathologies []string
}

// Key neurotransmitters affected by exercise
var (
	Dopamine = Neurotransmitter{
		Name:  "Dopamine",
		Class: "Monoamine",
		PrimaryFunctions: []string{
			"Reward and motivation",
			"Motor control",
			"Executive function",
			"Working memory",
		},
		SynthesisFrom: "Tyrosine",
		KeyEnzymes:    []string{"Tyrosine hydroxylase", "DOPA decarboxylase"},
		PrimaryRegions: []string{
			"Ventral tegmental area",
			"Substantia nigra",
			"Striatum",
			"Prefrontal cortex",
		},
		ExerciseRegulation: "Up",
		AcuteResponse:      "Increased release and turnover",
		ChronicAdaptation:  "Enhanced receptor sensitivity and signaling efficiency",
		ReleaseTimeframe:   "During and immediately post-exercise",
		RecoveryTimeframe:  "2-4 hours post-exercise",
		MoodEffect:         "Euphoria, pleasure, motivation",
		CognitiveEffect:    "Enhanced focus, working memory, and behavioral flexibility",
		RelatedPathologies: []string{
			"Parkinson's disease",
			"ADHD",
			"Addiction",
			"Depression",
			"Schizophrenia",
		},
	}

	Serotonin = Neurotransmitter{
		Name:  "Serotonin (5-HT)",
		Class: "Monoamine",
		PrimaryFunctions: []string{
			"Mood regulation",
			"Sleep-wake cycle",
			"Appetite control",
			"Pain modulation",
		},
		SynthesisFrom: "Tryptophan",
		KeyEnzymes:    []string{"Tryptophan hydroxylase", "Aromatic L-amino acid decarboxylase"},
		PrimaryRegions: []string{
			"Raphe nuclei",
			"Hypothalamus",
			"Limbic system",
			"Prefrontal cortex",
		},
		ExerciseRegulation: "Complex",
		AcuteResponse:      "Initial increase, possible decrease with prolonged exercise",
		ChronicAdaptation:  "Increased synthesis and receptor sensitivity",
		ReleaseTimeframe:   "During exercise, especially prolonged activity",
		RecoveryTimeframe:  "24+ hours for full normalization",
		MoodEffect:         "Anti-depressant, anxiolytic, satiety",
		CognitiveEffect:    "Improved impulse control, reduced rumination",
		RelatedPathologies: []string{
			"Depression",
			"Anxiety disorders",
			"OCD",
			"Eating disorders",
		},
	}

	Norepinephrine = Neurotransmitter{
		Name:  "Norepinephrine",
		Class: "Monoamine",
		PrimaryFunctions: []string{
			"Arousal and alertness",
			"Attention",
			"Stress response",
			"Heart rate and blood pressure",
		},
		SynthesisFrom: "Dopamine",
		KeyEnzymes:    []string{"Dopamine β-hydroxylase"},
		PrimaryRegions: []string{
			"Locus coeruleus",
			"Hypothalamus",
			"Amygdala",
			"Prefrontal cortex",
		},
		ExerciseRegulation: "Up",
		AcuteResponse:      "Strong increase proportional to intensity",
		ChronicAdaptation:  "More efficient noradrenergic signaling, receptor sensitivity",
		ReleaseTimeframe:   "Increases rapidly during exercise",
		RecoveryTimeframe:  "1-3 hours post-exercise",
		MoodEffect:         "Alertness, energy, potentially anxiety if too high",
		CognitiveEffect:    "Enhanced attention, vigilance, and working memory",
		RelatedPathologies: []string{
			"ADHD",
			"Depression",
			"PTSD",
			"Anxiety disorders",
		},
	}

	GABA = Neurotransmitter{
		Name:  "Gamma-Aminobutyric Acid (GABA)",
		Class: "Amino acid",
		PrimaryFunctions: []string{
			"Inhibitory neurotransmission",
			"Anxiety reduction",
			"Muscle relaxation",
			"Seizure prevention",
		},
		SynthesisFrom: "Glutamate",
		KeyEnzymes:    []string{"Glutamic acid decarboxylase (GAD)"},
		PrimaryRegions: []string{
			"Widespread throughout CNS",
			"Cortical interneurons",
			"Basal ganglia",
			"Cerebellum",
		},
		ExerciseRegulation: "Up",
		AcuteResponse:      "Moderate increase post-exercise",
		ChronicAdaptation:  "Enhanced GABAergic tone and receptor function",
		ReleaseTimeframe:   "Primarily post-exercise",
		RecoveryTimeframe:  "2-6 hours post-exercise",
		MoodEffect:         "Anxiolytic, relaxation, calmness",
		CognitiveEffect:    "Reduced rumination, improved stress coping",
		RelatedPathologies: []string{
			"Anxiety disorders",
			"Epilepsy",
			"Insomnia",
			"Huntington's disease",
		},
	}

	Glutamate = Neurotransmitter{
		Name:  "Glutamate",
		Class: "Amino acid",
		PrimaryFunctions: []string{
			"Excitatory neurotransmission",
			"Learning and memory",
			"Synaptic plasticity",
			"Cognitive function",
		},
		SynthesisFrom: "Glutamine, glucose",
		KeyEnzymes:    []string{"Glutaminase", "Glutamate dehydrogenase"},
		PrimaryRegions: []string{
			"Widespread throughout CNS",
			"Hippocampus",
			"Cortex",
			"Cerebellum",
		},
		ExerciseRegulation: "Complex",
		AcuteResponse:      "Balanced with GABA, transient increases",
		ChronicAdaptation:  "Optimized glutamatergic signaling efficiency",
		ReleaseTimeframe:   "During exercise, especially high-intensity",
		RecoveryTimeframe:  "1-3 hours post-exercise",
		MoodEffect:         "Mixed - excessive levels can cause anxiety",
		CognitiveEffect:    "Enhanced learning, memory formation, and plasticity",
		RelatedPathologies: []string{
			"Alzheimer's disease",
			"Schizophrenia",
			"Stroke",
			"Epilepsy",
			"Excitotoxicity",
		},
	}

	Endocannabinoids = Neurotransmitter{
		Name:  "Endocannabinoids (primarily Anandamide)",
		Class: "Lipid",
		PrimaryFunctions: []string{
			"Retrograde synaptic signaling",
			"Mood regulation",
			"Pain reduction",
			"Appetite regulation",
		},
		SynthesisFrom: "Arachidonic acid",
		KeyEnzymes:    []string{"N-acyl phosphatidylethanolamine phospholipase D (NAPE-PLD)"},
		PrimaryRegions: []string{
			"Widespread throughout CNS",
			"Hippocampus",
			"Basal ganglia",
			"Cerebellum",
		},
		ExerciseRegulation: "Up",
		AcuteResponse:      "Strong increase during moderate-intensity exercise",
		ChronicAdaptation:  "Enhanced endocannabinoid tone",
		ReleaseTimeframe:   "During and immediately post-exercise",
		RecoveryTimeframe:  "1-2 hours post-exercise",
		MoodEffect:         "Euphoria, calmness, 'runner's high'",
		CognitiveEffect:    "Reduced anxiety, enhanced reward processing",
		RelatedPathologies: []string{
			"Depression",
			"Anxiety disorders",
			"Chronic pain",
			"PTSD",
		},
	}

	Acetylcholine = Neurotransmitter{
		Name:  "Acetylcholine",
		Class: "Amine",
		PrimaryFunctions: []string{
			"Neuromuscular junction transmission",
			"Attention and memory",
			"Learning and plasticity",
			"Autonomic functions",
		},
		SynthesisFrom: "Choline and acetyl-CoA",
		KeyEnzymes:    []string{"Choline acetyltransferase"},
		PrimaryRegions: []string{
			"Motor neurons",
			"Basal forebrain",
			"Brainstem",
			"Striatum",
		},
		ExerciseRegulation: "Up",
		AcuteResponse:      "Increased release at neuromuscular junctions and in brain",
		ChronicAdaptation:  "Enhanced cholinergic signaling and receptor function",
		ReleaseTimeframe:   "During exercise",
		RecoveryTimeframe:  "0.5-2 hours post-exercise",
		MoodEffect:         "Alertness, attention",
		CognitiveEffect:    "Enhanced attention, memory formation, and cognitive precision",
		RelatedPathologies: []string{
			"Alzheimer's disease",
			"Myasthenia gravis",
			"Parkinson's disease",
		},
	}
)

// GetExerciseResponsiveNeurotransmitters returns neurotransmitters affected by exercise
func GetExerciseResponsiveNeurotransmitters() []Neurotransmitter {
	return []Neurotransmitter{
		Dopamine,
		Serotonin,
		Norepinephrine,
		GABA,
		Glutamate,
		Endocannabinoids,
		Acetylcholine,
	}
}

// CalculateNeuromodulatorResponse estimates neurotransmitter changes during and after exercise
func CalculateNeuromodulatorResponse(nt Neurotransmitter,
	exerciseType string, intensityPercent int, durationMinutes int, timePoint string) float64 {

	// Base level factor (centered at 1.0 = no change)
	changeFactor := 1.0

	// Calculate factors for determining response
	intensityFactor := float64(intensityPercent) / 100.0
	durationFactor := float64(durationMinutes) / 60.0 // Normalized to 1 hour

	// Apply neurotransmitter-specific response patterns
	switch nt.Name {
	case "Dopamine":
		if timePoint == "During" {
			// Increases significantly during exercise
			changeFactor = 1.0 + (0.5 * intensityFactor * (0.7 + 0.3*durationFactor))
		} else if timePoint == "Immediate post" {
			// Remains elevated immediately after
			changeFactor = 1.0 + (0.4 * intensityFactor)
		} else if timePoint == "1 hour post" {
			// Begins returning to baseline
			changeFactor = 1.0 + (0.2 * intensityFactor)
		} else {
			// Returns to baseline or slightly above after 3+ hours
			changeFactor = 1.0 + (0.05 * intensityFactor)
		}

	case "Serotonin (5-HT)":
		if timePoint == "During" {
			if durationMinutes < 60 {
				// Initial increase during shorter exercise
				changeFactor = 1.0 + (0.3 * intensityFactor)
			} else {
				// Complex pattern during prolonged exercise (can decrease with fatigue)
				changeFactor = 1.0 + (0.3 * intensityFactor) - (0.1 * (durationFactor - 1.0))
			}
		} else if timePoint == "Immediate post" {
			// Post-exercise elevation
			changeFactor = 1.0 + (0.3 * intensityFactor)
		} else if timePoint == "1 hour post" {
			// Sustained elevation
			changeFactor = 1.0 + (0.25 * intensityFactor)
		} else {
			// Gradually returns to baseline
			changeFactor = 1.0 + (0.1 * intensityFactor)
		}

	case "Norepinephrine":
		if timePoint == "During" {
			// Strong intensity-dependent increase
			changeFactor = 1.0 + (0.8 * intensityFactor * (0.8 + 0.2*durationFactor))
		} else if timePoint == "Immediate post" {
			// Begins declining but still elevated
			changeFactor = 1.0 + (0.5 * intensityFactor)
		} else if timePoint == "1 hour post" {
			// Returning to baseline
			changeFactor = 1.0 + (0.2 * intensityFactor)
		} else {
			// Normalized after 3+ hours
			changeFactor = 1.0 + (0.05 * intensityFactor)
		}

	case "Gamma-Aminobutyric Acid (GABA)":
		if timePoint == "During" {
			// Modest increase during exercise
			changeFactor = 1.0 + (0.2 * intensityFactor)
		} else if timePoint == "Immediate post" {
			// More pronounced post-exercise
			changeFactor = 1.0 + (0.3 * intensityFactor)
		} else if timePoint == "1 hour post" {
			// Sustained elevation
			changeFactor = 1.0 + (0.25 * intensityFactor)
		} else {
			// Gradual return to baseline
			changeFactor = 1.0 + (0.1 * intensityFactor)
		}

	case "Glutamate":
		if timePoint == "During" {
			// Intensity-dependent increase
			changeFactor = 1.0 + (0.3 * intensityFactor)
		} else if timePoint == "Immediate post" {
			// Quick normalization
			changeFactor = 1.0 + (0.15 * intensityFactor)
		} else {
			// Returns to baseline
			changeFactor = 1.0
		}

	case "Endocannabinoids (primarily Anandamide)":
		if intensityPercent >= 70 && intensityPercent <= 85 {
			// Greatest increase at moderate-to-high intensity
			intensityBoost := 1.0
		} else if intensityPercent > 85 {
			// Less pronounced at very high intensities
			intensityBoost := 0.8
		} else {
			// Lower response at lower intensities
			intensityBoost := 0.6
		}

		if timePoint == "During" {
			// Peaks during exercise
			changeFactor = 1.0 + (0.6 * intensityFactor * intensityBoost)
		} else if timePoint == "Immediate post" {
			// Remains elevated immediately after
			changeFactor = 1.0 + (0.5 * intensityFactor * intensityBoost)
		} else if timePoint == "1 hour post" {
			// Begins to decline
			changeFactor = 1.0 + (0.2 * intensityFactor)
		} else {
			// Returns to baseline
			changeFactor = 1.0
		}

	case "Acetylcholine":
		if timePoint == "During" {
			// Strong increase during exercise, especially for neuromuscular function
			changeFactor = 1.0 + (0.4 * intensityFactor * (0.8 + 0.2*durationFactor))
		} else if timePoint == "Immediate post" {
			// Begins returning to baseline
			changeFactor = 1.0 + (0.2 * intensityFactor)
		} else {
			// Rapid return to baseline
			changeFactor = 1.0 + (0.05 * intensityFactor)
		}
	}

	return changeFactor
}

// GetNeurotransmitterByCondition returns relevant neurotransmitters for specific conditions
func GetNeurotransmitterByCondition(condition string) []Neurotransmitter {
	var relevantNTs []Neurotransmitter

	switch condition {
	case "Depression":
		relevantNTs = append(relevantNTs, Serotonin, Dopamine, Norepinephrine, Endocannabinoids)
	case "Anxiety":
		relevantNTs = append(relevantNTs, GABA, Serotonin, Endocannabinoids)
	case "ADHD":
		relevantNTs = append(relevantNTs, Dopamine, Norepinephrine, Acetylcholine)
	case "Parkinson's disease":
		relevantNTs = append(relevantNTs, Dopamine, Acetylcholine, Glutamate)
	case "Alzheimer's disease":
		relevantNTs = append(relevantNTs, Acetylcholine, Glutamate)
	case "Chronic pain":
		relevantNTs = append(relevantNTs, Endocannabinoids, GABA, Serotonin)
	default:
		relevantNTs = GetExerciseResponsiveNeurotransmitters()
	}

	return relevantNTs
}
