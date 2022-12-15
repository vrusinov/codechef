#!/usr/bin/python3
# SPDX-FileCopyrightText: 2022 Google Inc
#
# SPDX-License-Identifier: Apache-2.0
"""https://www.codechef.com/problems/DIVISIBLEBY"""

import math
from typing import Sequence

def div_by_a_i(a: Sequence[int]) -> Sequence[int]:
    """Calculate GCD of the whole array, and Yi will be Ai / gcd."""
    gcd = a[0]
    for i in a:
        gcd = math.gcd(gcd, i)
    result = [i // gcd for i in a]
    return result

if __name__ == '__main__':
    t = int(input())
    for _ in range(t):
        unused_n = input()
        a = [int(x) for x in input().split()]
        result = div_by_a_i(a)
        print(' '.join(str(x) for x in result))