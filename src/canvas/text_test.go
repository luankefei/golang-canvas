package canvas

import (
	"fmt"
	"os"
	"testing"
	// "github.com/tdewolff/canvas"
)

// func TestDrawText(t *testing.T) {
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

// 测试文字
func TestDrawText(t *testing.T) {
	// var data []Drawer
	d, _ := os.Getwd()
	fmt.Println("location", d)
	Setup()

	text := Text{
		X:          50,
		Y:          50,
		Size:       20,
		LineHeight: 40,
		Color:      "#000000",
		Content:    "测试的文字，我爱北京天安门",
		FontStyle:  400,
		FontFamily: "PingFang",
	}

	// data = append(data, &text)
	// global := GlobalConfig{}
	// data[0].Draw()
	text.Draw()
}
