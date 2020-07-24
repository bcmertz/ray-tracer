package main

import (
	"fmt"
	"os"

	"ray-tracer/vector"
)

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
			v := vector.Vector{
				Color: [3]float64{float64(i / nx), float64(j / ny), 0.2},
			}
			var ir int = int(256 * v.R())
			var ig int = int(256 * v.G())
			var ib int = int(256 * v.B())
			_, err := f.WriteString(fmt.Sprintf("%v %v %v\n", ir, ig, ib))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func writePPMHeader(f *os.File) error {
	const color_type = "P3"
	const rows = 200
	const columns = 100
	const max_color = 255
	_, err := f.WriteString(fmt.Sprintf("%s\n%v %v\n%v\n", color_type, rows, columns, max_color))
	return err
}
