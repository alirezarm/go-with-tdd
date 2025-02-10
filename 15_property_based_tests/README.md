# Property Based Tests

## Requirements

* Write a function which converts an Arabic number (numbers 0 to 9) to a Roman Numeral and write another function for the reverse case, i.e. from Roman Numeral to Arabic number
    * I is one and III is three, V means five, but IV is 4 (not IIII)
    * MCMLXXXIV is 1984
    * Romans were into DRY: we can't have the same character repeated more than 3 times in a row
        * Instead we take the next highest symbol and then "subtract" by putting a symbol to the left of it. Not all symbols can be used as subtractors; only I (1), X (10) and C (100)

## Summary of key points

* Use `struct` to add more test cases (exmaple based tests)
* `strings.Builder`: used to efficiently build a string using Write methods which minimizes memory copying
* The for loop does not need to reply on an i and instead can rely on a condition
* We should view `switch` statements with a bit of suspicion, i.e. we are capturing a concept or data inside some imperative code when in fact it could be captured in a class structure instead
* The `HasPrefix(s, prefix)` checks whether string `s` starts with prefix and `TrimPrefix(s, prefix)` removes the prefix from `s`
* Roman Numerals has few rules:
    * Can't have more than 3 consecutive symbols
    * Only I (1), X (10) and C (100) can be "subtractors"
    * ConvertToArabic(ConvertToRoman(N)) = N
* Property based tests: The tests written using struct of exmaples is called "example" based tests where we provide examples for the tooling to verify. In property based testing, we take some rules that we know about our domain and somehow exercise them against our code. Property based tests uses random data to verify the rules to make sure they always hold true. The real challenge about property based tests is having a good understanding of the domain and writing these properties
* `testing/quick` package from the standard library: `quick.Check` function will run against a number of random inputs, if the function returns false it will be seen as failing the check
* We can't do negative numbers with Roman Numerals. Given our rule of a max of 3 consecutive symbols we can't represent a value greater than 3999 (well, kinda) and int has a much higher maximum value than 3999
    * Go has types for unsigned integers, which means they cannot be negative; so that rules out one class of bug in our code immediately. By adding 16 (i.e. uint16), it means it is a 16 bit integer which can store a max of 65535, which is still too big but gets us closer to what we need



