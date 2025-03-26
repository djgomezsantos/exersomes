from transformers import GPT4LMHeadModel, GPT2Tokenizer  # Ensure these imports are correct for GPT-4
import torch
import numpy as np

class TransformerRNASequenceModel:
    def __init__(self, model_name='gpt-4'):
        # Initialize the tokenizer and model for GPT-4
        self.tokenizer = GPT2Tokenizer.from_pretrained(model_name)  # Make sure this is correct for the GPT-4 tokenizer
        self.model = GPT4LMHeadModel.from_pretrained(model_name)    # Load the GPT-4 model

    def pad_sequences(self, sequences, max_length=1000, pad_token_id=0):
        """Pad sequences to the maximum length."""
        padded_sequences = []
        for seq in sequences:
            # Pad sequences shorter than max_length
            if len(seq) < max_length:
                padded_seq = seq + [pad_token_id] * (max_length - len(seq))  # Padding with pad_token_id
            else:
                padded_seq = seq[:max_length]  # Truncate longer sequences
            padded_sequences.append(padded_seq)
        return np.array(padded_sequences)

    def create_attention_mask(self, padded_sequences, pad_token_id=0):
        """Create an attention mask where 1 indicates real tokens and 0 indicates padding."""
        return (padded_sequences != pad_token_id).astype(float)

    def generate_sequence(self, input_sequence, max_length=1000, min_length=20, method='top_p', top_k=50, top_p=0.95):
        # Encode the input sequence
        inputs = self.tokenizer.encode(input_sequence, return_tensors='pt')

        # Pad the input sequences
        padded_inputs = self.pad_sequences([inputs.numpy()], max_length)  # Convert tensor to numpy array for padding
        attention_mask = self.create_attention_mask(padded_inputs)  # Create the attention mask

        # Convert to PyTorch tensors
        input_ids = torch.tensor(padded_inputs)  # Shape: (1, max_length)
        attention_mask = torch.tensor(attention_mask)  # Shape: (1, max_length)

        outputs = self.model.generate(
            input_ids=input_ids,
            attention_mask=attention_mask,  # Use the attention mask
            max_length=max_length,
            min_length=min_length,
            do_sample=True,
            top_k=top_k if method == 'top_k' else 0,
            top_p=top_p if method == 'top_p' else 1,
            num_return_sequences=1
        )
        
        # Decode the generated output to a readable format
        return self.tokenizer.decode(outputs[0], skip_special_tokens=True)

# Example usage:
model = TransformerRNASequenceModel()

# Assuming 'AUGCUACG' is your input sequence
input_sequence = "AUGCUACG"
generated_sequence = model.generate_sequence(input_sequence, max_length=1000, min_length=20)  # Set max_length to 1000
print(generated_sequence)