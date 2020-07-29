package intersection

import (
	"ray-tracer/vector"
)

type IntersectionPoint struct {
	IsHit    bool
	Position vector.Vector
	Normal   vector.Vector
}
