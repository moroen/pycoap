import _py3coap
from .errors import MalformedUri, MissingCredentials

GET = 0
PUT = 1
POST = 2


def setDebugLevel(level):
    _py3coap.DebugLevel(level)

def CloseConnection():
    _py3coap.DTLSCloseConnection()

def Request(uri, payload=None, method=GET, ident=None, key=None):

    protocol = uri.split("//")

    if not (protocol[0] == "coap:" or protocol[0] == "coaps:"):
        raise MalformedUri("Missing protocol")

    uri = protocol[1].split("/")

    if len(uri) < 2:
        raise MalformedUri(uri)

    host = uri[0]

    if not ":" in host:
        raise MalformedUri("Missing port")

    dest = "/".join(uri[1:])

    if protocol[0] == "coap:":
        if payload is None:
            return _py3coap.Request(host, dest)
        else:
            method = PUT if method == GET else method
            if method == PUT:
                return _py3coap.PutRequest(host, dest, payload)
            elif method == POST:
                return _py3coap.PostRequest(host, dest, payload)

    elif protocol[0] == "coaps:":
        if ident is None or key is None:
            raise MissingCredentials

        if payload is None:
            return _py3coap.DTLSRequest(host, dest, ident, key)
        else:
            method = PUT if method == GET else method

            if method == PUT:
                result = _py3coap.DTLSPutRequest(host, dest, payload, ident, key)
                return result
            elif method == POST:
                result = _py3coap.DTLSPostRequest(host, dest, payload, ident, key)
                return result

    else:
        raise MalformedUri("Unknow protocol: {}".format(protocol[0]))

    return None
