package shapes

import (
	"ray-tracer/intersection"
	"ray-tracer/ray"
)

type Shape interface {
	Intersect(r ray.Ray) intersection.IntersectionPoint
}
