# SPDX-FileCopyrightText: 2025 Vladimir Rusinov <vladimir.rusinov@gmail.com>
# SPDX-License-Identifier: Apache-2.0
"""
Determine the Score

https://www.codechef.com/practice/course/basic-programming-concepts/DIFF500/problems/DETSCORE
"""

if __name__ == "__main__":
    t = int(input())
    for _ in range(t):
        case = input().split()
        x = int(case[0])
        n = int(case[1])

        score = x / 10 * n

        print(int(score))
