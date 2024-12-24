# Pointers & Errors 

## Requirements

* Make a banking system: a `wallet` struct which let us deposit/widthraw `Bitcoin` but does not allow accessing fields directly with their name but rather access is possible via methods.

## Summary of key points

* In Go, a symbol (variable, type, func) starts with a lowercase name is private outside its package
* In Go, when we call a func/method, the arguments are copied, e.g., when calling `func (w Wallet) Deposit(amount int)` the `w` is a copy of whatever we called the method from
    * If we write a function that needs to mutate state we need it to take a pointer to the thing we want to change
    * Paasing a copy of values is useful for a lot of use-cases. However, sometimes we do not want to make a copy of something and instead need to pass a reference., e.g, referencing very large data structures or things where only one instance is necessary (like database connection pools)
* To find the memory address of a value, e.g., wallet, `&wallet.balance`
    * We get the pointer (memory address) of a value by placing an `&` at the beginning of the symbol
* Use `%p` to print memory addresses in base 16 notation with leading `0xs` 
*  Pointers point to some values which we can change, i.e., if take a pointer to wallet, we can change the original values within it. Without the pointer, we have a copy of the wallet
    * Change the receiver type of the methods to `*Wallet` rather than `Wallet`, i.e., a pointer to a wallet
    * Struct pointers are automatically dereferenced, i.e, Go permits writting w.balance, without an explicit dereference but using `(*w).balance` is correct too
* In Go, we can create a new type from an existing type using `type MyName OriginalType`, e.g., `type Bitcoin int`
* In Go, to indicate an error it is idiomatic for the func to return an err for the caller to check and act on, e.g., to signal a problem when using Withdraw
    * Withdraw should return an error if we attempt to withdraw more than we have and the balance should stay the same
        * Our initial intent for Withdraw was to just call it
    * In test, we check if an error has returned. If not, we fail the test if the error is `nil` (`nil` is synonymous with `null`)
    * Errors can be `nil` because the return type of Withdraw will be error (an interface)
    * If a function that takes arguments or returns values that are interfaces, they can be nillable
    * Accessing a value which is `nil` will throw a runtime `panic`. So, we should make sure that you check for nils
    * `errors.New` from package `errors` creates a new error with a use message
    * Errors can be converted to a string with the `.Error()` method
* Use linters to find unchecked errors, e.g., 
    * Install `errcheck` linter: `go install github.com/kisielk/errcheck@latest`
    * Run `errcheck .`