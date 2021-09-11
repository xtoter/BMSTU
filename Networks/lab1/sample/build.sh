#!/bin/bash
export GOPATH=`pwd`
go install ./src/client
go install ./src/server
