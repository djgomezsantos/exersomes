from transformers import GPT4LMHeadModel, GPT2Tokenizer  # Change import to the right one for GPT-4
import torch

class TransformerRNASequenceModel:
    def __init__(self, model_name='gpt-4'):
        self.tokenizer = GPT2Tokenizer.from_pretrained(model_name)  # Ensure this corresponds to the GPT-4 tokenizer
        self.model = GPT4LMHeadModel.from_pretrained(model_name)    # Load the GPT-4 model

    def generate_sequence(self, input_sequence, max_length=50, method='top_p', top_k=50, top_p=0.95):
        inputs = self.tokenizer.encode(input_sequence, return_tensors='pt')
        outputs = self.model.generate(
            inputs,
            max_length=max_length,
            do_sample=True,
            top_k=top_k if method == 'top_k' else 0,
            top_p=top_p if method == 'top_p' else 1,
            num_return_sequences=1
        )
        return self.tokenizer.decode(outputs[0], skip_special_tokens=True)

# Example input tensor with padding (padded to max length)
input_ids = torch.tensor([/* your padded sequence data */])

# Create an attention mask (1 for actual tokens, 0 for padding)
attention_mask = (input_ids != PAD_TOKEN_ID).type(torch.float)

# Forward pass
outputs = model(input_ids=input_ids, attention_mask=attention_mask)

# Example usage:
model = TransformerRNASequenceModel()
generated_sequence = model.generate_sequence("AUGCUACG", max_length=4000, min_length=20)
print(generated_sequence)