package color

import "math"

type Color struct {
	R float64
	G float64
	B float64
}

func (c1 Color) Add(c2 Color) Color {
	return Color{
		R: c1.R + c2.R,
		G: c1.G + c2.G,
		B: c1.B + c2.B,
	}
}

func (c1 Color) ColorMultiply(c2 Color) Color {
	return Color{
		R: c1.R * c2.R,
		G: c1.G * c2.G,
		B: c1.B * c2.B,
	}
}

func (c1 Color) ScalarMultiply(s float64) Color {
	return Color{
		R: c1.R * s,
		G: c1.G * s,
		B: c1.B * s,
	}
}

func (c1 Color) ScalarDivide(s float64) Color {
	return Color{
		R: c1.R / s,
		G: c1.G / s,
		B: c1.B / s,
	}
}

func (c1 Color) Clamp(lower, upper float64) Color {
	return Color{
		R: math.Max(lower, math.Min(c1.R, upper)),
		G: math.Max(lower, math.Min(c1.G, upper)),
		B: math.Max(lower, math.Min(c1.B, upper)),
	}
}

func (c1 Color) ToStr() string {
	return "Color (" + string(int(c1.R)) + ", " + string(int(c1.G)) + ", " + string(int(c1.B)) + ")"

}
