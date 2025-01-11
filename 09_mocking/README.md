# Mocking

## Requirements

* Write a program which counts down from 3, print each number on a new line (with a 1-second pause) and when it reaches zero it will print "Go!" and exit
    * Test that the program sleeps before each next print

## Summary of key points

* Slice up the requirements to iterate:
    * Print 3
    * Print 3, 2, 1 and Go!
    * Wait a second between each line
* We need to write to `stdout` and we use dependency injection (chapter 8) to facilitate testing this, i.e. use `io.Writer` interface
* The test takes 3 seconds to run (slow test) and we have a dependency on `time.Sleep` which we need to control in tests
    * If we can mock `time.Sleep`, we can use dependency injection to use the mocked version to spy on the calls to make assertions
* We define `time.Sleep` dependency as an interface
    * We use `time.Sleep` in `main` and a spy sleeper in test 
    * The spy sleeper is a type which implement `Sleep` interface
* Spies are a kind of mock which can record how a dependency is used, e.g., arguments sent in, number of times it has been called, etc
* If lots of things have to be mocked to test a functionality, it is a sign that the functionality is doing too many things (and thus too many dependencies to mock), so break the module apart so it does less
* Test useful behaviour not every implementation detail (except they are really important)
* Do not test private functions as they are implementation detail to support public behaviour. Private functions are normally "less stable" and we don't want to couple our tests to them
