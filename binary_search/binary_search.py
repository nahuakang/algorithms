def binary_search(arr, item):
    """Binary Search assumes that the input array is sorted."""
    low = 0
    high = len(arr) - 1

    while low <= high:
        mid = (low + high) // 2
        guess = arr[mid]

        if guess == item:
            return guess
        elif guess < item:
            low = mid + 1
        else:
            high = mid - 1

    return None


def run_test():
    assert binary_search([1, 2, 3, 4, 5, 6], 6) == 6
    assert binary_search([1, 2, 3, 4, 5, 6], 0) is None
    print("tests passed...")


if __name__ == "__main__":
    run_test()
