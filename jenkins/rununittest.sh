#!/bin/bash

set -x
export http_proxy=http://192.168.100.116:1080;
export https_proxy=http://192.168.100.116:1080;
go get -d -v golang.org/x/net/html
go get -u github.com/jstemmer/go-junit-report
go test -v 2>&1 > tmp
status=$?
$GOPATH/bin/go-junit-report < tmp > test_output.xml

exit ${status}