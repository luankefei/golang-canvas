package main

import (
	"image/color"
	"image/jpeg"
	"os"

	"github.com/tdewolff/canvas"
)

var fontFamily *canvas.FontFamily

func main() {
	// 加载本地字体
	leMiaoSrc := "../static/HanYiLeMiao_Regular.ttf"

	fontFamily = canvas.NewFontFamily("LeMiao")
	fontFamily.Use(canvas.CommonLigatures)
	if err := fontFamily.LoadFontFile(leMiaoSrc, canvas.FontRegular); err != nil {
		panic(err)
	}

	c := canvas.New(265, 90)
	ctx := canvas.NewContext(c)
	draw(ctx)
	c.SavePNG("out.png", 5.0)
}

func drawText(c *canvas.Context, x, y float64, halign, valign canvas.TextAlign, indent float64) {
	face := fontFamily.Face(6.0, color.Black, canvas.FontRegular, canvas.FontNormal)
	phrase := "测试的文本一，phase"

	text := canvas.NewTextBox(face, phrase, 60.0, 35.0, halign, valign, indent, 0.0)
	rect := text.Bounds()
	rect.Y = 0.0
	rect.H = -35.0
	c.SetFillColor(canvas.Whitesmoke)
	c.DrawPath(x, y, rect.ToPath())
	c.SetFillColor(canvas.Black)
	c.DrawText(x, y, text)
}

func drawImage(c *canvas.Context) {
	head, err := os.Open("../static/head.jpeg")
	if err != nil {
		panic(err)
	}

	img, err := jpeg.Decode(head)
	if err != nil {
		panic(err)
	}

	// x, y float64, img image.Image, dpm float64
	// Scale(1.0/dpm, 1.0/dpm)
	c.DrawImage(0, 0, img, 1)
}

func draw(c *canvas.Context) {
	face := fontFamily.Face(14.0, color.Black, canvas.FontRegular, canvas.FontNormal)
	c.SetFillColor(canvas.Black)
	c.DrawText(132.5, 88.0, canvas.NewTextBox(face, "测试的文本二", 0.0, 0.0, canvas.Center, canvas.Top, 0.0, 0.0))

	drawText(c, 5.0, 80.0, canvas.Left, canvas.Top, 10.0)
	drawText(c, 70.0, 80.0, canvas.Center, canvas.Top, 10.0)
	drawText(c, 135.0, 80.0, canvas.Right, canvas.Top, 10.0)
	drawText(c, 200.0, 80.0, canvas.Justify, canvas.Top, 10.0)
	drawText(c, 5.0, 40.0, canvas.Left, canvas.Top, 10.0)
	drawText(c, 70.0, 40.0, canvas.Left, canvas.Center, 10.0)
	drawText(c, 135.0, 40.0, canvas.Left, canvas.Bottom, 10.0)
	drawText(c, 200.0, 40.0, canvas.Left, canvas.Justify, 10.0)

	// test drawImage
	drawImage(c)
}
