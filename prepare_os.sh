#!/usr/bin/bash

files=`ls codegen/*.py`
#echo $files
for each in ${files}
do
    #echo $each
    python2.7 -m py_compile $each
done
rm -rf go-asn1/*
mkdir -p go-asn1/codegen
cp -f codegen/*.pyc go-asn1/codegen/
cp -arf asn1 go-asn1/
cp -arf libgo go-asn1/
cp README.md asn_errors.py  asn_parser.py go-asn1/
cp genasnpy.py go-asn1/go-asn1c
