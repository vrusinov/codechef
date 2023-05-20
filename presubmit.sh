#!/bin/bash

# SPDX-FileCopyrightText: 2022 Google Inc
# SPDX-FileCopyrightText: 2022 Vladimir Rusinov <vrusinov@google.com>
#
# SPDX-License-Identifier: Apache-2.0

set -e

pytype .

reuse lint
