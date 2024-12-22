package arrays

// Sum takes an array of numbers and returns the sum of them
func Sum(numbers []int) int {
	var sum int
	// for i := 0; i < 5; i++ {
	for _, number := range numbers {
		sum += number
	}
	return sum
}