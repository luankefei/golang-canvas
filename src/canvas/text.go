package canvas

import (
	"fmt"
	"image/color"

	"github.com/tdewolff/canvas"
)

// TODO 放函数内部是否更合理
var fontFamily *canvas.FontFamily

// Draw text
func (i *Text) Draw() {
	fmt.Printf("%v text draw", &i)
	testText()
}

func testLoadFont() {
	// 从文件加载新字体
	leMiaoSrc := "../static/HanYiLeMiao_Regular.ttf"

	fontFamily = canvas.NewFontFamily("LeMiao")
	fontFamily.Use(canvas.CommonLigatures)

	// TODO: 可以用LoadLocalFont
	if err := fontFamily.LoadFontFile(leMiaoSrc, canvas.FontRegular); err != nil {
		panic(err)
	}
}

func testText() {
	testLoadFont()

	c := canvas.New(1000, 1000)
	ctx := canvas.NewContext(c)

	matrix := canvas.Identity.Translate(0, 500)
	// .Rotate(180).ReflectY().

	ctx.SetView(matrix)
	// ctx.ComposeView(matrix)
	// ctx.ResetView()

	drawText(ctx, 0.0, 0.0, canvas.Left, canvas.Top, 0.0)

	// savePng的第二个参数是canvas导出时放大的倍数
	c.SavePNG("out.png", 1.0)
}

func drawText(c *canvas.Context, x, y float64, halign, valign canvas.TextAlign, indent float64) {
	face := fontFamily.Face(20.0, color.Black, canvas.FontRegular, canvas.FontNormal)
	// phrase := "测试的文本一，phase,测试的文本一，phase,测试的文本一，phase,测试的文本一，phase, 测试的文本一，phase,测试的文本一，phase,测试的文本一，phase"
	phrase := "hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world,hello, world, yeah"

	// box的宽高 0是auto
	text := canvas.NewTextBox(face, phrase, 0.0, 0.0, halign, valign, indent, 0.0)
	rect := text.Bounds()
	// rect.Y = 0.0
	rect.H = -35.0
	c.SetFillColor(canvas.Whitesmoke)
	c.DrawPath(x, y, rect.ToPath())
	c.SetFillColor(canvas.Black)
	c.DrawText(x, y, text)
}
