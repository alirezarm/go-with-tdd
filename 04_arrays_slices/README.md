# Iteration

## Requirements

Write a sum func which takes an array of elements and return the total

## Summary of key points

* Arrays have fixed capacity
* Array can be initialized in two ways:
    * [N]type{value1, value2, ..., valueN} e.g. `numbers := [5]int{1, 2, 3, 4, 5}`
    * [...]type{value1, value2, ..., valueN} e.g. `numbers := [...]int{1, 2, 3, 4, 5}`
* To print arrays in error messages, use `%v` (print default format)
* `array[index]`: value out of an array at a index
* Use of range for arrays in for loop, e.g., `for _, number := range numbers {}`
    * range iterate over an array, on each iteration returns the index and the value
* Slices do not have encoded size
* To check test coverage use: `go test -cover`
* 