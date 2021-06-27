#!/usr/bin/python3

def ciel(a: int, b: int) -> int:
    correct = a - b
    if (correct % 10) == 0:
        return correct + 1
    elif correct == 1:
        return 2
    else:
        return correct - 1


def main():
    a, b = [int(x) for x in input().split()]
    print(ciel(a, b))

if __name__ == '__main__':
    main()
