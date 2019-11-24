import _pycoap

GET = 0
PUT = 1
POST = 2


def setDebugLevel(level):
    _pycoap.DebugLevel(level)


def Request(uri, payload=None, method=GET, ident=None, key=None):
    uri = uri.split("/")
    host = uri[2]
    dest = "/".join(uri[3:])

    if (payload is None) and (ident is None):
        return _pycoap.Request(host, dest)

    if (payload is None) and (ident is not None):
        return _pycoap.DTLSRequest(host, dest, ident, key)

    if (payload is not None) and (ident is not None):
        if method == GET:
            method = PUT

        if method == PUT:
            result = _pycoap.DTLSPutRequest(host, dest, payload, ident, key)
            return result
        elif method == POST:
            result = _pycoap.DTLSPostRequest(host, dest, payload, ident, key)
            return result

        return None
