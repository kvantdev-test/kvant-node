#!/bin/bash


source ~/.profile

rm -R build
make build

echo build/kvant node
build/kvant version
build/kvant node

