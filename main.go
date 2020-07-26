package main

import (
	"fmt"
	"os"

	"ray-tracer/color"
	"ray-tracer/ray"
	"ray-tracer/vector"
)

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
	err = writePPMHeader(f)
	if err != nil {
		fmt.Println(err)
	}
	const nx = 200.0
	const ny = 100.0
	for j := ny - 1; j >= 0; j-- {
		for i := 0.0; i < nx; i++ {
			u := float64(i / nx)
			v := float64(j / ny)
			r := ray.Ray{
				Origin:    origin,
				Direction: lower_left.Add(horiz.ScalarMultiply(u)).Add(vert.ScalarMultiply(v)),
			}
			c := colorGradient(r)
			var ir int = int(256 * c.R)
			var ig int = int(256 * c.G)
			var ib int = int(256 * c.B)
			_, err := f.WriteString(fmt.Sprintf("%v %v %v\n", ir, ig, ib))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func colorGradient(r ray.Ray) color.Color {
	var direction vector.Vector = r.Direction
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

func writePPMHeader(f *os.File) error {
	const color_type = "P3"
	const rows = 200
	const columns = 100
	const max_color = 255
	_, err := f.WriteString(fmt.Sprintf("%s\n%v %v\n%v\n", color_type, rows, columns, max_color))
	return err
}
