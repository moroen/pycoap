from setuptools import Extension
from setuptools import setup


setup(
    name="py3coap",
    description="A GOLANG python extension for COAP/COAPS-requests",
    url="https://github.com/moroen/pycoap",
    version="0.9.0",
    author="moroen",
    author_email="no@email.com",
    classifiers=[
        "License :: OSI Approved :: MIT License",
        "Programming Language :: Python :: 3.6",
        "Programming Language :: Python :: Implementation :: CPython",
    ],
    ext_modules=[Extension("_py3coap", ["src/py3coap/py3coap.go"])],
    data_files=[('src/py3coap', ['src/py3coap/py3coap.h', 'src/py3coap/py3coap.c'])],
    packages=["py3coap"],
    build_golang={"root": "github.com/moroen/py3coap"},
    setup_requires=["setuptools-golang"],
    scripts=["pycoap-client"],
)
