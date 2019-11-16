package main

// #cgo pkg-config: python3
// #define Py_LIMITED_API
// #include <Python.h>
// int PyArg_ParseTuple_LL(PyObject *, long long *, long long *);
// int PyArg_ParseTuple_S(PyObject *, char *);
// char * ParseStringArgument(PyObject *);
import (
	"C"
	"log"
	"strconv"
	"strings"

	coap "github.com/moroen/gocoap"
)

// Python Functions

//export coapRequest
func coapRequest(gateway, uri *C.char) *C.char {
	/*
		param := coap.RequestParams{Host: }


		msg, err := GetRequest(C.GoString(gateway), C.GoString(uri))
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		// return C.PyUnicode_FromString(C.CString(msg.String()))
		return C.CString(msg.String())
	*/
	return nil
}

//export coapRequestDTLS
func coapRequestDTLS(gateway, uri, ident, key *C.char) *C.char {

	gw := strings.Split(C.GoString(gateway), ":")
	port, err := strconv.Atoi(gw[1])
	if err != nil {
		panic(err.Error())
	}

	params := coap.RequestParams{Host: gw[0], Port: port, Uri: C.GoString(uri), Id: C.GoString(ident), Key: C.GoString(key)}

	res, err := coap.GetRequest(params)
	if err != nil {
		// log.Println(err.Error())
		return nil
	}

	return C.CString(string(res))
}

//export coapPutRequest
func coapPutRequest(gateway, uri, payload *C.char) *C.char {
	/*
		msg, err := PutRequest(C.GoString(gateway), C.GoString(uri), C.GoString(payload))
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		return C.CString(msg.String())
	*/
	return nil
}

//export coapPutRequestDTLS
func coapPutRequestDTLS(gateway, uri, ident, key, payload *C.char) *C.char {
	gw := strings.Split(C.GoString(gateway), ":")
	port, err := strconv.Atoi(gw[1])
	if err != nil {
		panic(err.Error())
	}

	params := coap.RequestParams{Host: gw[0], Port: port, Uri: C.GoString(uri), Id: C.GoString(ident), Key: C.GoString(key), Payload: C.GoString(payload)}

	res, err := coap.PutRequest(params)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return C.CString(string(res))
}

//export coapPostRequestDTLS
func coapPostRequestDTLS(gateway, uri, ident, key, payload *C.char) *C.char {
	gw := strings.Split(C.GoString(gateway), ":")
	port, err := strconv.Atoi(gw[1])
	if err != nil {
		panic(err.Error())
	}

	params := coap.RequestParams{Host: gw[0], Port: port, Uri: C.GoString(uri), Id: C.GoString(ident), Key: C.GoString(key), Payload: C.GoString(payload)}

	res, err := coap.PostRequest(params)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return C.CString(string(res))
}

func main() {}
