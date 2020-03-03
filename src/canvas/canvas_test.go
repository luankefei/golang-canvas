package canvas

import (
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
func TestDrawText(t *testing.T) {
	// var data []Drawer

	text := Text{
		x:          0,
		y:          0,
		size:       10,
		lineHeight: 20,
		color:      "red",
		content:    "测试的文字，我爱北京天安门",
	}

	// data = append(data, &text)
	// global := GlobalConfig{}
	// data[0].Draw()
	text.Draw()
}
