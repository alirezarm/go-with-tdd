# Hello World

## Requirements

Write hello world to a recipient and in a specified language.  


## Summary of key points

* Go program have a `main` package with a `main` func in it
* Keyword `func` defines a function
* Import a package using `import` keyword, e.g., `import "fmt"`
* `Println` func is in `fmt` package
* Go test files should be names as `xxx_test.go`
* Name of test func within a Go test file should start with `Test`, e.g., `TestHello`
* Test func gets one argument `t *testing.T` (import testing package i.e. `import "testing"`) where `t` is the hook to the testing framework allowing us to use things like `t.Errorf`
* to access Go docs from commane line use `go doc fmt`
* Constants are defined using `const` keyword, e.g., `const helloPrefix = "Hello, "`
* Use `t.Run` to write subtests to write tests describing different scenarios
* `testing.TB` is an interface that both `*testing.T` (test) and `*testing.B` (benchmark) both satisfy
* Using `t.Helper()` in a mehtod tells the test suite that this method is a helper
