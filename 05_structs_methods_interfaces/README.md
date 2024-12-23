# Structs & Methods & Interfaces 

## Requirements

* Calculate the perimeter/area of a rectangle/circles/others given a height and width

## Summary of key points

* If we write perimeter/area functions providing width and height without no notion of rectangle, one might mistakenly provide width and height of a triangle and get the wrong answer
    * One solution is to use specific names, e.g., `RectangleArea`
    * A btter solution is to create a type representing rectangle to encapsulate the concept of rectangle. We can create a type using `struct` which is basically a named collection of fields to store data
* Define a type using `struct`, e.g., `type Rectangle struct {}`
    * `Rectangle` then can be passed to a func as a type, e.g., `func Perimeter(rectangle Rectangle) {}`
* Use of `%g` rather than `%f` will print a more precise decimal number in the error message
* Structs can have `methods`. A method is a function with a receiver. Methods are very similar to functions but they are called by invoking them on an instance of a particular type, i.e. we can only call methods on "things"
    * When declaring a method, the receiver must be specified, e.g. `func (receiverName ReceiverType) MethodName(args)`
    * When the method is called on a variable of type `ReceiverType`, we get a reference to its data via `receiverName` (this is similar to `this` in other languages)
    * Convetion is to use first letter of `ReceiverType` as `receiverName`, e.g., `r Rectangle`
* We want to get a colelction of shapes and calculate their area, calculate area of rectangles or circles and fail if the passed argument is not a shape
    * We can use `interfaces` which allows us to create functions that can be used with different types and create highly-decoupled code but maintainin type-safety (polymorphism)
    * `interface` is a type and defined as `type Shape interface { Area() float64 }`
    * This is different than other languages where the a type impelments the interface, i.e. Rectangle/Circle have a method called Area that returns a float64 so it satisfies the Shape interface. In Go, if the passed type satisfies what the interface asks for, it will compile (interface resolution is implicit)
    * It is easy to add a new shape, implement Area
* Table driven tests: build a list of test cases that can be tested in the same manner
    * We can use a slice of anonymous structs and fill them with test cases and then we iterate over the slice and run the tests


