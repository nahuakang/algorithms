"""
Given a list of numbers and a number k, return whether any two numbers from
the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

Bonus: Can you do this in one pass?
"""


def list_checker(arr, k):
    second_pairs = []

    for item in arr:
        if item == k or item in second_pairs:
            return True
        else:
            second_pairs.append(k - item)

    return False


def run_test():
    assert list_checker([10, 15, 3, 7], 17)
    assert list_checker([2, 0, -1, 3], 1)
    assert list_checker([1], 1)
    assert not list_checker([3, 10, 4, 6], 1)
    print("Tests passed...")


if __name__ == "__main__":
    run_test()
