# SPDX-FileCopyrightText: 2025 Vladimir Rusinov
# SPDX-License-Identifier: Apache-2.0
"""
Kitchen Timings

https://www.codechef.com/practice/course/basic-programming-concepts/DIFF500/problems/KITCHENTIME
"""

t = int(input())
for _ in range(t):
    x, y = map(int, input().split())
    print(y - x)
