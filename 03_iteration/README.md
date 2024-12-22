# Iteration

## Requirements

A function which repeats a char

## Summary of key points

* In Go, there is only `for` (no `while`, `do`, or `until`)
* `for` has no parentheses and always requires braces `{ }`
* Use `var` to declare a variable (and functions), e.g. `var repeated string`
    * `:=` is used to declare and initialize
* Benchmarks are a first-calss citizens of Go. They can be defined in test file using a func name starting with `Benchmark`
    * The `testing.B` gives access to the cryptically named `b.N`, i.e. benchmark runs `b.N` times and measures how long it takes. The framework determines the amount of times the code is run
    * Use `go test -bench=.` to run the benchmark