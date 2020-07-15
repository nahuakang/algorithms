def merge(left, right):
    i = j = 0
    length = len(left) + len(right)
    result = [None] * length
    
    for k in range(length):
        if j > len(right) - 1 and i <= len(left) - 1:
            result[k] = left[i]
            i += 1
        elif i > len(left) - 1 and j <= len(right) - 1:
            result[k] = right[j]
            j += 1
        elif left[i] <= right[j]:
            result[k] = left[i]
            i += 1
        else:
            result[k] = right[j]
            j += 1

    return result


def merge_sort(arr):
    length = len(arr)
    if length < 2:
        return arr

    left = arr[:length // 2]
    right = arr[length // 2:]

    return merge(merge_sort(left), merge_sort(right))


def test():
    test_arr_0 = [3, 3, 3, 2, 2, 2, 1, 1, 1]
    test_arr_1 = [5, 4, 1, 8, 7, 2, 6, 3]
    test_arr_2 = [9, 5, 4, 1, 8, 7, 2, 6, 3]
    assert merge_sort(test_arr_0) == [1, 1, 1, 2, 2, 2, 3, 3, 3]
    assert merge_sort(test_arr_1) == [1, 2, 3, 4, 5, 6, 7, 8]
    assert merge_sort(test_arr_2) == [1, 2, 3, 4, 5, 6, 7, 8, 9]


if __name__ == "__main__":
    test()
    print("tests passed...")
