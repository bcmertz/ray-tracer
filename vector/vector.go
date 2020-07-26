package vector

import (
	"math"

	"ray-tracer/color"
)

type Vector struct {
	X     float64
	Y     float64
	Z     float64
	Color color.Color
}

func (v1 Vector) Equals(v2 Vector) (equals bool) {
	return (v1.X == v2.X) && (v1.Y == v2.Y) && (v1.Z == v2.Z)
}

func (v1 Vector) Add(v2 Vector) Vector {
	v := Vector{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
	return v
}

func (v1 Vector) Subtract(v2 Vector) Vector {
	v := Vector{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
		Z: v1.Z - v2.Z,
	}
	return v
}

func (v1 Vector) ScalarMultiply(s float64) Vector {
	v := Vector{
		X: s * v1.X,
		Y: s * v1.Y,
		Z: s * v1.Z,
	}
	return v
}

func (v1 Vector) ScalarDivide(s float64) Vector {
	v := Vector{
		X: v1.X / s,
		Y: v1.Y / s,
		Z: v1.Z / s,
	}
	return v
}

func (v1 Vector) Dot(v2 Vector) float64 {
	return ((v1.X * v2.X) + (v1.Y * v2.Y) + (v1.Z * v2.Z))
}

func (v1 Vector) Cross(v2 Vector) Vector {
	v := Vector{
		X: v1.Y*v2.Z - v1.Z*v2.Y,
		Y: v1.Z*v2.X - v1.X*v2.Z,
		Z: v1.X*v2.Y - v1.Y*v2.X,
	}
	return v
}

func (v1 Vector) Length() float64 {
	return math.Sqrt(v1.Dot(v1))
}

func (v1 Vector) Distance(v2 Vector) float64 {
	return v2.Subtract(v1).Length()
}

func (v1 Vector) Normalize(v2 Vector) Vector {
	length := v1.Length()
	return Vector{
		X: v1.X / length,
		Y: v1.Y / length,
		Z: v1.Z / length,
	}
}

func (v1 Vector) ToString() string {
	return "Vector (" + string(int(v1.X)) + ", " + string(int(v1.Y)) + ", " + string(int(v1.Z)) + ")"
}
