from werkzeug.wrappers import Request, Response

# middleware to validate pubsub request format
class pubsub ():
    def __init__ (self, app):
        self.app = app

    def __call__ (self, environ, start_response):
        request = Request(environ)
        message = request.get_json()

        if not message:
            response = Response("received message not json formatted", mimetype= 'text/plain', status=400)
            return response(environ, start_response)

        if not isinstance(message, dict) or "message" not in message:
            response = Response("received message not pub/sub formatted", mimetype= 'text/plain', status=400)
            return response(environ, start_response)

        environ["filename"] = message["data"]
        environ["attributes"] = message["attributes"]

        return self.app(environ, start_response)
