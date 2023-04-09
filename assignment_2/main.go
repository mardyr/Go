package main

/* Write a program that create two customs struct types called
'triangle' & 'square

Type square shall be a struct with a field called 'sideLength'
of type float64

The triangle type shall be a struct with field called height
of type float64 and field called `base` of type float64

Both types should have function called `getArea` that returns
the calculated area of the square or triangle

Area of triangle = 0.5*base*height
Area of a square = sideLength*sideLength

Add a `shape` interface that defines function called `printArea`
This function shall calculate the area of the given shape and
print it out to the terminal. Design the interface so that the
`printArea` function can be called with either a triangle or a square

*/

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLenght float64
}

type triangle struct {
	height float64
	base   float64
}

func main() {

	mySquare := square{
		sideLenght: 5,
	}
	myTriangle := triangle{
		height: 8,
		base:   12,
	}
	printArea(mySquare)
	printArea(myTriangle)

}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

func (sq square) getArea() float64 {
	return sq.sideLenght * sq.sideLenght
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.height * tr.base
}
