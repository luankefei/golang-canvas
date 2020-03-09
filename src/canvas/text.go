package canvas

import (
	"fmt"
	// "image/color"

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

	fmt.Println("Draw color", color)
	// TODO: fontFamily依赖loadFont的加载，理论上只需要加载一次，多字体可以实现并存
	face := fontFamily.Face(t.Size, color, t.FontStyle, canvas.FontNormal)

	// box的宽高 0是auto
	// c := canvas.New(265, 90)
	// ctx := canvas.NewContext(c)

	// rect := text.Bounds()
	// rect.Y = 0.0
	// rect.H = -35.0

	c := canvas.New(500, 300)
	ctx := canvas.NewContext(c)
	// ctx.SetView(canvas.Identity.Translate(0.0, 0.0))
	text := canvas.NewTextBox(face, content, 0.0, 0.0, canvas.Left, canvas.Top, indent, lineStretch)
	fmt.Println("text draw ", t.Size, content, indent, lineStretch, face)
	rect := text.Bounds()
	rect.Y = 0.0
	rect.H = -35.0
	ctx.SetFillColor(canvas.Whitesmoke)
	ctx.DrawPath(t.X, t.Y, rect.ToPath())
	ctx.DrawText(t.X, t.Y, text)

	ctx.SetFillColor(color)

	// 尽量导出2x或者3x的尺寸，但坐标是1x的，需要更多测试
	c.SavePNG("out.png", 3.0)
}

// func drawText(c *canvas.Context, x, y float64, halign, valign canvas.TextAlign, indent float64) {
// 	face := fontFamily.Face(20.0, color.Black, canvas.FontRegular, canvas.FontNormal)
// 	// phrase := "测试的文本一，phase,测试的文本一，phase,测试的文本一，phase,测试的文本一，phase, 测试的文本一，phase,测试的文本一，phase,测试的文本一，phase"
// 	phrase := "hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world, yeah"

// 	// box的宽高 0是auto
// 	text := canvas.NewTextBox(face, phrase, 0.0, 0.0, halign, valign, indent, 0.0)
// 	rect := text.Bounds()
// 	rect.Y = 0.0
// 	rect.H = -35.0
// 	c.SetFillColor(canvas.Whitesmoke)
// 	c.DrawPath(x, y, rect.ToPath())
// 	c.SetFillColor(canvas.Black)
// 	c.DrawText(x, y, text)
// }

// LoadFont 从本地文件注册字体
func LoadFont(filepath string, name string, style canvas.FontStyle) {
	fmt.Println("func.LoadFont: ", filepath, name)
	fontFamily = canvas.NewFontFamily(name)

	// TODO: 这里将字体下载到本地 gitignore
	if err := libs.DownloadFile("_font_"+name, filepath); err != nil {
		panic(err)
	}

	if err := fontFamily.LoadFontFile("_font_"+name, style); err != nil {
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
