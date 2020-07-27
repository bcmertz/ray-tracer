package main

import (
	"fmt"
	"os"

	"ray-tracer/color"
	"ray-tracer/ray"
	"ray-tracer/vector"
)

// coordinate grid dimensions, pixes are mapped onto this grid we define
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
	// define dimensions of the image
	const lines_x = 200.0
	const lines_y = 100.0
	// iterate through the image pixels
	for y := lines_y - 1; y >= 0; y-- {
		for x := 0.0; x < lines_x; x++ {
			// sanity check
			fmt.Printf("\rAt image coordinate %v %v", x, y)
			// scale pixel number to x,y grid size
			x_coord := horiz.ScalarMultiply(float64(x / lines_x))
			y_coord := vert.ScalarMultiply(float64(y / lines_y))
			// define the begining point of the ray and the direction vector into the image (from the camera)
			r := ray.Ray{
				Origin:    origin,
				Direction: lower_left.Add(x_coord).Add(y_coord), // from lower left hand side of the screen move up y_coord and right x_coord
			}
			c := colorGradient(r)
			// convert rgb colors to 255 values from unit color values
			var ir int = int(255 * c.R)
			var ig int = int(255 * c.G)
			var ib int = int(255 * c.B)
			// write out the pixel value
			_, err := f.WriteString(fmt.Sprintf("%v %v %v\n", ir, ig, ib))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

// create linear gradient of blue and white
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

// write ppm file header
func writePPMHeader(f *os.File) error {
	const color_type = "P3"
	const rows = 200
	const columns = 100
	const max_color = 255
	_, err := f.WriteString(fmt.Sprintf("%s\n%v %v\n%v\n", color_type, rows, columns, max_color))
	return err
}
