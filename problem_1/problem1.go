package problem1

// ListChecker checks if any two elements in a list adds up to the sum
func ListChecker(arr []int, sum int) bool {
	secondPairs := []int{}
	for _, element := range arr {
		if element != sum && !contains(secondPairs, element) {
			secondPairs = append(secondPairs, sum-element)
		} else {
			return true
		}
	}
	return false
}

func contains(arr []int, target int) bool {
	for _, element := range arr {
		if element == target {
			return true
		}
	}
	return false
}
