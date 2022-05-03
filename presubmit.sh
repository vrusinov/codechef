#!/bin/bash

set -e

# Run go linters in all go directories
for d in $(find . -name \*.go | sed -r 's|/[^/]+$||' |sort -u) ; do
    golint $d
    which gometalinter && gometalinter $d
done

pytype .
