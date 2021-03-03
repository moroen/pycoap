#!/bin/bash
set -e -x

# Collect the pythons
pys=(/opt/python/*/bin)

# Filter out Python2
pys=(${pys[@]//*2*/})

# Filter out Python 3.4
pys=(${pys[@]//*34*/})

# Filter out Python 3.5
pys=(${pys[@]//*35*/})

# Install go
yum -y install wget
wget https://golang.org/dl/go1.16.linux-amd64.tar.gz
tar -xzf go1.16.linux-amd64.tar.gz

export PATH=$PATH:`pwd`/go/bin
export GOPATH=/io/
export GO111MODULE=off

go get -d -v github.com/moroen/gocoaplib
go build -buildmode=c-archive -o /io/lib/static/libgocoap.a /io/src/github.com/moroen/gocoaplib/main.go

# Compile wheels
for PYBIN in "${pys[@]}"; do
    "${PYBIN}/pip" install -r /io/azure/dev-requirements.txt
    "${PYBIN}/pip" wheel /io/ -w wheelhouse/
done

# Bundle external shared libraries into the wheels
for whl in wheelhouse/$package_name-*.whl; do
    auditwheel repair --plat $PLAT "$whl" -w /io/wheelhouse/
done

# Remove 2010-wheels
rm /io/wheelhouse/*2010*