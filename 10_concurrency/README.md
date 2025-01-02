# Concurrency

## Requirements

* Write a func which check a list of hundreds of websites and return a boolean for their response status


## Summary of key points

* To run benchmark test in current folder: `go test -bench="."`
* A blocking operation makes us wait for it to finish
* A `goroutine` is a separate process which runs a non-blocking operation
* Starting a `goroutine`: use `go` keyword to turn a function call into a go statement, i.e. `go doSomething()`
    * The only way to start a goroutine is to put `go` in front of a func call, we often use anonymous functions when we want to start a goroutine
* Two features of anonymous functions:
    * They can be executed at the same time that they're declared - this is what the () at the end of the anonymous function is doing
    * They maintain access to the `lexical scope` in which they are defined, i.e. all the variables that are available at the point when we declare the anonymous function are also available in the body of the function
* Race condition: a bug that occurs when the output of our software is dependent on the timing and sequence of events that we have no control over
    * Go's maps do not allow more than one goroutine to write to them at once and give fatal error
    * Go has built in race detector which can be run with the `race` flag: `go test --bench="." -race`
* Channels: a go data structure that can both receive and send values and thus allow communication between different processes
    * Channels solve data race by coordinating goroutines, e.g., communication between the parent process and each of the goroutines that it makes to do some work
    * Instead of writing to the map directly, we send the result struct for each call to `wc` function to the `resultChannel` with a send statement using the `<-` operator: `resultChannel <- result{u, wc(u)}`
    * Use a receive expression using the `<-` operator, which assigns a value received from a channel to a variable `r := <-resultChannel`
    * By sending the results into a channel, we can control the timing of each write into the results map, ensuring that it happens one at a time