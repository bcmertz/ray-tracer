package ray

import (
	"ray-tracer/vector"
)

type Ray struct {
	Origin    vector.Vector
	Direction vector.Vector
}

func (r Ray) PointAtParameter(t float64) vector.Vector {
	return r.Origin.Add(r.Direction.ScalarMultiply(t))
}
