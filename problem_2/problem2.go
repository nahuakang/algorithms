/*
Given an array of integers, return a new array such that each element
at index i of the new array is the product of all the numbers in the original
array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output
would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1],
the expected output would be [2, 3, 6].

Follow-up: what if you can't use division?
*/

package problem2

// Multiplier accepts an array and returns a new array with each element at
// index i having the product of all elements in the array except its own value
func Multiplier(arr []int) []int {
	length := len(arr)
	prod := ArrayProduct(arr)
	result := make([]int, length)

	for i, v := range arr {
		result[i] = prod / v
	}

	return result
}

// ArrayProduct returns an integer as the product of all elements in the array
func ArrayProduct(arr []int) int {
	prod := 1
	for _, val := range arr {
		prod *= val
	}
	return prod
}
