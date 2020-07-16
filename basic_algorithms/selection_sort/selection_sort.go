package selectionSort

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0

	for i, v := range arr {
		if v < smallest {
			smallest = v
			smallestIndex = i
		}
	}

	return smallestIndex
}

// SelectionSort uses O(n^2) to sort an array from smallest to largest
func SelectionSort(arr []int) []int {
	var results []int
	loopTimes := len(arr)

	for i := 0; i < loopTimes; i++ {
		smallestIndex := findSmallest(arr)
		results = append(results, arr[smallestIndex])
		arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...)
	}

	return results
}
