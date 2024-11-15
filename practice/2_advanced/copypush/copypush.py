#!/usr/bin/python3

# SPDX-FileCopyrightText: 2022 Google Inc
# SPDX-FileCopyrightText: 2022,2024 Vladimir Rusinov <vrusinov@google.com>
# SPDX-License-Identifier: Apache-2.0

from typing import Literal

COPY = "c"
PUSH_BACK = "p"


def copypush(line: str, prev_op: str = "") -> Literal["yes"] | Literal["no"]:
    """Copy and Push Back.

    Idea: start from given line and try to apply Anon's actions in reverse.
    Try splitting or cutting characters from line to see if we can end up
    in empty string. Keep track of previous actions.
    """
    line_len = len(line)
    if line_len == 0:
        return "yes"
    # First, try to see if we can split the string into two equals.
    if line_len % 2 == 0:
        if line[: line_len // 2] == line[line_len // 2 :]:
            if copypush(line[line_len // 2 :], prev_op=COPY) == "yes":
                return "yes"
    # We tried splitting and it didn't work. Let's see if we can append.
    if prev_op == PUSH_BACK:
        return "no"
    return copypush(line[:-1], PUSH_BACK)


if __name__ == "__main__":
    t = input()
    for _ in range(int(t)):
        unused_n = input()
        line = input()
        print(copypush(line))
