import torch

from google.cloud import storage

def save (model, filename):
    client = storage.Client()
    bucket = client.bucket("pipeline")
    blob = bucket.blob(filename)

    with blob.open('wb', ignore_flush=True) as file:
        torch.save(model, file)

def load (filename):
    client = storage.Client()
    bucket = client.bucket("pipeline")
    blob = bucket.blob(filename)

    with blob.open('r') as file:
         return torch.load(file)
