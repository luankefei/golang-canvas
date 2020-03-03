package canvas

import (
	"fmt"
	"image/color"

	"github.com/luankefei/golang-canvas/src/libs"
	"github.com/tdewolff/canvas"
)

// TODO 放函数内部是否更合理
var fontFamily *canvas.FontFamily

// Draw text
func (t *Text) Draw() {

	// TODO: color is not support HEX string, using RGBA instead
	color := HexToColor(t.Color)
	content := t.Content
	indent := float64(0)
	lineStretch := 0.0
	face := fontFamily.Face(t.Size, color, canvas.FontRegular, canvas.FontNormal)

	// box的宽高 0是auto
	text := canvas.NewTextBox(face, content, 0.0, 0.0, canvas.Left, canvas.Top, indent, lineStretch)
	// rect := text.Bounds()
	// rect.Y = 0.0
	// rect.H = -35.0

	fmt.Printf("text draw %v", t, text)
	// c := canvas.New(1000, 1000)
	// ctx := canvas.NewContext(c)
	// c.SetFillColor(canvas.Whitesmoke)
	// c.DrawPath(x, y, rect.ToPath())
	// c.SetFillColor(canvas.Black)
	// c.DrawText(x, y, text)
}

func drawText(c *canvas.Context, x, y float64, halign, valign canvas.TextAlign, indent float64) {
	face := fontFamily.Face(20.0, color.Black, canvas.FontRegular, canvas.FontNormal)
	// phrase := "测试的文本一，phase,测试的文本一，phase,测试的文本一，phase,测试的文本一，phase, 测试的文本一，phase,测试的文本一，phase,测试的文本一，phase"
	phrase := "hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world, yeah"

	// box的宽高 0是auto
	text := canvas.NewTextBox(face, phrase, 0.0, 0.0, halign, valign, indent, 0.0)
	rect := text.Bounds()
	rect.Y = 0.0
	rect.H = -35.0
	c.SetFillColor(canvas.Whitesmoke)
	c.DrawPath(x, y, rect.ToPath())
	c.SetFillColor(canvas.Black)
	c.DrawText(x, y, text)
}

// LoadFont 从本地文件注册字体
func LoadFont(filepath string, name string, style canvas.FontStyle) {
	fmt.Println("func.LoadFont: ", filepath, name)
	font := canvas.NewFontFamily(name)
	if err := libs.DownloadFile(name, filepath); err != nil {
		panic(err)
	}

	if err := font.LoadFontFile(name, style); err != nil {
		panic(err)
	}
}

// func testText() {
// 	c := canvas.New(1000, 1000)
// 	ctx := canvas.NewContext(c)

// 	matrix := canvas.Identity.Translate(0, 500)
// 	// .Rotate(180).ReflectY().

// 	ctx.SetView(matrix)
// 	// ctx.ComposeView(matrix)
// 	// ctx.ResetView()

// 	drawText(ctx, 0.0, 0.0, canvas.Left, canvas.Top, 0.0)

// 	// savePng的第二个参数是canvas导出时放大的倍数
// 	c.SavePNG("out.png", 1.0)
// }
