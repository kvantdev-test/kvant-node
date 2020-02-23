#!/bin/bash


source ~/.profile

rm -R build
make build

echo build/kvant node
version=`build/kvant version`
echo $version

build/kvant node
cp build/kvant release/kvant_${version}


