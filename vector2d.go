package main

import "math"

type Vector2d struct {
	x, y float64
}

func (v1 Vector2d) Add(v2 Vector2d) Vector2d {
	return Vector2d{v1.x + v2.x, v1.y + v2.y}
}
func (v1 Vector2d) Substract(v2 Vector2d) Vector2d {
	return Vector2d{v1.x - v2.x, v1.y - v2.y}
}
func (v1 Vector2d) Multiply(v2 Vector2d) Vector2d {
	return Vector2d{v1.x * v2.x, v1.y * v2.y}
}
func (v1 Vector2d) AddValue(value float64) Vector2d {
	return Vector2d{v1.x + value, v1.y + value}
}
func (v1 Vector2d) SubstractValue(value float64) Vector2d {
	return Vector2d{v1.x - value, v1.y - value}
}
func (v1 Vector2d) MultiplyValue(value float64) Vector2d {
	return Vector2d{v1.x * value, v1.y * value}
}
func (v1 Vector2d) DivideValue(value float64) Vector2d {
	return Vector2d{v1.x / value, v1.y / value}
}
func (v1 Vector2d) limit(lower, upper float64) Vector2d {
	return Vector2d{math.Min(math.Max(v1.x, lower), upper), math.Min(math.Max(v1.y, lower), upper)}
}
func (v1 Vector2d) distance(v2 Vector2d) float64 {
	return math.Sqrt(math.Pow(v1.x-v2.x, 2) + math.Pow(v1.y-v2.y, 2))
}
