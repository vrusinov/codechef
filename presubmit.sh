#!/bin/bash

set -e

# Run go linters in all go directories
for d in $(find . -name \*.go | sed -r 's|/[^/]+$||' |sort -u) ; do
    pushd $d
    golangci-lint run
    popd    
done

pytype .

# Install presubmit
ln -sf ../../presubmit.sh .git/hooks/pre-commit
