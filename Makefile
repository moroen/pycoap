gocoap = vendor/src/github.com/moroen/gocoap

cflags = $(shell pkg-config --cflags python3)
ldflags = -Wl,$(shell pkg-config --libs python3-embed)
hdir = $(shell pwd)

target = _py3coap.cpython-38-x86_64-linux-gnu.so


# ldflags = -Wl,-undefined,$(shell pkg-config --libs python3-embed)

dummy:



$(gocoap):
	go get github.com/moroen/gocoap

lib: $(gocoap)
	CGO_CFLAGS=$(cflags) CGO_LDFLAGS=$(ldflags) go build -buildmode=c-shared -o $(hdir)/$(target) ./src/py3coap/

module:
	python3 setup.py build --user

develop: 
	python3 setup.py develop --user

clean:
	-rm -rf dist
	-rm -rf build
	-rm -rf *egg-info
	-rm *.so
	-rm -rf pkg
	-rm -rf bin