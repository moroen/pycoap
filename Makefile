dummy:

module:
	python3 setup.py build

develop: 
	python3 setup.py develop

clean:
	rm -rf build
	rm -rf *egg-info
	rm *.so
	rm pkg/*