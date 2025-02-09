# Context

## Requirements

* Start a web server that when hit kicks off a potentially long-running process to fetch some data for it to return in the response. A user cancels the request before the data can be retrieved and we have to make sure the process is told to give up



## Summary of key points

* `Context` package helps us manage long-running, resource-intensive processes (often in goroutines). If the action that caused this gets cancelled or fails for some reason we need to stop these processes in a consistent way through your application
    * The context package provides functions to derive new Context values from existing ones. These values form a tree: when a Context is canceled, all Contexts derived from it are also canceled.
    * Contexts must be derived so that cancellations are propagated throughout the call stack for a given request, e.g, we derive a new `cancellingCtx` from the request which returns a `cancel` function. We then schedule that function to be called in 5 milliseconds by using `time.AfterFunc`. Finally we use this new context in our request by calling `request.WithContext`
    * Context has a method `Done()` which returns a channel which gets sent a signal when the context is "done" or "cancelled". We want to listen to that signal and call `store.Cancel` if we get it but we want to ignore it if our Store manages to Fetch before it. To manage this we run `Fetch` in a goroutine and it will write the result into a new channel `data`. We then use `select` to effectively race to the two asynchronous processes and then we either write a response or Cancel
* Our web server should not be concerned with manually cancelling Store. The store can depend on other slow-running processes. We have to make sure that `Store.Cancel` correctly propagates the cancellation to all of its dependants
* Context is a consistent way of offering cancellation:
    * Incoming requests to a server should create a Context, and outgoing calls to servers should accept a Context
    * The chain of function calls between them must propagate the Context, optionally replacing it with a derived Context created using WithCancel, WithDeadline, WithTimeout, or WithValue. When a Context is canceled, all Contexts derived from it are also canceled
    * A context parameter must be passed as the first argument to every function on the call path between incoming and outgoing requests. This allows Go code developed by many different teams to interoperate well. It provides simple control over timeouts and cancellation and ensures that critical values like security credentials transit Go programs properly
    * Following this approach, we pass through the context to our Store and let it be responsible. That way it can also pass the context through to its dependants and they too can be responsible for stopping themselves


