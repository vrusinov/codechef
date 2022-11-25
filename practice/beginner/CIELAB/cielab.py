#!/usr/bin/python3

# SPDX-FileCopyrightText: 2022 Google Inc
# SPDX-FileCopyrightText: 2022 Vladimir Rusinov <vrusinov@google.com>
#
# SPDX-License-Identifier: Apache-2.0

def ciel(a: int, b: int) -> int:
    correct = a - b
    if (correct % 10) == 0:
        return correct + 1
    elif correct == 1:
        return 2
    else:
        return correct - 1


def main():
    a, b = [int(x) for x in input().split()]
    print(ciel(a, b))

if __name__ == '__main__':
    main()
