#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=hertz_demo
echo "$CURDIR/${BinaryName}"
exec $CURDIR/${BinaryName} serve