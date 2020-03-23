package canvas

import (
	"fmt"
	"testing"
)

// 测试画图主入口的参数传入
func TestCreateImage(t *testing.T) {
	// TODO: 临时测试代码
	// Setup()

	var data []Drawer

	i := Image{
		x: 1,
		y: 2,
	}
	data = append(data, &i)
	global := GlobalConfig{
		Width:    1000,
		Height:   1000,
		FileName: "out.png",
	}

	CreateImage(data, global)
}

func TestRGBAToColor(t *testing.T) {
	fmt.Println(RGBAToColor("rgba(255,255,255,1)"))
}

func TestHexToColor(t *testing.T) {
	fmt.Println(HexToColor("#F58C4BFF"))
}
