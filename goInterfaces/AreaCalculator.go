package main

import (
	"fmt"
	"math"
)

type Reactangle struct {
	Length float64
	Width  float64
}

type Square struct {
	Side float64
}

type Circle struct {
	Radius float64
}

type Area interface {
	CalcArea() float64
}

func (r *Reactangle) CalcArea() float64 {
	return r.Length * r.Width
}

func (r *Reactangle) CalcPermimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func (s *Square) CalcArea() float64 {
	return math.Pow(s.Side, 2)
}

func (s *Square) CalcPermimeter() float64 {
	return 4 * s.Side
}

func (c *Circle) CalcArea() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c *Circle) CalcPermimeter() float64 {
	return math.Pi * (2 * c.Radius)
}

func FindArea(a Area) {
	fmt.Println("Area is:", math.Round(a.CalcArea()))
}

func main() {
	rectangle := Reactangle{Length: 5.00, Width: 6.00}
	Square := Square{Side: 5.00}
	circle := Circle{Radius: 5.00}

	FindArea(&rectangle)
	FindArea(&Square)
	FindArea(&circle)

}
