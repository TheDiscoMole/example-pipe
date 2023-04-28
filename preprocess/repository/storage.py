from google.cloud import storage

client = storage.Client()
bucket = client.bucket("pipeline")

# open a file pointer for a google storage blob
def open (f, p):
    return bucket.blob(f).open(p)
