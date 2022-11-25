#!/bin/bash

# SPDX-FileCopyrightText: 2022 Google Inc
# SPDX-FileCopyrightText: 2022 Vladimir Rusinov <vrusinov@google.com>
#
# SPDX-License-Identifier: Apache-2.0

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
