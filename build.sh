#!/bin/bash
## Linux x86 y x86_64
echo "-> Linux"
echo "* x86"
cd $GOPATH/src/github.com/xxxtonixxx/goPassword/build/linux/x86/
env GOARCH=386 GOOS=linux go build ../../../
echo "* x86_64"
cd ../x86_64/
env GOOS=linux go build ../../../

## Windows x86 y x86_64
echo "-> Windows"
echo "* x86"
cd ../../windows/x86
env GOOS=windows go build ../../../
echo "* x86_64"
cd ../x86_64
env GOOS=windows GOARCH=386 go build ../../../
