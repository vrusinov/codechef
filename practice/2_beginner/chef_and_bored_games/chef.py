#!/usr/bin/python3

# SPDX-FileCopyrightText: 2022 Google Inc
# SPDX-FileCopyrightText: 2022 Vladimir Rusinov <vrusinov@google.com>
#
# SPDX-License-Identifier: Apache-2.0

"""
Description:

Chef has recently been playing a lot of chess in preparation for the ICCT
(International Chef Chess Tournament).

Since putting in long hours is not an easy task, Chef's mind wanders elsewhere.
He starts counting the number of squares with odd side length on his chessboard..

However, Chef is not satisfied. He wants to know the number of squares of odd
side length on a generic N*NN∗N chessboard.


Solution:

Maths / combinatorics :(

- in 2x2 matrix we have 2*2 + 1 squares
  2*2 of them are odd-length

- in 3x3 matrix we have 1(3*3) + 4(2*2) + 9(1*1) squares
                        ^^^      ^^^     ^^^
                        odd size even    odd

- in 4x4 matrix we have 1(4*4) + 4(3*3) + 9(2*2) + 16(1*1) squares
  so 16 + 4 = 20.

Looks like the formula is n**2 + (n-2)**2 + ...
"""

def num_squares(n):
    """Returns number of squares with odd length."""
    r = 0
    for i in range(n, 0, -2):
        r += i**2
    return r


if __name__ == '__main__':
    t = int(input())
    for _ in range(t):
        n = int(input())
        print(num_squares(n))