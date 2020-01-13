#include <Python.h>

#ifndef _PYCOAP_H
#define _PYCOAP_H

#define error_urinotfound 1
#define error_handshake 2
#define error_notallowed 3
#define error_readerror 4
#define error_writeerror 5

struct pycoap_result
{
    char * result;
    int error;
};

typedef struct pycoap_result coapresult;

// Coap Functions
int coapDebugLevel(int);
coapresult coapRequest(char *, char*);
coapresult coapPutRequest(char *, char *, char *);
coapresult coapRequestDTLS(char *, char*, char *, char *);
coapresult coapPutRequestDTLS(char *, char *, char *, char *, char *);
coapresult coapPostRequestDTLS(char *, char *, char *, char *, char *);


// int raiseError(int);
void raiseError(int);

static PyObject *UriNotFoundError = NULL;
static PyObject *HandshakeError = NULL;
static PyObject *ReadTimeoutError = NULL;
static PyObject *WriteTimeoutError = NULL;
static PyObject *MethodNotAllowedError = NULL;

#endif

