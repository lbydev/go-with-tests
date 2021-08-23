package main

import (
	"math"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func Perimeter(width float64, height float64) float64 {
	return 2*(width + height)
}

func Area(width float64, height float64) float64 {
	return width * height
}

type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (c Triangle) Area() float64 {
	return (c.Base * c.Height) * 0.5
}

type Shape interface {
	Area() float64
}