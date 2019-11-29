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
	"unicode/utf8"
)

var _debugLevel int

// Python Functions

//export coapDebugLevel
func coapDebugLevel(level C.int) C.int {

	_debugLevel = int(level)
	return level
}

//export coapRequest
func coapRequest(gateway, uri *C.char) *C.char {
	gw := strings.Split(C.GoString(gateway), ":")
	port, err := strconv.Atoi(gw[1])
	if err != nil {
		panic(err.Error())
	}

	params := coap.RequestParams{Host: gw[0], Port: port, Uri: C.GoString(uri)}
	response, err := coap.GetRequest(params)
	if err != nil {
		if _debugLevel == 1 {
			log.Println(err.Error())
		}
		return nil
	}

	return C.CString(validateResponse(response))
}

//export coapRequestDTLS
func coapRequestDTLS(gateway, uri, ident, key *C.char) *C.char {

	gw := strings.Split(C.GoString(gateway), ":")
	port, err := strconv.Atoi(gw[1])
	if err != nil {
		panic(err.Error())
	}

	params := coap.RequestParams{Host: gw[0], Port: port, Uri: C.GoString(uri), Id: C.GoString(ident), Key: C.GoString(key)}

	response, err := coap.GetRequest(params)
	if err != nil {
		if _debugLevel == 1 {
			log.Println(err.Error())
		}
		return nil
	}

	return C.CString(validateResponse(response))
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

	response, err := coap.PutRequest(params)
	if err != nil {
		if _debugLevel == 1 {
			log.Println(err.Error())
		}
		return nil
	}

	return C.CString(validateResponse(response))
}

//export coapPostRequestDTLS
func coapPostRequestDTLS(gateway, uri, ident, key, payload *C.char) *C.char {
	gw := strings.Split(C.GoString(gateway), ":")
	port, err := strconv.Atoi(gw[1])
	if err != nil {
		panic(err.Error())
	}

	params := coap.RequestParams{Host: gw[0], Port: port, Uri: C.GoString(uri), Id: C.GoString(ident), Key: C.GoString(key), Payload: C.GoString(payload)}

	response, err := coap.PostRequest(params)
	if err != nil {
		if _debugLevel == 1 {
			log.Println(err.Error())
		}
		return nil
	}

	return C.CString(validateResponse(response))
}

func validateResponse(response []byte) (string) {
	res := string(response)

	if !utf8.ValidString(res) {
        v := make([]rune, 0, len(res))
        for i, r := range res {
            if r == utf8.RuneError {
                _, size := utf8.DecodeRuneInString(res[i:])
                if size == 1 {
                    continue
                }
            }
            v = append(v, r)
        }
        res = string(v)
    } else {
		res = string(res)
	}

	return res
}

func main() {}
