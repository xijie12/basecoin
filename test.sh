#!/bin/bash
go build *.go
rm -rf blockChain.db
rm -rf vlockChain.db.lock

./block.exe
