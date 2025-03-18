#!/bin/bash
RUN_NAME=hertz_demo
mkdir -p output
cp script/bootstrap.sh output 2>/dev/null
chmod +x output/bootstrap.sh
make update_api
go build -o output/${RUN_NAME}