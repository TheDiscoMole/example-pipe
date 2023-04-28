import os

from flask import Flask, request

from middleware import pubsub
from service import audio, text, image

app = Flask(__name__)
app.wsgi_app = pubsub(app.wsgi_app)

# pre-process audio files
@app.route("/audio", methods=["POST"])
async def audio_handler ():
    try:
        filename = pubsub.parse(request.environ["filename"])
        await audio.preprocess(filename)
    except Exception as e:
        abort(400, e)

    return 200, ""

# pre-process text files
@app.route("/text", methods=["POST"])
async def text_handler ():
    try:
        filename = pubsub.parse(request.environ["filename"])
        await text.preprocess(filename)
    except Exception as e:
        abort(400, e)

    return 200, ""

# pre-process image files
@app.route("/image", methods=["POST"])
async def image_handler ():
    try:
        filename = pubsub.parse(request.environ["filename"])
        await image.preprocess(filename)
    except Exception as e:
        abort(400, e)

    return 200, ""

if __name__ == "__main__":
    # local server environment
    PORT = int(os.getenv("PORT")) if os.getenv("PORT") else 8080
    app.run(host="127.0.0.1", port=PORT, debug=True)
