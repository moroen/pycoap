export GOPATH=$(shell pwd)
export GO111MODULE=off

cflags = $(shell pkg-config --cflags python3-embed)
ldflags = -Wl,$(shell pkg-config --libs python3-embed)
hdir = $(shell pwd)

target = _py3coap$(shell python3-config --extension-suffix)

gocoaplib = src/github.com/moroen/gocoaplib


static_lib = lib/static/libgocoap.a
shared_lib = lib/shared/libgocoap.so
lib_source = src/github.com/moroen/gocoaplib/gocoaplib.go

dummy:
	echo $(target)


libs: $(static_lib) $(shared_lib)

$(gocoaplib):
	go get -v github.com/moroen/gocoaplib

$(static_lib): $(gocoaplib)
	go build -buildmode=c-archive -o $(static_lib) $(lib_source)

$(shared_lib): $(gocoaplib)
	go build -buildmode=c-shared -o $(shared_lib) $(lib_source)


# $(gocoap):
#	go get github.com/moroen/gocoap

pythonlib: $(static_lib)
	gcc src/py3coap/py3coap.c -shared -fPIC -Ilib/static -Llib/static $(cflags) $(ldflags) -lgocoap $(python-config --cflags --ldflags --embed) -o $(target)

module: $(static_lib)
	python3 setup.py install --user

develop: $(static_lib)
	python3 setup.py develop --user

clean:
	-rm -rf dist
	-rm -rf build
	-rm -rf *egg-info
	-rm *.so
	-rm -rf bin
	-rm -rf lib
	-rm -rf src/github.com