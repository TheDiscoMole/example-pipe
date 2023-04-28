from PIL import Image

from ..repository import storage

# resize image file
async def preprocess (filename):
    async with storage.open(f'ingest/{filename}', 'r') as file:
        image = Image.open(file)

    # get dimensions
    width, height = image.size

    # shrink smallest dimension to 256 and other dimension with respect to aspect ratio
    if width > height: width, height = int(width * 256 / height), 256
    else: width, height = 256, int(height * 256 / width)

    # resize and format image
    image = image.resize((width, height))
    image = image.convert("RGB")

    async with storage.open(f'preprocess/{filename}', 'w') as file:
        image.save(file)
