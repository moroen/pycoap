from setuptools import Extension
from setuptools import setup


setup(
    name="pycoap",
    description="A GOLANG python extension for COAP/COAPS-requests",
    url="https://github.com/moroen/pycoap",
    version="0.6.2",
    author="moroen",
    author_email="no@email.com",
    classifiers=[
        "License :: OSI Approved :: MIT License",
        "Programming Language :: Python :: 3.6",
        "Programming Language :: Python :: Implementation :: CPython",
    ],
    ext_modules=[Extension("_pycoap", ["src/pycoap/pycoap.go"])],
    packages=["pycoap"],
    build_golang={"root": "github.com/moroen/pycoap"},
    setup_requires=["setuptools-golang"],
)
