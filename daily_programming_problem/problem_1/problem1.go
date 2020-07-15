/*
Given a list of numbers and a number k, return whether any two numbers from
the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

Bonus: Can you do this in one pass?
*/

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
