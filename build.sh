#!/bin/bash


source ~/.profile

rm -R build
make build

echo build/kvant node
version=`build/kvant version`
echo $version

build/kvant node
filename="kvant_${version}"
echo $filename > bin_name
cp build/kvant release/$filename


