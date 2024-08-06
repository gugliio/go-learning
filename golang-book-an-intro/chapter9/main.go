package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func main() {
	// c := new(Circle)
	// c2 := Circle{x: 0, y: 0, r: 5}

	c := Circle{0, 0, 5}

	fmt.Println(c.x, c.y, c.r)
	c.x = 10
	c.y = 5

	area := c.area()
	fmt.Printf("Area of circle with radius %.2f: %.2f\n", c.r, area)
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
