package canvas

import (
	"fmt"

	"github.com/luankefei/golang-canvas/src/libs"
	"github.com/tdewolff/canvas"
)

// TODO 使用单例替代
var fontFamily map[string]*canvas.FontFamily
var ptPerMm = 2.8346456692913384

// Draw text
func (t *Text) Draw(c *canvas.Context) {
	// color is HEX string, recommend use RGBA instead
	color := HexToColor(t.Color)
	content := t.Content
	indent := float64(0)

	// TODO: 不确定lineStretch的使用是否正确
	lineStretch := t.LineHeight - t.Size
	limit := float64(t.Limit)
	align := t.Align

	// TODO: fontFamily依赖loadFont的加载，理论上只需要加载一次，多字体可以实现并存
	fontKey := fmt.Sprintf("_font_%s_%d", t.FontFamily, t.FontStyle)

	// canvas.FontNormal
	face := fontFamily[fontKey].Face(t.Size*ptPerMm, color, t.FontStyle, canvas.FontNormal)

	metrics := face.Metrics()
	diff := metrics.Ascent - metrics.CapHeight

	// TODO: feature stroke
	// fmt.Println("draw stroke", t.X, t.Y*-1+diff+t.LineHeight, content)
	// p, _ := face.ToPath(content)
	// c.DrawPath(t.X, t.Y*-1+diff-t.LineHeight, p.Stroke(0.75, canvas.RoundCap, canvas.RoundJoin))

	// fmt.Println("metrics", metrics.Ascent, metrics.CapHeight, diff)
	// fmt.Println("line_break", limit, lineStretch)

	// TODO: height为0可能会导致折行不生效
	text := canvas.NewTextBox(
		face,
		content,
		limit,
		0,
		align,
		canvas.Top,
		indent,
		lineStretch,
	)

	fmt.Println("text_draw ", t.Size, content, indent, lineStretch)
	c.DrawText(t.X, t.Y*-1+diff, text)
	c.SetFillColor(color)
}

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
