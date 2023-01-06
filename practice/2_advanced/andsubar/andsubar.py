# SPDX-FileCopyrightText: 2022 Google Inc
# SPDX-FileCopyrightText: 2022 Vladimir Rusinov <vrusinov@google.com>
#
# SPDX-License-Identifier: Apache-2.0

"""https://www.codechef.com/submit/ANDSUBAR"""

import math


def and_subarr(n):
    if n == 1:
        return 1
    # Need highest power of two and all following digits.
    highest_p2 = 2 ** int(math.log(n, 2))
    remainder = n - highest_p2
    # Here we either go from highest power of 2 to the end of array, or take the
    # length of the sequence following previous highest power of two.
    return max(remainder + 1, highest_p2 // 2)


if __name__ == "__main__":
    t = int(input())
    for _ in range(t):
        n = int(input())
        print(and_subarr(n))
