#!/usr/bin/python3

def eopr(x: int, y: int) -> int:
    """Choose one positive odd integer a and one positive even integer b.

    Do the cleaning.

    Returns:
        Minimum number of rounds (possibly zero) to make lab clean
    """
    if x == y:
        return 0
    if x < y:
        # We set b = 2.
        if (x % 2 == 0) and (y % 2 == 0):
            a = y - x  # even, but must be odd
            a = a // 2
            if a % 2 == 0:
                # a += 1
                # x = x + a
                # x = x + a
                # b = 2
                # x = x - b
                return 3
            else:
                # x = x + a
                # x = x + a
                return 2
        if (x % 2 == 1) and (y % 2 == 0):
            # a = y -x # odd
            # x = x + a
            return 1
        if (x % 2 == 0) and (y % 2 == 1):
            # same as above
            return 1
        if (x % 2 == 1) and (y % 2 == 1):
            # same as when both even
            a = y - x  # even, but must be odd
            a = a // 2
            if a % 2 == 0:
                return 3
            else:
                return 2
    else:  # x > y
        if (x % 2 == 0) ^ (y % 2 == 0):
            # a = 1
            # b = y - x + 1 # odd
            # x = x - b
            # x = x + a
            return 2
        else:
            # if both are even or both are odd, we can do single x = x - b
            return 1
    return -1  # Should never be reached but pylance complains without it


def main():
    # Number of test cases
    t = int(input())
    for _ in range(t):
        x, y = [int(i) for i in input().split()]
        print(eopr(x, y))

if __name__ == '__main__':
    main()
