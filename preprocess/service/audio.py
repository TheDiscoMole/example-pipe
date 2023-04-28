import librosa
import numpy as np

from ..repository import storage

# mel-transform audio file
async def preprocess (filename):
    async with storage.open(f'ingest/{filename}', 'r') as file:
        y, sr = librosa.load(file)

    # compute mel-transform
    mel = librosa.feature.melspectrogram(y=y, sr=sr, n_fft=2048, hop_length=512, n_mels=128)
    mel = librosa.power_to_db(mel, ref=np.max)

    async with storage.open(f'preprocess/{filename}', 'w') as file:
        np.save(file, mel)
