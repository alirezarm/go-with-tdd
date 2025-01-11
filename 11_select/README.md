# Select

## Requirements

* Create a func called `WebsiteRacer` which takes two URLs and hits both URLs with an HTTP GET and returning the URL which returned first
    * If none of them return within 10 seconds, returns an error


## Summary of key points

* We use goroutine `select` to synchronise processes
* Use `http.Get` to perform an HTTP GET request
    * It returns an `http.Response` and an `error`
* Time functions
    * `time.Now`: returns current time 
    * `time.Since`: returns a `time.Duration` of the difference of current time and its argument
* Testing the functionality against real websites is slow and may or may not work
    * Go's standard libraray has a package called `net/http/httptest` which enables us to create a mock HTTP server
    * `httptest.NewServer` takes an `http.HandlerFunc` which we are sending in via an anonymous function
        * `http.HandlerFunc` is a type: `type HandlerFunc func(ResponseWriter, *Request)`
        * I t needs a function that takes a `ResponseWriter` and a `Request`
        * This is how a real HTTP server in Go is written but we wrap it in an `httptest.NewServer` which makes it easier to use with testing, i.e., it finds an open port to listen on and then we can close it when the test is done
* `defer`: if we prefix a func call with `defer`, it will be called at the end of the containing func
    * It is good for cleaning up resources, e.g., closing a file or in our case closing a server so that it does not continue to listen to a port
* Go's `select` construct
    * Concurrently check several things, e.g., in the case of website racer we can check them at the same time and see which one comes back first (thier exact response times do not matter)
    * `ping` func creates a `chan struct{}` type which is the smallest data type with no memeory allocation required (in comparision to using type like `bool`)
        * We just close the chan and do not send any data to it
        * As soon as goroutine is done, it sends a signal to the chan which is what we want to know which website is faster
        * We set up two channels, one for each of our URLs, i.e. whichever writes to its channel first will have its code executed in the select
    * In Go concurrency, we can wait for values to be sent to a channel with `myVar := <-ch` which isa blocking call, as we're waiting for a value. `select` allows us to wait on multiple channels where the first one to send a value "wins" and the code underneath the case is executed
* Always use `make` to create channel
    * Using `var ch chan struct{}` to create a chan will initialized it with "zero" value of the type, e.g. for channel zero value is `nil`
        * If we try to send to it with `<-`, it will block forever as we cannot send to `nil` channels
* It is possible that channels we are listening on never return a value
    * `time.After` returns a chan (like ping func) and will send a signal to it after the defined amount of time. This means `time.After` is a very handy function when using `select`