package canvas

// import (
// 	"image"
// 	"image/color"

// 	"github.com/tdewolff/canvas"
// )

// // Path is alias of *canvas.Path
// type Path struct {
// 	p canvas.Path
// 	i *Image
// }

// // ColorModel 色彩
// func (p *Path) ColorModel() color.Model {
// 	return color.AlphaModel
// }

// // Bounds 区域
// func (p *Path) Bounds() image.Rectangle {
// 	// c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r
// 	b := p.p.Bounds()
// 	return image.Rect(int(b.X), int(b.Y), int(b.X+b.W), int(b.X+b.H))
// 	// return image.Rect(int(p.i.X), int(p.i.Y), int(p.i.X+p.i.Width), int(p.i.Y+p.i.Height))
// }

// // At 点值
// func (p *Path) At(x, y int) color.Color {
// 	b := p.p.Bounds()
// 	// xx, yy, rr := float64(x-p.i.X)+0.5, float64(y-p.i.Y)+0.5, float64(c.r)
// 	if float64(x) > (b.X+b.W) || float64(x) < b.X || float64(y) > (b.Y+b.H) || float64(y) < b.Y {
// 		// 原来的值是255
// 		return color.Alpha{255}
// 	}
// 	return color.Alpha{0}
// }
