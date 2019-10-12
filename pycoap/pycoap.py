import _pycoap



def Request(uri, payload=None, ident=None, key=None):
    if (payload is None) and (ident is None):
        print("Coap-request")
        return
    
    if (payload is None) and (ident is not None):
        
        
        uri = uri.split("/")
        
        host = uri[2]
        
        dest = "/".join(uri[3:])
        
        result = _pycoap.DTLSRequest(host, dest, ident, key)
        return result

    if (payload is not None) and (ident is not None):
        uri = uri.split("/")
        host = uri[2]
        dest = "/".join(uri[3:])

        result = _pycoap.DTLSPutRequest(host, dest, payload, ident, key)
        return result
    