package intersection

import (
	"ray-tracer/material"
	"ray-tracer/vector"
)

type IntersectionPoint struct {
	IsHit    bool
	Position vector.Vector
	Normal   vector.Vector
	Distance float64
	Material material.Material
}
