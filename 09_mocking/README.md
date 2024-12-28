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