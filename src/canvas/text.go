package canvas

import (
	"fmt"
	// "image/color"

	"github.com/luankefei/golang-canvas/src/libs"
	"github.com/tdewolff/canvas"
)

// TODO 放函数内部是否更合理
var fontFamily map[string]*canvas.FontFamily

// Draw text
func (t *Text) Draw(c *canvas.Context) {
	// TODO: color is not support HEX string, using RGBA instead
	color := HexToColor(t.Color)
	content := t.Content
	indent := float64(0)
	lineStretch := 0.0

	fmt.Println("Draw color", color)
	// TODO: fontFamily依赖loadFont的加载，理论上只需要加载一次，多字体可以实现并存
	fontKey := fmt.Sprintf("_font_%s_%d", t.FontFamily, t.FontStyle)
	face := fontFamily[fontKey].Face(t.Size, color, t.FontStyle, canvas.FontNormal)

	// c := canvas.New(750, 750)
	// ctx := canvas.NewContext(c)
	// ctx.SetView(canvas.Identity.Translate(0.0, 0.0))
	text := canvas.NewTextBox(face, content, 0.0, 0.0, canvas.Left, canvas.Top, indent, lineStretch)
	fmt.Println("text draw ", t.Size, content, indent, lineStretch, face)
	c.DrawText(t.X, t.Y, text)
	c.SetFillColor(color)
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
	fontKey := fmt.Sprintf("_font_%s_%d", name, style)
	fontFamily[fontKey] = canvas.NewFontFamily(name)

	fmt.Println("func.LoadFont: ", filepath, name, fontKey)

	// TODO: 这里将字体下载到本地 gitignore
	if err := libs.DownloadFile(fontKey, filepath); err != nil {
		panic(err)
	}

	// loads a font from a file.
	if err := fontFamily[fontKey].LoadFontFile(fontKey, style); err != nil {
		panic(err)
	}
}

// InitFont 初始化字体的入口, 由canvas.Setup函数调用
func InitFont() {
	fonts := make([]Font, 0)
	filepath := "../config/font.json"
	libs.LoadConfigFromJSON(filepath, &fonts)

	// 初始化fontFamily
	fontFamily = make(map[string]*canvas.FontFamily, len(fonts))

	for _, v := range fonts {
		// 从配置文件加载新字体
		LoadFont(v.FileName, v.Name, v.Style)
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

// rect := text.Bounds()
// rect.Y = 0.0
// rect.H = -35.0
// ctx.SetFillColor(canvas.Whitesmoke)
// ctx.DrawPath(t.X, t.Y, rect.ToPath())
