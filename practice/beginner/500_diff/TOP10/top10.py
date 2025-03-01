"""
SPDX-License-Identifier: Apache-2.0
SPDX-FileCopyrightText: Copyright 2025 Vladimir Rusinov

Masterchef finals
https://www.codechef.com/practice/course/basic-programming-concepts/DIFF500/problems/TOP10
"""


def is_top10(position: int) -> bool:
    """Check if in top10"""
    return position <= 10


if __name__ == "__main__":
    t = int(input())
    for _ in range(t):
        x = int(input())
        print("YES" if is_top10(x) else "NO")
