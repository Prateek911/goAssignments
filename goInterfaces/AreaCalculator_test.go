package main

import (
	"math"
	"testing"
)

func TestRectangleArea(t *testing.T) {
	rect := Reactangle{Length: 5.00, Width: 6.00}
	expectedArea := 30.00
	if area := rect.CalcArea(); area != expectedArea {
		t.Errorf("Actual area is, %2f, got %2f", expectedArea, area)
	}
}

func TestSquareArea(t *testing.T) {
	square := Square{Side: 5.00}
	expectedArea := 25.00
	if area := square.CalcArea(); area != expectedArea {
		t.Errorf("Actual area is, %2f, got %2f", expectedArea, area)
	}
}

func TestCircleArea(t *testing.T) {
	circle := Circle{Radius: 5.00}
	expectedArea := math.Pi * math.Pow(5, 2)
	if area := circle.CalcArea(); area != expectedArea {
		t.Errorf("Actual area is, %2f, got %2f", expectedArea, area)
	}
}
