"""
SPDX-License-Identifier: Apache-2.0
SPDX-FileCopyrightText: Copyright 2025 Vladimir Rusinov

Biryani classes
https://www.codechef.com/practice/course/basic-programming-concepts/DIFF500/problems/BIRYANI
"""


def get_biryani_class_cost(weeks: int, cost_per_week: int) -> int:
    """Get biryani class cost"""
    return weeks * cost_per_week


if __name__ == "__main__":
    t = int(input())
    for _ in range(t):
        x, y = map(int, input().split())
        print(get_biryani_class_cost(x, y))
