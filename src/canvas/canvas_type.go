package canvas

import "github.com/tdewolff/canvas"

// DrawWrapper 为json parse预留
type DrawWrapper struct {
	DrawType string `json:type`
}

// Drawer is general interface
type Drawer interface {
	Draw(c *canvas.Context)
}

// GlobalConfig type
type GlobalConfig struct {
	Width, Height, Compression float64
	FileName, MimeType         string
}

// ImageClip is clip image config
type ImageClip struct {
	Width, Height, X, Y int
}

// Text config
type Text struct {
	DrawWrapper
	Align                         canvas.TextAlign
	Size, X, Y, LineHeight, Limit float64
	Color, Content, FontFamily    string
	FontStyle                     canvas.FontStyle
}

// lineHeight 和 limit 主要用于文字多行需要计算折行的情况
// 如果文字是居中对齐，x和y值需要传入水平居中的中心点坐标
// EAlign: left | center | right | justify
// EFontWeight: bold | regular | normal
// IText {
//   x: number
//   y: number
//   content: string
//   color: string
//   size: number
//   align?: EAlign
//   fontWeight?: EFontWeight
//   lineHeight?: number
//   limit?: number
// }

// Image config
type Image struct {
	DrawWrapper
	X, Y, Width, Height, Opacity, BorderRadius float64
	ImageURL, MimeType                         string
	Buffer                                     []byte
	Clip                                       ImageClip
}

// Rect is a rectangle in 2D defined by a position and its width and height.
type Rect struct {
	X, Y, W, H float64
}

// Font define
type Font struct {
	FileName string           `json:filename`
	Name     string           `json:name`
	Style    canvas.FontStyle `json:style`
}

// import (
// 	"image"
// 	"image/color"
// )

// //Circle 圆形，实现了image.Image
// type Path struct {
// 	p image.Point
// 	r int
// }

// // ColorModel 色彩
// func (c *Circle) ColorModel() color.Model {
// 	return color.AlphaModel
// }

// // Bounds 区域
// func (c *Circle) Bounds() image.Rectangle {
// 	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
// }

// // At 点值
// func (c *Circle) At(x, y int) color.Color {
// 	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
// 	if xx*xx+yy*yy < rr*rr {
// 		return color.Alpha{255}
// 	}
// 	return color.Alpha{0}
// }
