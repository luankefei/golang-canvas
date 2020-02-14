
package canvas

import (
	"fmt"
)

// Draw text
func (i *Image) Draw() {
	fmt.Printf("%v text draw", &i)
	testText()
}

func testText() {
	c := canvas.New(265, 90)
	ctx := canvas.NewContext(c)
	draw(ctx)
	c.SavePNG("out.png", 5.0)
}

func drawText(c *canvas.Context, x, y float64, halign, valign canvas.TextAlign, indent float64) {
	face := fontFamily.Face(6.0, color.Black, canvas.FontRegular, canvas.FontNormal)
	// phrase := "测试的文本一，phase,测试的文本一，phase,测试的文本一，phase,测试的文本一，phase, 测试的文本一，phase,测试的文本一，phase,测试的文本一，phase"
	phrase := "hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world, yeah"

	text := canvas.NewTextBox(face, phrase, 60.0, 0.0, halign, valign, indent, 4.0)
	rect := text.Bounds()
	rect.Y = 0.0
	rect.H = -35.0
	c.SetFillColor(canvas.Whitesmoke)
	c.DrawPath(x, y, rect.ToPath())
	c.SetFillColor(canvas.Black)
	c.DrawText(x, y, text)
}