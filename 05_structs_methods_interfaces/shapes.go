package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

// interface definition for Shape type which has Area() method that returns float64
type Shape interface{
	Area() float64
}

// method for Rectangle type that returns area as float64
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

// method for Circle type that returns area as float64
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base  float64
	Height float64
}

// method for Triangle type that returns area as float64
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Width * rectangle.Height
// }