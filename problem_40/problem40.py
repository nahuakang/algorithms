"""
Given an array of integers where every integer occurs three times
except for one integer, which only occurs once, find and return
the non-duplicated integer.

For example,
given [6, 1, 3, 3, 3, 6, 6], return 1,
given [13, 19, 13, 13], return 19.

Do this in O(N) time and O(1) space.
"""
WORD_SIZE = 5


def check_unique(arr):
    non_duplicate = 0

    for i in range(0, WORD_SIZE):
        sum_i_position_bits = 0
        x = 1 << i
        print(f"x is now {x} or {bin(x)}")
        for j in range(len(arr)):
            print(f"arr[j] is currently {arr[j]} or {bin(arr[j])}")
            if arr[j] & x:
                sum_i_position_bits += 1
                print(f"sum_i_position_bits is now {sum_i_position_bits}")

        if sum_i_position_bits % 3:
            non_duplicate |= x

    return non_duplicate


def run_test():
    assert check_unique([6, 1, 3, 3, 3, 6, 6]) == 1
    assert check_unique([13, 19, 13, 13]) == 19
    assert check_unique([0, -1, -1, -1]) == 0
    assert check_unique([2, 2, 2, 6, 10, 10, 10]) == 6
    print("all tests passed...")


if __name__ == "__main__":
    run_test()
