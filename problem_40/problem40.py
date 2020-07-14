"""
Given an array of integers where every integer occurs three times
except for one integer, which only occurs once, find and return
the non-duplicated integer.

For example,
given [6, 1, 3, 3, 3, 6, 6], return 1,
given [13, 19, 13, 13], return 19.

Do this in O(N) time and O(1) space.
NOTE: I choose to use O(log N) space instead of bitwise operators.
"""


def check_unique(arr):
    results = dict()
    for val in arr:
        if val in results:
            results[val] += 1
        else:
            results[val] = 1

    for key in results:
        if results[key] == 1:
            return key


def run_test():
    assert check_unique([6, 1, 1, 1, 3, 6, 6]) == 3
    assert check_unique([13, 19, 13, 13]) == 19
    assert check_unique([0, -1, -1, -1]) == 0
    assert check_unique([2, 2, 2, 6, 10, 10, 10]) == 6
    print("all tests passed...")


if __name__ == "__main__":
    run_test()
