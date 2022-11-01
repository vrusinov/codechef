#!/usr/bin/python3
def subarray_removal(arr):
    num_zeroes = 0
    num_ones = 0
    for i in arr:
        if i == 1:
            num_ones += 1
        else:
            num_zeroes += 1
    m = min(num_zeroes, num_ones)
    if num_zeroes > num_ones:
        return num_ones
    else:
        return num_zeroes + (num_ones - num_zeroes) // 3

if __name__ == '__main__':
    t = int(input())
    for _ in range(t):
        unused_n = input()
        arr = [int(x) for x in input().split()]
        print(subarray_removal(arr))