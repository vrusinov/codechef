# SPDX-FileCopyrightText: 2025 Vladimir Rusinov <vladimir.rusinov@gmail.com>
# SPDX-License-Identifier: Apache-2.0
"""
Bone Appetit

https://www.codechef.com/practice/course/basic-programming-concepts/DIFF500/problems/BNE_APT
"""

n, m = map(int, input().split())
x, y = map(int, input().split())

treats = n * x + m * y
print(treats)
