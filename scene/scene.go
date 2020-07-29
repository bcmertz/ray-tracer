package scene

import (
	"ray-tracer/intersection"
	"ray-tracer/ray"
	"ray-tracer/shapes"
)

type Scene struct {
	Geometry []shapes.Shape
}

func (scene Scene) AddShape(shape shapes.Shape) {
	scene.Geometry = append(scene.Geometry, shape)
}

func (scene Scene) Intersect(ray ray.Ray) intersection.IntersectionPoint {
	closestIntersection := scene.Geometry[0].Intersect(ray)
	for _, object := range scene.Geometry {
		currentIp := object.Intersect(ray)
		if currentIp.IsHit && (!closestIntersection.IsHit || (closestIntersection.Distance > currentIp.Distance)) {
			closestIntersection = currentIp
		}
	}
	return closestIntersection
}
