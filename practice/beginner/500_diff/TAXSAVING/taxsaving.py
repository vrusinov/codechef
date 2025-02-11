# SPDX-FileCopyrightText: 2025 Vladimir Rusinov <vladimir.rusinov@gmail.com>
# SPDX-License-Identifier: Apache-2.0
"""
Saving Taxes.

https://www.codechef.com/practice/course/basic-programming-concepts/DIFF500/problems/TAXSAVING
"""

t = int(input())

for _ in range(t):
    case = input().split()
    x = int(case[0])
    y = int(case[1])
    print(x - y)
