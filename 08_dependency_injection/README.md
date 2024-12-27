# Dependency Injection

## Requirements

* Write a greeting func similar to hello-world (chapter 1) and test the actual printing

## Summary of key points

* Requires understanding of interfaces, see structs (chapter 5)
* `fmt.Printf` prints to stdout, which is hard for to capture using the testing framework
* We need to `inject` (`pass in`) the dependency of printing
* We can accept an `interface` rather than a concrete type because the func doesn't need to know where/how the printing happens
    * This way, we can change the implementation to print to something we control so that we can test it
    * In "real life" we would inject in something that writes to stdout
    * Under the hood, `Printf` passes `os.Stdout` to `Fprintf` which expects an interface called `io.Writer` which put its argument data somewhere
    * To make our `Greet` func testable, instead of using `bytes.Buffer` as writer we use `io.Writer` as writer
        * This enables use to easily test `Greet` by passing a buffer during test
        * In the application (i.e. in `main` func), we can use other writer, e.g., `os.Stdout` or `http.ResponseWriter` (which also implements the `io.Writer` interface)
        * This also separate the concerns: decoupling where the data goes (e.g., where to send the greet?) from how to generate it (e.g. how to greet?)
