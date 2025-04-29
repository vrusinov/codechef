# SPDX-FileCopyrightText: 2025 Vladimir Rusinov
# SPDX-License-Identifier: Apache-2.0
"""
IPL Ticket Rush

https://www.codechef.com/practice/course/basic-programming-concepts/DIFF500/problems/IPLTRSH
"""

t = int(input())
for _ in range(t):
    n, m = map(int, input().split())
    print(max(n - m, 0))
