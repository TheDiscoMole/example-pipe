import sentencepiece

from ..repository import storage

# embed text file
async def preprocess (filename):
    # load text embedding model
    spp = sentencepiece.SentencePieceProcessor(model_file='service/model/bpemb')

    async with storage.open(f'ingest/{filename}', 'r') as file:
        lines = file.readlines()

    # embed lines of text
    lines = [spp.encode(line) for line in lines]

    async with storage.open(f'preprocess/{filename}', 'w') as file:
        file.writelines(lines)
