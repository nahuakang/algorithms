def count_inversion(arr):
    if len(arr) < 2:
        return arr, 0

    mid = len(arr) // 2
    left = arr[:mid]
    right = arr[mid:]

    sorted_left, left_inv_count = count_inversion(left)
    sorted_right, right_inv_count = count_inversion(right)

    return merge_and_count(sorted_left, sorted_right,
                           left_inv_count+right_inv_count)


def merge_and_count(left, right, previous_count):
    result = [None] * (len(left) + len(right))
    i = j = split_inv_count = 0

    for k in range(len(result)):
        if i > len(left) - 1 and j <= len(right) - 1:
            result[k] = right[j]
            j += 1
        elif j > len(right) - 1 and i <= len(left) - 1:
            result[k] = left[i]
            i += 1
        elif left[i] <= right[j]:
            result[k] = left[i]
            i += 1
        elif left[i] > right[j]:
            result[k] = right[j]
            split_inv_count += len(left) - i
            j += 1

    return result, (split_inv_count + previous_count)


def test():
    test_arr_0 = [0, 0, 0, 0, 0, 0, 0]
    test_arr_1 = [1, 3, 5, 2, 4, 6]
    test_arr_2 = [1, 3, 5, 7, 2, 4, 6]
    test_arr_3 = [5, 3, 1, 6, 4, 2]
    test_arr_4 = [6, 5, 4, 3, 2, 1]
    assert count_inversion(test_arr_0)[1] == 0
    assert count_inversion(test_arr_1)[1] == 3
    assert count_inversion(test_arr_2)[1] == 6
    assert count_inversion(test_arr_3)[1] == 9
    assert count_inversion(test_arr_4)[1] == 15


if __name__ == "__main__":
    test()
    print("test passed")
