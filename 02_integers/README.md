# Integers

## Requirements

Writing and add functions 

## Summary of key points

* When a func has more than one arguments of the same type, the signature can be shorten, e.g. instead of `Add(int x, int y)` use `Add(x, y int)` 
* We can add `Testable Examples` using functions with name starting with `Example` in go test file. Example functions are compiled together with tests, thus the documentation's examples always reflect current code behavior. To see examples while running test:
    * Install `pkgsite` package: `go install golang.org/x/pkgsite/cmd/pkgsite@latest`
    * Add it to your path: e.g., on MacOS, `export PATH=$HOME/go/bin:$PATH`
    * The example will always will be always compilled but to execute it, add a special comment line at the end of the example func, e.g., `// Output: 9`