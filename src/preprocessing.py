import pandas as pd

def pad_sequences(sequences, max_length=4000, pad_token='PAD'):
    padded_sequences = []
    for seq in sequences:
        # Pad sequences that are less than max_length
        if len(seq) < max_length:
            padded_seq = seq + [pad_token] * (max_length - len(seq))  # Apply padding
        else:
            padded_seq = seq[:max_length]  # Truncate sequences longer than max_length

        padded_sequences.append(padded_seq)
    return padded_sequences

def load_and_preprocess_data(file_path):
    # Load your RNA data (as an example)
    data = pd.read_csv(file_path)

    # Example: Assuming data has a column 'RNA_sequence' that contains the sequences
    sequences = data['RNA_sequence'].tolist()
    
    # Convert sequences to a suitable format (e.g., list of lists if necessary)
    sequences = [list(seq) for seq in sequences]  # Convert strings to lists of characters/nucleotides

    # Apply padding
    padded_sequences = pad_sequences(sequences)

    # Proceed with other preprocessing steps as needed
    # ...

    return padded_sequences

# Example usage
if __name__ == '__main__':
    file_path = 'path/to/your/data.csv'  # Replace with actual data path
    padded_sequences = load_and_preprocess_data(file_path)
    print(padded_sequences)
