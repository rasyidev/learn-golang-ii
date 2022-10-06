package tambahan

import (
	"fmt"
	"math"
	"testing"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

// function volume yang ada di luar interface
func (c circle) volume() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func TestContohInterface(t *testing.T) {
	c1 := circle{radius: 5}
	r1 := rectangle{width: 3, height: 2}

	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of r1: %T\n", r1)

	fmt.Println(c1.area())
}
