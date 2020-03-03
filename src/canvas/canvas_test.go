package canvas

import (
	"fmt"
	"testing"
)

// 测试画图主入口的参数传入
func TestCreateImage(t *testing.T) {
	var data []Drawer

	i := Image{
		x: 1,
		y: 2,
	}
	data = append(data, &i)
	global := GlobalConfig{}

	CreateImage(data, global)
}

// 测试文字
// func TestDrawText(t *testing.T) {
// 	// var data []Drawer

// 	text := Text{
// 		X:          0,
// 		Y:          0,
// 		Size:       10,
// 		LineHeight: 20,
// 		Color:      "red",
// 		Content:    "测试的文字，我爱北京天安门",
// 	}

// 	// data = append(data, &text)
// 	// global := GlobalConfig{}
// 	// data[0].Draw()
// 	text.Draw()
// }

func TestRGBAToColor(t *testing.T) {
	fmt.Println(RGBAToColor("rgba(255,255,255,1)"))
}

func TestHexToColor(t *testing.T) {
	fmt.Println(HexToColor("#F58C4BFF"))
}
