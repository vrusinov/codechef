#!/usr/bin/python3

# SPDX-FileCopyrightText: 2022 Google Inc
# SPDX-FileCopyrightText: 2022 Vladimir Rusinov <vrusinov@google.com>
#
# SPDX-License-Identifier: Apache-2.0

"""ATM Machine

https://www.codechef.com/problems/ATM2
"""

from typing import Any, Sequence


def atm(n: Any, k: int, a: Sequence[int]):
    result = []
    del n  # unused
    for withdrawal in a:
        if k >= withdrawal:
            k -= withdrawal
            result.append(1)
        else:
            result.append(0)
    return "".join(str(x) for x in result)


def main():
    t = int(input())
    for _ in range(t):
        line = input().split()
        n = int(line[0])
        k = int(line[1])
        a = [int(x) for x in input().split()]
        print(atm(n, k, a))


if __name__ == "__main__":
    main()
