# Reflection

## Requirements

* Create a counter which is safe to use concurrently
    * Start with an unsafe counter in a single-threaded env
    * Confirm it is unsafe with multiple goroutines and fix it 

## Summary of key points

* `sync.WaitGroup`: a convenient way of synchronising concurrent processes
    * A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished
    * After `wg.Wait()` finished, we can be sure all of our goroutines have attempted to Inc the Counter
* When multiple goroutines are trying to mutate the value of the counter at the same time, the final value of the counter will not be correct
    * Simple solution is to add a lock to our Counter using `Mutex` to ensure only one goroutine can increment the counter at a time
        * A `Mutex` is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex
        * Any goroutine calling Inc will acquire the lock on Counter if they are first. All the other goroutines will have to wait for it to be Unlocked before getting access
        * A Mutex must not be copied after first use
        * When we pass our Counter (by value) to `assertCounter` it will try and create a copy of the mutex, i.e. we should pass in a pointer instead
* Vet as in `go vet` examines Go source code and reports suspicious constructs, such as Printf calls whose arguments do not align with the format string. Vet uses heuristics that do not guarantee all reports are genuine problems, but it can find errors not caught by the compilers
