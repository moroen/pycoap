#!/usr/bin/env python3

import pycoap

import argparse

parser = argparse.ArgumentParser()
parser.add_argument("uri")
parser.add_argument("payload", nargs="?")
parser.add_argument("--ident")
parser.add_argument("--key")
parser.add_argument("--method", default=pycoap.PUT)

class MalformedURI(Exception):
    pass

class MissingCredentials(Exception):
    pass

def select_method(method):
    method = method.upper()

    methods = {
        "GET":  pycoap.GET,
        "PUT":  pycoap.PUT,
        "POST": pycoap.POST
    }

    return methods.get(method, pycoap.GET)

if __name__ == '__main__':    
    result = None
    args = parser.parse_args()

    fullUri = args.uri
    
    uri = fullUri.split("/")

    method = select_method(args.method)
    print(method)


    try:
        if not (uri[0] == "coap:" or uri[0] == "coaps:"):
            raise (MalformedURI())
        if not uri[1]=="":
            raise (MalformedURI("Missing //"))
        if not uri[2].find(":")>0:
            raise (MalformedURI("Missing port"))

        dest = "/".join(uri[3:])
        
        if uri[0]=="coap:":
            if args.payload != None:
                print(args.payload)
                result = pycoap.PutRequest(uri[2], dest, args.payload)
            else:
                result = pycoap.Request(uri[2], dest)

        if uri[0]=="coaps:":
            if args.ident==None or args.key==None:
                raise MissingCredentials
            if args.payload != None:
                result = pycoap.Request(fullUri, payload=args.payload, method=method, ident=args.ident, key=args.key)
            else:
                result = pycoap.Request(fullUri, method=pycoap.GET, ident=args.ident, key=args.key)

        if result!=None:
            print("Response: {}".format(result))
    


    except MissingCredentials:
        print("Error: Missing credentials for DTLS-connection!")
    except MalformedURI:
        print("Error: Malformed uri {0}".format(args.uri))
