# Iteration

## Requirements

* Write a Sum func which takes an array/slice of elements and return the total
* Extend Sum to SumAll  func to get varying number of slices and return a new slice with the sum for each input slice
* Extend Sum to SumAllTails to calculate the sum of the tails of each slice where the tail of a slice is its items except the first one

## Summary of key points

* Arrays have fixed capacity
* Array can be initialized in two ways:
    * [N]type{value1, value2, ..., valueN} e.g. `numbers := [5]int{1, 2, 3, 4, 5}`
    * [...]type{value1, value2, ..., valueN} e.g. `numbers := [...]int{1, 2, 3, 4, 5}`
* To print arrays in error messages, use `%v` (print default format)
* `array[index]`: value out of an array at a index
* Use of `range` for arrays in for loop, e.g., `for _, number := range numbers {}`
    * range iterate over an array, on each iteration returns the index and the value
* Slices do not have encoded size
    * Slices/arrays work with any other type too, including arrays/slices themselves, e.g. `[][]string` declare a variable representing a 2D array 
* To check test coverage use: `go test -cover`
* To pass a varying number of inputs to a func (i.e., `variadic functions`) use `name ...[]type` syntax, e.g., `func SumAll(numbersToSum ...[]int) []int { }`
* Equality operators do not work with slices. We can either 
    * iterate over the slices and use equality operators
    * use `slices.Equal` which does a simple shallow compare on slices which is type safe but the elements must be comparable. 
    * use `reflect.DeepEqual` (not type safe, i.e., `reflect.DeepEqual([]int{1,2}, "abc")` will compile!)
* Func to get length of a slice is `len` but `cap` func returns its capacity (i.e. number of elements it can keep)
* `append` func takes a slice and a value and returns a new slice
* Syntax to slice a slice: `slice[low:high]`
    * By omitting `low` or `high`, we include everything to that side of the slice
* We can assign a func to a variable, e.g., `checkSums := func(t testing.TB, got, want []int) { }` (i.e., functions are value too)
    * This way, we hide the func that do not need to be exported 



