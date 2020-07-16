def find_smallest(arr):
    smallest = arr[0]
    smallest_index = 0

    for i in range(1, len(arr)):
        if arr[i] < smallest:
            smallest = arr[i]
            smallest_index = i

    return smallest_index


def selection_sort(arr):
    results = []
    for i in range(len(arr)):
        smallest = find_smallest(arr)
        results.append(arr.pop(smallest))

    return results


def run_test():
    assert selection_sort([1, 2, 3, 4, 5]) == [1, 2, 3, 4, 5]
    assert selection_sort([5, 4, 3, 2, 1]) == [1, 2, 3, 4, 5]
    assert selection_sort([-3, 3, -2, 2, -1, 1, 0]) == [-3, -2, -1, 0, 1, 2, 3]
    assert selection_sort([0, 0]) == [0, 0]
    print("all tests passed...")


if __name__ == "__main__":
    run_test()
