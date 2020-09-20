dummy:

lib:
	CGO_CFLAGS='-I/usr/local/include -I/usr/local/opt/openssl@1.1/include -I/usr/local/opt/sqlite/include -I/usr/local/Cellar/python@3.8/3.8.5/Frameworks/Python.framework/Versions/3.8/include/python3.8' CGO_LDFLAGS=-Wl,-undefined,dynamic_lookup go build -buildmode=c-shared -o /Users/mak/projects/python/pycoap/_py3coap.cpython-38-darwin.so ./src/py3coap/

module:
	python3 setup.py build

develop: 
	python3 setup.py develop

clean:
	-rm -rf dist
	-rm -rf build
	-rm -rf *egg-info
	-rm *.so
	-rm -rf pkg
	-rm -rf bin