#define Py_LIMITED_API
#include <Python.h>
#include <stdio.h>

PyObject * sum(PyObject *, PyObject *);
// PyObject * coapRequest(PyObject *, PyObject *);
char * coapRequest(char *, char*);
char * coapPutRequest(char *, char *, char *);
char * coapRequestDTLS(char *, char*, char *, char *);
char * coapPutRequestDTLS(char *, char *, char *, char *, char *);


// Workaround missing variadic function support
// https://github.com/golang/go/issues/975
int PyArg_ParseTuple_LL(PyObject * args, long long * a, long long * b) {  
    return PyArg_ParseTuple(args, "LL", a, b);
}

int PyArg_ParseTuple_SSS(PyObject * args, char * a, char *b, char *c) {
    return PyArg_ParseTuple(args, "sss", a,b,c);
}

const char * ParseStringArgument(PyObject * args) {
    const char* s;

    if (!PyArg_ParseTuple(args, "s", &s))
        return NULL;

    return s;
}

PyObject * request(PyObject *self, PyObject *args) {
    char *gateway, *uri, *res;

    if (!PyArg_ParseTuple(args, "ss", &gateway, &uri))
        Py_RETURN_NONE;
    
    res = coapRequest(gateway, uri);
    if (!res)
        Py_RETURN_NONE;

    return PyUnicode_FromString(res);
}

PyObject * putRequest(PyObject *self, PyObject *args) {
    char *gateway, *uri, *payload, *res;

    if (!PyArg_ParseTuple(args, "sss", &gateway, &uri, &payload))
        Py_RETURN_NONE;

    res = coapPutRequest(gateway, uri, payload);
    if (!res)
        Py_RETURN_NONE;

    return PyUnicode_FromString(res);
    
}

PyObject * requestDTLS(PyObject *self, PyObject *args) {
    char *gateway, *uri, *ident, *key, *res;

    if (!PyArg_ParseTuple(args, "ssss", &gateway, &uri, &ident, &key))
        Py_RETURN_NONE;
    
    res = coapRequestDTLS(gateway, uri, ident, key);
    if (!res)
        Py_RETURN_NONE;

    return PyUnicode_FromString(res);
}

PyObject * putRequestDTLS(PyObject *self, PyObject *args) {
    char *gateway, *uri, *ident, *key, *payload, *res;

    if (!PyArg_ParseTuple(args, "sssss", &gateway, &uri, &payload, &ident, &key))
        Py_RETURN_NONE;

    res = coapPutRequestDTLS(gateway, uri, ident, key, payload);
    if (!res)
        Py_RETURN_NONE;

    return PyUnicode_FromString(res);
    
}

static PyMethodDef CoapMethods[] = {
    {"Request", request, METH_VARARGS, "Make a COAP Request."},
    {"PutRequest", putRequest, METH_VARARGS, "Make a COAP Put Request."},
    {"DTLSRequest", requestDTLS, METH_VARARGS, "Make a COAP Request."},
    {"DTLSPutRequest", putRequestDTLS, METH_VARARGS, "Make a COAP Put Request."},
    {NULL, NULL, 0, NULL}
};

static struct PyModuleDef coapmodule = {  
   PyModuleDef_HEAD_INIT, "coap", NULL, -1, CoapMethods
};

PyMODINIT_FUNC  
PyInit__pycoap(void)  
{
    return PyModule_Create(&coapmodule);
}

