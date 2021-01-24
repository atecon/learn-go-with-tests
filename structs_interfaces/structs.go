package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

// func (receiverName ReceiverType) MethodName(args)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

func Area(r Rectangle) (area float64) {
	return r.Width * r.Height
}

// tell Go what a Shape is using an interface declaration
// By declaring an interface the helper is decoupled from
// the concrete types and just has the method it needs to do its job.
type Shape interface {
	Area() float64
}
