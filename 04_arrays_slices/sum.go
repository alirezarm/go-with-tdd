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

// SumAll takes a varying number of slices, returning a new slice 
// containing the totals for each slice passed in
func SumAll(numbersToSum ...[]int) []int {
	// lengthOfNumbers := len(numbersToSum)
	// sums := make([]int, lengthOfNumbers)
	var sums []int
	for _, numbers := range numbersToSum {
		// sums[i] = Sum(numbers)
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// SumAllTails takes a varying number of slices, returning a new slice 
// containing the totals for each slice passed in, excluding the first element (the "head")
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		// Handle the case where the slice is empty
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
