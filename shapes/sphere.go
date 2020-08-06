package shapes

import (
	"math"
	"ray-tracer/intersection"
	"ray-tracer/material"
	"ray-tracer/ray"
	"ray-tracer/vector"
)

type Sphere struct {
	Center   vector.Vector
	Radius   float64
	Material material.Material
}

func (sphere Sphere) Intersect(ray ray.Ray) intersection.IntersectionPoint {
	var centerToOrigin vector.Vector = ray.Origin.Subtract(sphere.Center)
	var a float64 = ray.Direction.Dot(ray.Direction)
	var b float64 = 2.0 * centerToOrigin.Dot(ray.Direction)
	var c float64 = centerToOrigin.Dot(centerToOrigin) - sphere.Radius*sphere.Radius
	discriminant := b*b - 4*a*c
	position := ray.PointAtParameter(-b - math.Sqrt(discriminant)/(2.0*a))
	var normal vector.Vector = position.Subtract(sphere.Center).Normalize()
	var isHit bool
	if discriminant < 0 {
		isHit = false
	} else {
		isHit = true
	}
	distance := position.Distance(ray.Origin)
	return intersection.IntersectionPoint{
		IsHit:    isHit,
		Normal:   normal,
		Position: position,
		Distance: distance,
		Material: sphere.Material,
	}
}
