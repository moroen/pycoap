[![Build Status](https://dev.azure.com/moroen/PyCOAP/_apis/build/status/pycoap?branchName=master)](https://dev.azure.com/moroen/PyCOAP/_build/latest?definitionId=1&branchName=master)

# PyCOAP

A python extension module written in GOLANG to provide basic COAP-capabilities, including support for DTLS.


## Requirements
To  compile and install PyCOAP you need:

### Python-development libraries, upgraded pip and setuptools

```shell
  $ sudo apt install python3 python3-dev python3-pip
  $ sudo pip3 install -U pip setuptools
```

### golang-compiler

```shell
  $ sudo apt install golang
```

## Installation

```shell
  $ git clone https://github.com/moroen/pycoap.git
  $ cd pycoap
  $ sudo python3 setup.py install
```

## Usage
See included pyclient.py for a small example
