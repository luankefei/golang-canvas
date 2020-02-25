package libs

import (
	"github.com/tdewolff/canvas"
)

// @see https://github.com/tdewolff/canvas/blob/master/examples/html-canvas/main.go
func RegisterFont() {
	family := canvas.NewFontFamily("dejavu-serif")
	family.LoadFontFile("font/DejaVuSerif.ttf", canvas.FontRegular)

	// size float64, col color.Color, style FontStyle, variant FontVariant, deco ...FontDecorator
	// face := family.Face(12.0*ptPerMm, canvas.Black, canvas.FontRegular, canvas.FontNormal)
	// test.Float(t, face.fauxBold, 0.0)
	// test.T(t, face.boldness(), 400)
}
