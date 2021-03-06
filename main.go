package main

import (
	"fmt"
	"math/rand"
	"os"

	"ray-tracer/color"
	"ray-tracer/light"
	"ray-tracer/material"
	"ray-tracer/ray"
	"ray-tracer/scene"
	"ray-tracer/shapes"
	"ray-tracer/vector"
)

// coordinate grid dimensions, pixes are mapped onto this grid we define
// relative size of orthogonal vector lengths changes scaling / stretch of image
// if lines_x = 200 lines_y = 200 then X/Y vectors should be equal magnitude instead
// of 2:1 as is currently used for lines_x = 400 lines_y = 200
var lower_left = vector.Vector{
	X: -2.0,
	Y: -1.0,
	Z: -1.0,
}
var horiz = vector.Vector{
	X: 4.0,
	Y: 0.0,
	Z: 0.0,
}
var vert = vector.Vector{
	X: 0.0,
	Y: 2.0,
	Z: 0.0,
}
var origin = vector.Vector{
	X: 0.0,
	Y: 0.0,
	Z: 0.0,
}

func main() {
	f, err := os.Create("test.ppm")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	// define dimensions of the image
	const lines_x = 400.0
	const lines_y = 200.0
	err = writePPMHeader(f, lines_x, lines_y)
	if err != nil {
		fmt.Println(err)
	}
	// setup scene
	red := material.Material{
		DiffuseAlbedo: color.Color{
			R: 1.0,
			G: 0.0,
			B: 0.0,
		},
		Specular: false,
	}
	white := material.Material{
		DiffuseAlbedo: color.Color{
			R: 0.8,
			G: 0.8,
			B: 0.8,
		},
		SpecularAlbedo: color.Color{
			R: 0.2,
			G: 0.2,
			B: 0.2,
		},
		Specular: true,
	}
	blue := material.Material{
		DiffuseAlbedo: color.Color{
			R: 0.0,
			G: 0.0,
			B: 1.0,
		},
		Specular: false,
	}
	sphere1 := shapes.Sphere{
		Center: vector.Vector{
			X: 0.0,
			Y: 0.0,
			Z: -2.0,
		},
		Radius:   0.5,
		Material: red,
	}
	sphere2 := shapes.Sphere{
		Center: vector.Vector{
			X: 2.0,
			Y: 1.5,
			Z: -2.0,
		},
		Radius:   0.5,
		Material: white,
	}
	sphere3 := shapes.Sphere{
		Center: vector.Vector{
			X: 0.0,
			Y: -0.4,
			Z: -1.0,
		},
		Radius:   0.4,
		Material: blue,
	}
	light1 := light.Light{
		Position: vector.Vector{
			X: -1.0,
			Y: 200.0,
			Z: -1.0,
		},
		Power: color.Color{
			R: 0.0,
			G: 0.0,
			B: 10.0,
		},
		Attenuation: vector.Vector{X: 0.0, Y: 0.0, Z: 1},
	}
	scene := scene.Scene{
		Geometry: []shapes.Shape{sphere1, sphere2, sphere3},
		Lights:   []light.Light{light1},
	}
	samples_per_pixel := 20
	// iterate through the image pixels
	for y := lines_y - 1; y >= 0; y-- {
		for x := 0.0; x < lines_x; x++ {
			// sanity check
			fmt.Printf("\rAt image coordinate %v %v", x, y)
			var color color.Color
			for s := 0; s < samples_per_pixel; s++ {
				// scale pixel number to x,y grid size
				// also add in antialiasing w random offset sampling
				x_coord := horiz.ScalarMultiply(float64((x + rand.Float64()) / lines_x))
				y_coord := vert.ScalarMultiply(float64((y + rand.Float64()) / lines_y))
				// define the begining point of the ray and the direction vector into the image (from the camera)
				r := ray.Ray{
					Origin:    origin,
					Direction: lower_left.Add(x_coord).Add(y_coord), // from lower left hand side of the screen move up y_coord and right x_coord
				}
				c := trace(scene, r)
				color = color.Add(c)
			}
			// average out color from all samples
			color = color.ScalarDivide(float64(samples_per_pixel))
			// convert rgb colors to 255 values from unit color values
			var ir int = int(255 * color.R)
			var ig int = int(255 * color.G)
			var ib int = int(255 * color.B)
			// write out the pixel value
			_, err := f.WriteString(fmt.Sprintf("%v %v %v\n", ir, ig, ib))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

// see if we hit any objects, if not return background if yes return color of the shape
func trace(scene scene.Scene, r ray.Ray) color.Color {
	col := color.Color{
		R: 0,
		G: 0,
		B: 0,
	}
	intersectionPoint := scene.Intersect(r)
	if intersectionPoint.IsHit {
		col = intersectionPoint.Material.DiffuseAlbedo
		// do lighting
		//for _, light := range scene.Lights {
		// 	directionToLight := light.Position.Subtract(intersectionPoint.Position).Normalize()
		// 	shadowRay := ray.Ray{
		// 		Origin:    intersectionPoint.Position.Add(directionToLight).ScalarMultiply(0.00001),
		// 		Direction: directionToLight,
		// 	}
		// 	shadowIp := scene.Intersect(shadowRay)
		// 	if (!shadowIp.IsHit || (shadowIp.Distance > light.Distance(intersectionPoint.Position))) && directionToLight.Dot(intersectionPoint.Normal) > 0 {
		// 		intensity := light.IntensityAt(intersectionPoint.Position)
		// 		col = col.Add(intensity.ColorMultiply(intersectionPoint.Material.DiffuseAlbedo.ScalarMultiply(intersectionPoint.Normal.Dot(directionToLight))))
		// 	}
		// }

		return col
	} else {
		return renderBackground(r.Direction)
	}

}

// render linear gradient background if no objects are hit
func renderBackground(direction vector.Vector) color.Color {
	// gradient varies over y direction
	t := 0.5 * (direction.Y + 1)
	c1 := color.Color{
		R: 1.0,
		G: 1.0,
		B: 1.0,
	}
	c2 := color.Color{
		R: 0.5,
		G: 0.7,
		B: 1.0,
	}
	return c1.ScalarMultiply(1.0 - t).Add(c2.ScalarMultiply(t))

}

// write ppm file header
func writePPMHeader(f *os.File, rows, columns float64) error {
	const color_type = "P3"
	const max_color = 255
	_, err := f.WriteString(fmt.Sprintf("%s\n%v %v\n%v\n", color_type, rows, columns, max_color))
	return err
}
