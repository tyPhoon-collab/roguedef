package ds

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

func (v Vec2) Distance(v2 Vec2) float64 {
	return v.Sub(v2).Len()
}

func (v Vec2) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}

func (v Vec2) AngleTo(v2 Vec2) float64 {
	return v.Sub(v2).Angle()
}

func (v Vec2) Normalize() (Vec2, error) {
	length := v.Len()
	if length == 0 {
		return Vec2{}, &ZeroDivisionError{}
	}
	return v.DivScalar(length)
}

func (v Vec2) DirTo(v2 Vec2) (Vec2, error) {
	return v2.Sub(v).Normalize()
}

func (v Vec2) String() string {
	return fmt.Sprintf("(%f, %f)", v.X, v.Y)
}

func Unpack[T int | float64 | float32](v Vec2) (T, T) {
	return T(v.X), T(v.Y)
}
