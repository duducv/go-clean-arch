#!/bin/bash
go test -coverprofile ../test/cover.out ../internal/core/... 

if [ $? -eq 0 ]; then
    go tool cover -html=../test/cover.out -o ../test/cover.html
else
    echo "One or more test failed, the coverage report will not be generated."
fi