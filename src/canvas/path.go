package canvas

import (
	"image"
	"image/color"
	"github.com/tdewolff/canvas"
)

// Path is alias of *canvas.Path
type Path struct {
	*canvas.Path
}

// ColorModel 色彩
func (p *Path) ColorModel() color.Model {
	return color.AlphaModel
}

// Bounds 区域
func (p *Path) Bounds() image.Rectangle {
// c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r
	return image.Rect(0, 0, 100, 100)
}

// At 点值
func (p *Path) At(x, y int) color.Color {
	// xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	// if xx*xx+yy*yy < rr*rr {
		// 原来的值是255
		return color.Alpha{170}
	// }
	// return color.Alpha{0}
}