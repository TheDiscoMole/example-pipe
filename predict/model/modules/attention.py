import torch

class AttentionEncoder (torch.nn.Module):
    def __init__ (self, model_dim,
        nhead=8,
        activation=torch.nn.ReLU,
        dropout=1e-1,
        layer_norm_eps=1e-5
    ):
        self.dropout = torch.nn.Dropout(dropout)
        self.activation = activation

        self.attention = torch.nn.MultiheadAttention(model_dim, nhead, dropout=dropout)
        self.normalize1 = torch.nn.LayerNorm(attention_dim, eps=layer_norm_eps)
        self.linear1 = torch.nn.Linear(model_dim, model_dim)
        self.normalize2 = torch.nn.LayerNorm(model_dim, eps=layer_norm_eps)

    def forward (self, key, query, value):

        attention = self.attention(key, query, value)
        attention = self.dropout(attention)

        output = self.normalize1(input + attention)
        output = self.normalize2(output + self.linear2(output))

        return self.activation(output)
