package main

import (
	"fmt"
	"math"
)

type Vec2 struct {
	X, Y float64
}

func (v Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func (v Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{
		X: v.X - v2.X,
		Y: v.Y - v2.Y,
	}
}

func (v Vec2) Mul(v2 Vec2) Vec2 {
	return Vec2{
		X: v.X * v2.X,
		Y: v.Y * v2.Y,
	}
}

func (v Vec2) MulScalar(x float64) Vec2 {
	return Vec2{
		X: v.X * x,
		Y: v.Y * x,
	}
}

func (v Vec2) Div(v2 Vec2) (Vec2, error) {
	if v2.X == 0 || v2.Y == 0 {
		return Vec2{}, &ZeroDivisionError{}
	}
	return Vec2{
		X: v.X / v2.X,
		Y: v.Y / v2.Y,
	}, nil
}

func (v Vec2) DivScalar(x float64) (Vec2, error) {
	if x == 0 {
		return Vec2{}, &ZeroDivisionError{}
	}
	return Vec2{
		X: v.X / x,
		Y: v.Y / x,
	}, nil
}

func (v Vec2) Dot(v2 Vec2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v Vec2) Cross(v2 Vec2) float64 {
	return v.X*v2.Y - v.Y*v2.X
}

func (v Vec2) Len() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec2) Normalize() (Vec2, error) {
	length := v.Len()
	if length == 0 {
		return Vec2{}, &ZeroDivisionError{}
	}
	return v.DivScalar(length)
}

func (v Vec2) String() string {
	return fmt.Sprintf("(%f, %f)", v.X, v.Y)
}
