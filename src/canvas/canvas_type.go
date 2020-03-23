package canvas

import "github.com/tdewolff/canvas"

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
	width, height, x, y int32
}

// Text config
type Text struct {
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
	x, y, width, height, opacity, borderRadius int32
	imageURL                                   string
	clip                                       ImageClip
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
