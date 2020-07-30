package light

import (
	"ray-tracer/color"
	"ray-tracer/vector"
)

type Light struct {
	Position vector.Vector
	Power color.Color
	Attenuation vector.Vector
}

func (light Light) Distance(point vector.Vector) float64 {
	return light.Position.Distance(point)
}

func (light Light) IntensityAt(point vector.Vector) color.Color {
	distance := light.Distance(point)
	intensity := light.Power.ScalarDivide((light.Attenuation.X) + (light.Attenuation.Y * distance) + (light.Attenuation.Z * distance * distance))
	return intensity
}
