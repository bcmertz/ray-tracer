package material

import (
	"ray-tracer/color"
)

type Material struct {
	DiffuseAlbedo  color.Color
	SpecularAlbedo color.Color
	Specular       bool
}
