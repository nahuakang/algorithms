"""
Given an array of integers, return a new array such that each element
at index i of the new array is the product of all the numbers in the original
array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output
would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1],
the expected output would be [2, 3, 6].

Follow-up: what if you can't use division?
"""
from functools import reduce


def multiplier(arr):
    prod = reduce(lambda x, y: x * y, arr, 1)
    return [int(prod / x) for x in arr]


def run_test():
    assert multiplier([1, 2, 3, 4, 5]) == [120, 60, 40, 30, 24]
    assert multiplier([3, 2, 1]) == [2, 3, 6]
    print("all tests passed...")


if __name__ == "__main__":
    run_test()
