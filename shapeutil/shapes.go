package shapeutil

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func closures() {
	var functions []func()

	fn := func(val int) func() {
		return func() {
			fmt.Println(val)
		}
	}

	for i := 0; i < 10; i++ {
		functions = append(functions, fn(i))
	}

	for _, f := range functions {
		f()
	}
}

func zero(xPtr *int) {
	*xPtr = 0
}

func pointerTest() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0
}