import argparse  # For command-line argument parsing
from src.preprocessing import load_and_preprocess_data  # Import your data loading and preprocessing code
from src.model import TransformerRNASequenceModel  # Adjust this path based on where your model class is defined

def main(input_sequence):
    # Initialize the model
    model = TransformerRNASequenceModel(model_name='gpt-4')

    # Generate the RNA sequence
    generated_sequence = model.generate_sequence(input_sequence, max_length=4000, min_length=20)
    
    # Print results
    print(f"Input Sequence: {input_sequence}")
    print(f"Generated Sequence: {generated_sequence}")

if __name__ == "__main__":
    # Set up the argument parser
    parser = argparse.ArgumentParser(description="Generate exerkine RNA sequences using a Transformer model")
    parser.add_argument('--input_sequence', type=str, required=True, help='RNA sequence to start generation from')
    
    # Parse arguments
    args = parser.parse_args()
    
    # Run the main function
    main(args.input_sequence)